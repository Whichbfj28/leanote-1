package service

import (
	//	"fmt"
	"context"
	"sort"
	"strings"
	"time"

	"github.com/coocn-cn/leanote/app/info"
	. "github.com/coocn-cn/leanote/app/lea"
	"github.com/coocn-cn/leanote/app/note/model"
	"github.com/coocn-cn/leanote/app/note/repository"
	"github.com/coocn-cn/leanote/pkg/errcode"
	"github.com/coocn-cn/leanote/pkg/log"
	"gopkg.in/mgo.v2/bson"
	//	"html"
)

// 笔记本

type NotebookService struct {
	note    repository.NoteRepository
	book    repository.BookRepository
	content repository.ContentRepository
}

// 排序
func sortSubNotebooks(eachNotebooks info.SubNotebooks) info.SubNotebooks {
	// 遍历子, 则子往上进行排序
	for _, eachNotebook := range eachNotebooks {
		if eachNotebook.Subs != nil && len(eachNotebook.Subs) > 0 {
			eachNotebook.Subs = sortSubNotebooks(eachNotebook.Subs)
		}
	}

	// 子排完了, 本层排
	sort.Sort(&eachNotebooks)
	return eachNotebooks
}

// 整理(成有关系)并排序
// GetNotebooks()调用
// ShareService调用
func ParseAndSortNotebooks(ctx context.Context, userNotebooks []*model.Book, noParentDelete, needSort bool) info.SubNotebooks {
	// 整理成info.Notebooks
	// 第一遍, 建map
	// notebookId => info.Notebooks
	userNotebooksMap := make(map[bson.ObjectId]*info.Notebooks, len(userNotebooks))
	for _, book := range userNotebooks {
		each := book.MustData(ctx)
		newNotebooks := info.Notebooks{Subs: info.SubNotebooks{}}
		newNotebooks.NotebookId = each.NotebookId
		newNotebooks.Title = each.Title
		//		newNotebooks.Title = html.EscapeString(each.Title)
		newNotebooks.Title = strings.Replace(strings.Replace(each.Title, "<script>", "", -1), "</script", "", -1)
		newNotebooks.Seq = each.Seq
		newNotebooks.UserId = each.UserId
		newNotebooks.ParentNotebookId = each.ParentNotebookId
		newNotebooks.NumberNotes = each.NumberNotes
		newNotebooks.IsTrash = each.IsTrash
		newNotebooks.IsBlog = each.IsBlog

		// 存地址
		userNotebooksMap[each.NotebookId] = &newNotebooks
	}

	// 第二遍, 追加到父下

	// 需要删除的id
	needDeleteNotebookId := map[bson.ObjectId]bool{}
	for id, each := range userNotebooksMap {
		// 如果有父, 那么追加到父下, 并剪掉当前, 那么最后就只有根的元素
		if each.ParentNotebookId.Hex() != "" {
			if userNotebooksMap[each.ParentNotebookId] != nil {
				userNotebooksMap[each.ParentNotebookId].Subs = append(userNotebooksMap[each.ParentNotebookId].Subs, each) // Subs是存地址
				// 并剪掉
				// bug
				needDeleteNotebookId[id] = true
				// delete(userNotebooksMap, id)
			} else if noParentDelete {
				// 没有父, 且设置了要删除
				needDeleteNotebookId[id] = true
				// delete(userNotebooksMap, id)
			}
		}
	}

	// 第三遍, 得到所有根
	final := make(info.SubNotebooks, len(userNotebooksMap)-len(needDeleteNotebookId))
	i := 0
	for id, each := range userNotebooksMap {
		if !needDeleteNotebookId[id] {
			final[i] = each
			i++
		}
	}

	// 最后排序
	if needSort {
		return sortSubNotebooks(final)
	}
	return final
}

// 得到某notebook
func (this *NotebookService) GetNotebook(notebookId, userId string) info.Notebook {
	book := this.GetNotebookById(notebookId)

	if book.UserId.Hex() != userId {
		return info.Notebook{}
	}

	return book
}

func (this *NotebookService) GetNotebookById(notebookId string) info.Notebook {
	m := this
	ctx := context.Background()

	book, err := m.book.Find(ctx, repository.BookID(notebookId))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记本失败")
		return info.Notebook{}
	}

	if book == nil {
		return info.Notebook{}
	}

	return info.Notebook(book.MustData(ctx))
}

func (this *NotebookService) GetNotebookByUserIdAndUrlTitle(userId, notebookIdOrUrlTitle string) info.Notebook {
	m := this
	ctx := context.Background()

	var predicate repository.Predicater
	if IsObjectId(notebookIdOrUrlTitle) {
		predicate = repository.BookID(notebookIdOrUrlTitle)
	} else {
		predicate = repository.BookUserAndURLTitle(userId, notebookIdOrUrlTitle)
	}

	book, err := m.book.Find(ctx, predicate)
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记本失败")
		return info.Notebook{}
	}

	if book == nil {
		return info.Notebook{}
	}

	return info.Notebook(book.MustData(ctx))
}

// 同步的方法
func (m *NotebookService) GeSyncNotebooks(userId string, afterUsn, maxEntry int) []info.Notebook {
	ctx := context.Background()

	books, err := m.book.FindAll(ctx, repository.BookUSNNextBooks(userId, afterUsn, maxEntry))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记本失败")
		return nil
	}

	resp := make([]info.Notebook, 0, len(books))
	for _, v := range books {
		resp = append(resp, info.Notebook(v.MustData(ctx)))
	}

	return resp
}

// 得到用户下所有的notebook
// 排序好之后返回
// [ok]
func (m *NotebookService) GetNotebooks(userId string) info.SubNotebooks {
	ctx := context.Background()

	userNotebooks, err := m.book.FindAll(ctx, repository.BookUserAndNotDelete(userId))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记本失败")
		return nil
	}

	if len(userNotebooks) == 0 {
		return nil
	}

	return ParseAndSortNotebooks(ctx, userNotebooks, true, true)
}

// share调用, 不需要删除没有父的notebook
// 不需要排序, 因为会重新排序
// 通过notebookIds得到notebooks, 并转成层次有序
func (m *NotebookService) GetNotebooksByNotebookIds(notebookIds []bson.ObjectId) info.SubNotebooks {
	ctx := context.Background()

	ids := make([]string, 0, len(notebookIds))
	for _, v := range notebookIds {
		ids = append(ids, v.Hex())
	}

	userNotebooks, err := m.book.FindAll(ctx, repository.BookIDs(ids))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记本失败")
		return nil
	}

	if len(userNotebooks) == 0 {
		return nil
	}

	return ParseAndSortNotebooks(ctx, userNotebooks, false, false)
}

// 添加
func (m *NotebookService) AddNotebook(notebook info.Notebook) (bool, info.Notebook) {
	ctx := context.Background()

	if notebook.NotebookId == "" {
		notebook.NotebookId = bson.NewObjectId()
	}

	notebook.UrlTitle = GetUrTitle(notebook.UserId.Hex(), notebook.Title, "notebook", notebook.NotebookId.Hex())
	notebook.Usn = userService.IncrUsn(notebook.UserId.Hex())
	now := time.Now()
	notebook.CreatedTime = now
	notebook.UpdatedTime = now

	err := m.book.Save(ctx, m.book.New(ctx, model.BookData(notebook)))
	if err != nil {
		return false, notebook
	}

	return true, notebook
}

// 更新笔记, api
func (m *NotebookService) UpdateNotebookApi(userId, notebookId, title, parentNotebookId string, seq, usn int) (bool, string, info.Notebook) {
	ctx := context.Background()

	if notebookId == "" {
		return false, "notebookIdNotExists", info.Notebook{}
	}

	err := m.update(ctx, repository.BookUserAndID(userId, notebookId), func(book *model.Book) error {
		if book == nil {
			return errcode.NotFound(ctx, "notExists", notebookId)
		}

		return book.UpdateNotebookApi(ctx, userId, title, parentNotebookId, seq, usn, userService.IncrUsn(userId))
	})
	if err != nil {
		return false, err.Error(), info.Notebook{}
	}

	return true, "", m.GetNotebook(userId, notebookId)
}

// 判断是否是blog
func (this *NotebookService) IsBlog(notebookId string) bool {
	m := this
	ctx := context.Background()

	book, err := m.book.Find(ctx, repository.BookID(notebookId))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记本失败")
		return false
	}

	if book == nil {
		return false
	}

	return book.MustData(ctx).IsBlog
}

// 判断是否是我的notebook
func (m *NotebookService) IsMyNotebook(notebookId, userId string) bool {
	ctx := context.Background()

	count, err := m.book.Count(ctx, repository.BookUserAndID(userId, notebookId))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记本失败")
		return false
	}

	return count != 0
}

// 更新笔记本标题
// [ok]
func (m *NotebookService) UpdateNotebookTitle(notebookId, userId, title string) bool {
	ctx := context.Background()

	err := m.update(ctx, repository.BookUserAndID(userId, notebookId), func(book *model.Book) error {
		if book == nil {
			return nil
		}

		return book.UpdateTitle(ctx, userId, title, userService.IncrUsn(userId))
	})

	if err != nil {
		log.G(ctx).WithError(err).Error("更新笔记本失败")
		return false
	}

	return true
}

// ToBlog or Not
func (m *NotebookService) ToBlog(userId, notebookId string, isBlog bool) bool {
	ctx := context.Background()

	// 笔记本
	err := m.update(ctx, repository.BookUserAndID(userId, notebookId), func(book *model.Book) error {
		if book == nil {
			return nil
		}

		newUSN := userService.IncrUsn(userId)
		err := updateNodes(m.note, ctx, repository.NoteBookID(book.MustData(ctx).NotebookId.Hex()), func(notes []*model.Note) error {
			ids := make([]string, 0, len(notes))
			for _, note := range notes {
				ids = append(ids, note.MustData(ctx).NoteId.Hex())

				if err := note.SetBlogStatus(ctx, isBlog, newUSN); err != nil {
					return err
				}
			}

			return updateContents(m.content, ctx, repository.ContentNoteIDs(ids), func(contents []*model.Content) error {
				for _, content := range contents {
					if err := content.SetBlogStatus(ctx, isBlog); err != nil {
						return err
					}
				}

				return nil
			})
		})
		if err != nil {
			return err
		}

		return book.SetBlogStatus(ctx, isBlog, newUSN)
	})

	if err != nil {
		log.G(ctx).WithError(err).Error("更新笔记本失败")
		return false
	}

	// 重新计算tags
	go (func() {
		blogService.ReCountBlogTags(userId)
	})()

	return true
}

func (m *NotebookService) DeleteNotebook(userId, notebookId string) (bool, string) {
	ctx := context.Background()

	err := m.update(ctx, repository.BookUserAndID(userId, notebookId), func(book *model.Book) error {
		if book == nil {
			return nil
		}

		data := book.MustData(ctx)

		childs, err := m.book.FindAll(ctx, repository.BookUserAndParentIDAndDelete(data.UserId.Hex(), data.NotebookId.Hex(), false))
		if err != nil {
			return err
		}

		if len(childs) != 0 {
			return errcode.FailedPrecondition(ctx, "笔记本下有子笔记本")
		}

		notes, err := m.note.FindAll(ctx, repository.NoteBookID(data.NotebookId.Hex()))
		if err != nil {
			return err
		}

		if len(notes) != 0 {
			return errcode.FailedPrecondition(ctx, "笔记本下有笔记")
		}

		return book.SoftDelete(ctx, userService.IncrUsn(userId))
	})

	if err != nil {
		return false, err.Error()
	}

	return true, ""

}

// API调用, 删除笔记本, 不作笔记控制
func (m *NotebookService) DeleteNotebookForce(userId, notebookId string, usn int) (bool, string) {
	ctx := context.Background()

	book, err := m.book.Find(ctx, repository.BookUserAndID(userId, notebookId))
	if err != nil {
		return false, err.Error()
	}

	if book == nil {
		return false, "notExists"
	}

	if book.MustData(ctx).Usn != usn {
		return false, "conflict"
	}

	if err := m.book.Delete(ctx, book); err != nil {
		return false, err.Error()
	}

	return true, ""
}

// 排序
// 传入 notebookId => Seq
// 为什么要传入userId, 防止修改其它用户的信息 (恶意)
// [ok]
func (m *NotebookService) SortNotebooks(userId string, notebookId2Seqs map[string]int) bool {
	ctx := context.Background()

	if len(notebookId2Seqs) == 0 {
		return false
	}

	ids := make([]string, 0, len(notebookId2Seqs))
	for id := range notebookId2Seqs {
		ids = append(ids, id)
	}

	err := m.updates(ctx, repository.BookUserAndIDs(userId, ids), func(book *model.Book) error {
		return book.SetSortWeight(ctx, notebookId2Seqs[book.MustData(ctx).NotebookId.Hex()], userService.IncrUsn(userId))
	})

	if err != nil {
		log.G(ctx).WithError(err).Error("更新笔记本失败")
		return false
	}

	return true
}

// 排序和设置父
func (m *NotebookService) DragNotebooks(userId string, curNotebookId string, parentNotebookId string, siblings []string) bool {
	ctx := context.Background()

	err := m.update(ctx, repository.BookUserAndID(userId, curNotebookId), func(book *model.Book) error {
		if book == nil {
			return nil
		}

		var err error
		var perent *model.Book = nil

		if parentNotebookId != "" {
			perent, err = m.book.Find(ctx, repository.BookID(parentNotebookId))
			if err != nil {
				return err
			}
		}

		return book.SetParent(ctx, perent, userService.IncrUsn(userId))
	})

	if err != nil {
		log.G(ctx).WithError(err).Error("更新笔记本失败")
		return false
	}

	// 排序
	notebookId2Seqs := make(map[string]int, len(siblings))
	for seq, notebookId := range siblings {
		notebookId2Seqs[notebookId] = seq
	}

	return m.SortNotebooks(userId, notebookId2Seqs)
}

// 重新统计笔记本下的笔记数目
// noteSevice: AddNote, CopyNote, CopySharedNote, MoveNote
// trashService: DeleteNote (recove不用, 都统一在MoveNote里了)
func (m *NotebookService) ReCountNotebookNumberNotes(notebookId string) bool {
	ctx := context.Background()

	err := m.update(ctx, repository.BookID(notebookId), func(book *model.Book) error {
		if book == nil {
			return nil
		}

		notes, err := m.note.FindAll(ctx, repository.NoteBookID(book.MustData(ctx).NotebookId.Hex()))
		if err != nil {
			return err
		}

		return book.RefreshNumberNotes(ctx, len(notes))
	})

	if err != nil {
		log.G(ctx).WithError(err).Error("更新笔记本失败")
		return false
	}
	// Log(count)
	// Log(notebookId)
	return true
}

func (m *NotebookService) update(ctx context.Context, predicate repository.Predicater, f func(*model.Book) error) error {
	book, err := m.book.Find(ctx, predicate)
	if err != nil {
		return err
	}

	if err := f(book); err != nil {
		return err
	}

	return m.book.Save(ctx, book)
}

func (m *NotebookService) updates(ctx context.Context, predicate repository.Predicater, f func(*model.Book) error) error {
	books, err := m.book.FindAll(ctx, predicate)
	if err != nil {
		return err
	}

	for _, book := range books {
		if err := f(book); err != nil {
			return err
		}
	}

	return m.book.Save(ctx, books...)
}
