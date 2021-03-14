package service

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/coocn-cn/leanote/app/info"
	. "github.com/coocn-cn/leanote/app/lea"
	"github.com/coocn-cn/leanote/app/note/model"
	"github.com/coocn-cn/leanote/app/note/repository"
	"github.com/coocn-cn/leanote/pkg/errcode"
	"github.com/coocn-cn/leanote/pkg/log"
	"gopkg.in/mgo.v2/bson"
)

type NoteService struct {
	note    repository.NoteRepository
	content repository.ContentRepository
	history repository.HistoryRepository
}

// 通过id, userId得到note, 包含已删除的
func (this *NoteService) GetNote(noteId, userId string) info.Note {
	ctx := context.Background()

	note, err := this.note.Find(ctx, repository.ID(noteId).WithUser(userId))
	if err != nil || note == nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return info.Note{}
	}

	return info.Note(note.MustData(ctx))
}

// fileService调用
// 不能是已经删除了的, life bug, 客户端删除后, 竟然还能在web上打开
func (this *NoteService) GetNoteById(noteId string) info.Note {
	ctx := context.Background()
	if noteId == "" {
		return info.Note{}
	}

	note, err := this.note.Find(ctx, repository.ID(noteId).WithDeleted(false))
	if err != nil || note == nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return info.Note{}
	}

	return info.Note(note.MustData(ctx))
}

// 通过id, userId得到note, 不包含已删除的
func (this *NoteService) GetNoteByIdAndUserId(noteId, userId string) info.Note {
	ctx := context.Background()
	if noteId == "" || userId == "" {
		return info.Note{}
	}

	note, err := this.note.Find(ctx, repository.ID(noteId).WithUser(userId).WithDeleted(false))
	if err != nil || note == nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return info.Note{}
	}

	return info.Note(note.MustData(ctx))
}

// 得到blog, blogService用
// 不要传userId, 因为是公开的
func (this *NoteService) GetBlogNote(noteId string) info.Note {
	ctx := context.Background()

	note, err := this.note.Find(ctx, repository.NoteIDAndBlog(noteId, true).WithTrash(false).WithDeleted(false))
	if err != nil || note == nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return info.Note{}
	}

	return info.Note(note.MustData(ctx))
}

// 通过id, userId得到noteContent
func (this *NoteService) GetNoteContent(noteContentId, userId string) (noteContent info.NoteContent) {
	ctx := context.Background()

	content, err := this.content.Find(ctx, repository.ID(noteContentId).WithUser(userId))
	if err != nil || content == nil {
		log.G(ctx).WithError(err).Error("获取笔记内容失败")
		return info.NoteContent{}
	}

	return info.NoteContent(content.MustData(ctx))
}

// 得到笔记和内容
func (this *NoteService) GetNoteAndContent(noteId, userId string) (noteAndContent info.NoteAndContent) {
	note := this.GetNote(noteId, userId)
	noteContent := this.GetNoteContent(noteId, userId)
	return info.NoteAndContent{note, noteContent}
}

func (this *NoteService) GetNoteBySrc(src, userId string) info.Note {
	ctx := context.Background()
	if src == "" {
		return info.Note{}
	}

	note, err := this.note.Find(ctx, repository.NoteSrc(src).WithUser(userId).WithSort("-Usn"))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return info.Note{}
	}

	return info.Note(note.MustData(ctx))
}

func (this *NoteService) GetNoteAndContentBySrc(src, userId string) (noteId string, noteAndContent info.NoteAndContentSep) {
	note := this.GetNoteBySrc(src, userId)
	if note.NoteId != "" {
		noteId = note.NoteId.Hex()
		noteContent := this.GetNoteContent(note.NoteId.Hex(), userId)
		return noteId, info.NoteAndContentSep{note, noteContent}
	}
	return
}

// 获取同步的笔记
// > afterUsn的笔记
func (this *NoteService) GetSyncNotes(userId string, afterUsn, maxEntry int) []info.ApiNote {
	ctx := context.Background()

	books, err := this.note.FindAll(ctx, repository.NoteNexts(afterUsn).WithUser(userId).WithLimit(maxEntry))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return nil
	}

	notes := make([]info.Note, 0, len(books))
	for _, v := range books {
		notes = append(notes, info.Note(v.MustData(ctx)))
	}

	return this.ToApiNotes(notes)
}

// note与apiNote的转换
func (this *NoteService) ToApiNotes(notes []info.Note) []info.ApiNote {
	// 2, 得到所有图片, 附件信息
	// 查images表, attachs表
	if len(notes) > 0 {
		noteIds := make([]bson.ObjectId, len(notes))
		for i, note := range notes {
			noteIds[i] = note.NoteId
		}
		noteFilesMap := this.getFiles(noteIds)
		// 生成info.ApiNote
		apiNotes := make([]info.ApiNote, len(notes))
		for i, note := range notes {
			noteId := note.NoteId.Hex()
			apiNotes[i] = this.ToApiNote(&note, noteFilesMap[noteId])
		}
		return apiNotes
	}
	// 返回空的
	return []info.ApiNote{}
}

// note与apiNote的转换
func (this *NoteService) ToApiNote(note *info.Note, files []info.NoteFile) info.ApiNote {
	apiNote := info.ApiNote{
		NoteId:      note.NoteId.Hex(),
		NotebookId:  note.NotebookId.Hex(),
		UserId:      note.UserId.Hex(),
		Title:       note.Title,
		Tags:        note.Tags,
		IsMarkdown:  note.IsMarkdown,
		IsBlog:      note.IsBlog,
		IsTrash:     note.IsTrash,
		IsDeleted:   note.IsDeleted,
		Usn:         note.Usn,
		CreatedTime: note.CreatedTime,
		UpdatedTime: note.UpdatedTime,
		PublicTime:  note.PublicTime,
		Files:       files,
	}
	return apiNote
}

// getDirtyNotes, 把note的图片, 附件信息都发送给客户端
// 客户端保存到本地, 再获取图片, 附件

// 得到所有图片, 附件信息
// 查images表, attachs表
// [待测]
func (this *NoteService) getFiles(noteIds []bson.ObjectId) map[string][]info.NoteFile {
	noteImages := noteImageService.getImagesByNoteIds(noteIds)
	noteAttachs := attachService.getAttachsByNoteIds(noteIds)

	noteFilesMap := map[string][]info.NoteFile{}

	for _, noteId := range noteIds {
		noteIdHex := noteId.Hex()
		noteFiles := []info.NoteFile{}
		// images
		if images, ok := noteImages[noteIdHex]; ok {
			for _, image := range images {
				noteFiles = append(noteFiles, info.NoteFile{
					FileId: image.FileId.Hex(),
					Type:   image.Type,
				})
			}
		}

		// attach
		if attachs, ok := noteAttachs[noteIdHex]; ok {
			for _, attach := range attachs {
				noteFiles = append(noteFiles, info.NoteFile{
					FileId:   attach.AttachId.Hex(),
					Type:     attach.Type,
					Title:    attach.Title,
					IsAttach: true,
				})
			}
		}

		noteFilesMap[noteIdHex] = noteFiles
	}

	return noteFilesMap
}

// 列出note, 排序规则, 还有分页
// CreatedTime, UpdatedTime, title 来排序
func (this *NoteService) ListNotes(userId, notebookId string, isTrash bool, pageNumber, pageSize int, sortField string, isAsc bool, isBlog bool) (int, []info.Note) {
	ctx := context.Background()
	skipNum, sortFieldR := parsePageAndSort(pageNumber, pageSize, sortField, isAsc)

	predicate := repository.All()
	if isBlog {
		predicate = repository.NoteBlog()
	}
	if notebookId != "" {
		predicate = repository.NoteBookIDAndBlog(notebookId, isBlog)
	}

	predicate = predicate.WithUser(userId).WithDeleted(false).WithTrash(isTrash)

	count, err := this.note.Count(ctx, predicate)
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return 0, nil
	}

	notes, err := this.note.FindAll(ctx, predicate.WithSort(sortFieldR).WithPage(skipNum, pageSize))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return 0, nil
	}

	resp := make([]info.Note, 0, len(notes))
	for _, v := range notes {
		resp = append(resp, info.Note(v.MustData(ctx)))
	}

	return count, resp
}

// 通过noteIds来查询
// ShareService调用
func (this *NoteService) ListNotesByNoteIdsWithPageSort(noteIds []bson.ObjectId, userId string, pageNumber, pageSize int, sortField string, isAsc bool, isBlog bool) []info.Note {
	ctx := context.Background()
	skipNum, sortFieldR := parsePageAndSort(pageNumber, pageSize, sortField, isAsc)

	ids := make([]string, 0, len(noteIds))
	for _, v := range noteIds {
		ids = append(ids, v.Hex())
	}

	userNotebooks, err := this.note.FindAll(ctx, repository.IDs(ids).WithTrash(false).WithSort(sortFieldR).WithPage(skipNum, pageSize))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记本失败")
		return nil
	}

	notes := make([]info.Note, 0, len(userNotebooks))
	for _, note := range userNotebooks {
		notes = append(notes, info.Note(note.MustData(ctx)))
	}

	return notes
}

// shareService调用
func (this *NoteService) ListNotesByNoteIds(noteIds []bson.ObjectId) []info.Note {
	ctx := context.Background()

	ids := make([]string, 0, len(noteIds))
	for _, v := range noteIds {
		ids = append(ids, v.Hex())
	}

	userNotebooks, err := this.note.FindAll(ctx, repository.IDs(ids))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记本失败")
		return nil
	}

	notes := make([]info.Note, 0, len(userNotebooks))
	for _, note := range userNotebooks {
		notes = append(notes, info.Note(note.MustData(ctx)))
	}

	return notes
}

// blog需要
func (this *NoteService) ListNoteContentsByNoteIds(noteIds []bson.ObjectId, fields []string) (notes []info.NoteContent) {
	ctx := context.Background()

	ids := make([]string, 0, len(noteIds))
	for _, v := range noteIds {
		ids = append(ids, v.Hex())
	}
	predicate := repository.IDs(ids)
	if len(fields) != 0 {
		predicate = predicate.WithSelect(fields)
	}
	records, err := this.content.FindAll(ctx, predicate)
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记本失败")
		return nil
	}

	resp := make([]info.NoteContent, 0, len(records))
	for _, row := range records {
		resp = append(resp, info.NoteContent(row.MustData(ctx)))
	}

	return resp
}

// 只得到abstract, 不需要content
func (this *NoteService) ListNoteAbstractsByNoteIds(noteIds []bson.ObjectId) (notes []info.NoteContent) {
	return this.ListNoteContentsByNoteIds(noteIds, []string{"_id", "Abstract"})
}

func (this *NoteService) ListNoteContentByNoteIds(noteIds []bson.ObjectId) (notes []info.NoteContent) {
	return this.ListNoteContentsByNoteIds(noteIds, []string{"_id", "Abstract", "Content"})
}

// 添加笔记
// 首先要判断Notebook是否是Blog, 是的话设为blog
// [ok]

func (this *NoteService) AddNote(note info.Note, fromApi bool) info.Note {
	ctx := context.Background()
	if note.NoteId.Hex() == "" {
		noteId := bson.NewObjectId()
		note.NoteId = noteId
	}

	// 关于创建时间, 可能是客户端发来, 此时判断时间是否有
	note.CreatedTime = FixUrlTime(note.CreatedTime)
	note.UpdatedTime = FixUrlTime(note.UpdatedTime)

	note.IsTrash = false
	note.UpdatedUserId = note.UserId
	note.UrlTitle = GetUrTitle(note.UserId.Hex(), note.Title, "note", note.NoteId.Hex())
	note.Usn = userService.IncrUsn(note.UserId.Hex())

	notebookId := note.NotebookId.Hex()

	// api会传IsBlog, web不会传
	if !fromApi {
		// 设为blog
		note.IsBlog = notebookService.IsBlog(notebookId)
	}
	//	if note.IsBlog {
	note.PublicTime = note.UpdatedTime
	//	}

	if err := this.note.Save(ctx, this.note.New(ctx, model.NoteData(note))); err != nil {
		return note
	}

	// tag1
	tagService.AddTags(note.UserId.Hex(), note.Tags)

	// recount notebooks' notes number
	notebookService.ReCountNotebookNumberNotes(notebookId)

	return note
}

// 添加共享d笔记
func (this *NoteService) AddSharedNote(note info.Note, myUserId bson.ObjectId) info.Note {
	// 判断我是否有权限添加
	if shareService.HasUpdateNotebookPerm(note.UserId.Hex(), myUserId.Hex(), note.NotebookId.Hex()) {
		note.CreatedUserId = myUserId // 是我给共享我的人创建的
		return this.AddNote(note, false)
	}
	return info.Note{}
}

// 添加笔记本内容
// [ok]
func (this *NoteService) AddNoteContent(noteContent info.NoteContent) info.NoteContent {
	ctx := context.Background()
	noteContent.CreatedTime = FixUrlTime(noteContent.CreatedTime)
	noteContent.UpdatedTime = FixUrlTime(noteContent.UpdatedTime)

	noteContent.UpdatedUserId = noteContent.UserId
	if err := this.content.Save(ctx, this.content.New(ctx, model.ContentData(noteContent))); err != nil {
		return noteContent
	}

	// 更新笔记图片
	noteImageService.UpdateNoteImages(noteContent.UserId.Hex(), noteContent.NoteId.Hex(), "", noteContent.Content)

	return noteContent
}

// 添加笔记和内容
// 这里使用 info.NoteAndContent 接收?
func (this *NoteService) AddNoteAndContentForController(note info.Note, noteContent info.NoteContent, updatedUserId string) info.Note {
	if note.UserId.Hex() != updatedUserId {
		if !shareService.HasUpdateNotebookPerm(note.UserId.Hex(), updatedUserId, note.NotebookId.Hex()) {
			Log("NO AUTH11")
			return info.Note{}
		} else {
			Log("HAS AUTH -----------")
		}
	}
	return this.AddNoteAndContent(note, noteContent, bson.ObjectIdHex(updatedUserId))
}

func (this *NoteService) AddNoteAndContent(note info.Note, noteContent info.NoteContent, myUserId bson.ObjectId) info.Note {
	if note.NoteId.Hex() == "" {
		noteId := bson.NewObjectId()
		note.NoteId = noteId
	}
	noteContent.NoteId = note.NoteId
	if note.UserId != myUserId {
		note = this.AddSharedNote(note, myUserId)
	} else {
		note = this.AddNote(note, false)
	}
	if note.NoteId != "" {
		this.AddNoteContent(noteContent)
	}
	return note
}

func (this *NoteService) AddNoteAndContentApi(note info.Note, noteContent info.NoteContent, myUserId bson.ObjectId) info.Note {
	if note.NoteId.Hex() == "" {
		noteId := bson.NewObjectId()
		note.NoteId = noteId
	}
	noteContent.NoteId = note.NoteId
	if note.UserId != myUserId {
		note = this.AddSharedNote(note, myUserId)
	} else {
		note = this.AddNote(note, true)
	}
	if note.NoteId != "" {
		this.AddNoteContent(noteContent)
	}
	return note
}

// 修改笔记
// 这里没有判断usn
func (this *NoteService) UpdateNote(updatedUserId, noteId string, needUpdate bson.M, usn int) (bool, string, int) {
	ctx := context.Background()

	userId := ""
	newUSN := 0
	needRecountTags := false
	err := updateNote(this.note, ctx, repository.ID(noteId), func(note *model.Note) error {
		// 是否存在
		if note == nil {
			return errcode.NotFound(ctx, "notebookIdNotExists", noteId)
		}

		data := note.MustData(ctx)

		userId = data.UserId.Hex()
		// updatedUserId 要修改userId的note, 此时需要判断是否有修改权限
		if userId != updatedUserId {
			if !shareService.HasUpdatePerm(userId, updatedUserId, noteId) {
				Log("NO AUTH2")
				return errcode.NotFound(ctx, "noAuth", userId, updatedUserId, noteId)
			} else {
				Log("HAS AUTH -----------")
			}
		}

		// 是否已自定义
		if data.IsBlog && data.HasSelfDefined {
			delete(needUpdate, "ImgSrc")
			delete(needUpdate, "Desc")
		}

		// 是否修改了isBlog
		// 也要修改noteContents的IsBlog
		if isBlog, ok := needUpdate["IsBlog"]; ok {
			isBlog2 := isBlog.(bool)
			if data.IsBlog != isBlog2 {
				this.UpdateNoteContentIsBlog(noteId, userId, isBlog2)

				// 重新发布成博客
				if !data.IsBlog {
					needUpdate["PublicTime"] = needUpdate["UpdatedTime"]
				}

				needRecountTags = true
			}
		}

		// 添加tag2
		// TODO 这个tag去掉, 添加tag另外添加, 不要这个
		if tags, ok := needUpdate["Tags"]; ok {
			tagService.AddTagsI(userId, tags)

			// 如果是博客, 标签改了, 那么重新计算
			if data.IsBlog {
				needRecountTags = true
			}
		}

		newUSN = userService.IncrUsn(userId)
		return note.Updete_needdelete_(ctx, needUpdate, updatedUserId, newUSN)
	})
	if err != nil {
		return false, err.Error(), 0
	}

	if needRecountTags {
		// 重新计算tags
		go (func() {
			blogService.ReCountBlogTags(userId)
		})()
	}

	// 重新获取之
	note := this.GetNoteById(noteId)

	hasRecount := false

	// 如果修改了notebookId, 则更新notebookId'count
	// 两方的notebook也要修改
	notebookIdI := needUpdate["NotebookId"]
	if notebookIdI != nil {
		notebookId := notebookIdI.(bson.ObjectId)
		if notebookId != "" {
			notebookService.ReCountNotebookNumberNotes(note.NotebookId.Hex())
			notebookService.ReCountNotebookNumberNotes(notebookId.Hex())
			hasRecount = true
		}
	}

	// 不要多次更新, isTrash = false, = true都要重新统计
	if isTrashI, ok := needUpdate["IsTrash"]; ok {
		// 如果是垃圾, 则删除之共享
		isTrash := isTrashI.(bool)
		if isTrash {
			shareService.DeleteShareNoteAll(noteId, userId)
		}
		if !hasRecount {
			notebookService.ReCountNotebookNumberNotes(note.NotebookId.Hex())
		}
	}

	return true, "", newUSN
}

// 当设置/取消了笔记为博客
func (m *NoteService) UpdateNoteContentIsBlog(noteId, userId string, isBlog bool) {
	ctx := context.Background()

	err := updateContent(m.content, ctx, repository.ID(noteId).WithUser(userId), func(content *model.Content) error {
		return content.SetBlogStatus(ctx, isBlog)
	})
	if err != nil {
		log.G(ctx).WithError(err).Error("设置/取消了笔记为博客失败")
		return
	}

	return
}

// 附件修改, 增加noteIncr
func (this *NoteService) IncrNoteUsn(noteId, userId string) int {
	ctx := context.Background()

	usn := userService.IncrUsn(userId)
	err := updateNote(this.note, ctx, repository.ID(noteId).WithUser(userId), func(note *model.Note) error {
		return note.SetUSN(ctx, usn)
	})
	if err != nil {
		log.G(ctx).WithError(err).Error("IncrNoteUsn 失败")
		return usn
	}
	return usn
}

// 这里要判断权限, 如果userId != updatedUserId, 那么需要判断权限
// [ok] TODO perm还没测 [del]
func (this *NoteService) UpdateNoteTitle(userId, updatedUserId, noteId, title string) bool {
	ctx := context.Background()
	// updatedUserId 要修改userId的note, 此时需要判断是否有修改权限
	if userId != updatedUserId {
		if !shareService.HasUpdatePerm(userId, updatedUserId, noteId) {
			println("NO AUTH")
			return false
		}
	}

	err := updateNote(this.note, ctx, repository.ID(noteId).WithUser(userId), func(note *model.Note) error {
		return note.SetTitle(ctx, updatedUserId, title, userService.IncrUsn(userId))
	})
	if err != nil {
		log.G(ctx).WithError(err).Error("UpdateNoteTitle 失败")
		return false
	}

	return true
}

// 修改笔记本内容
// [ok] TODO perm未测
// hasBeforeUpdateNote 之前是否更新过note其它信息, 如果有更新, usn不用更新
// TODO abstract这里生成
func (this *NoteService) UpdateNoteContent(operateUserID, noteId, noteContent, abstract string, hasBeforeUpdateNote bool, usn int, updatedTime time.Time) (bool, string, int) {
	m := this
	ctx := context.Background()

	err := updateNote(m.note, ctx, repository.ID(noteId), func(note *model.Note) error {
		// 是否已自定义
		if note == nil {
			return errcode.NotFound(ctx, "notExists")
		}

		data := note.MustData(ctx)
		// updatedUserId 要修改userId的note, 此时需要判断是否有修改权限
		userId := data.UserId.Hex()
		if userId != operateUserID && !shareService.HasUpdatePerm(userId, operateUserID, noteId) {
			return errcode.PermissionDenied(ctx, operateUserID, noteId)
		}

		// 如果之前没有修改note其它信息, 那么usn++
		if !hasBeforeUpdateNote {
			// 需要验证
			if usn >= 0 && data.Usn != usn {
				return errcode.DeadlineExceeded(ctx, usn)
			}

			note.NoteMutation.SetField("Usn", userService.IncrUsn(userId))
		}

		return updateContent(m.content, ctx, repository.ContentNoteID(noteId), func(content *model.Content) error {
			if content == nil {
				content = m.content.New(ctx, model.ContentData{
					UserId: bson.ObjectIdHex(userId),
					NoteId: bson.ObjectIdHex(noteId),
				})
			}

			if !data.IsBlog || !data.HasSelfDefined {
				content.SetAbstract(ctx, abstract)
			}

			if err := content.SaveContent(ctx, operateUserID, noteContent, updatedTime); nil != err {
				return err
			}

			// 添加历史记录
			history, err := m.history.Find(ctx, repository.HistoryNoteID(note.MustData(ctx).NoteId.Hex()))
			if err != nil {
				return err
			}

			if history == nil {
				history = m.history.New(ctx, model.HistoryData{
					UserId: bson.ObjectIdHex(userId),
					NoteId: bson.ObjectIdHex(noteId),
				})
			}

			eachHistory := info.EachHistory{
				UpdatedUserId: bson.ObjectIdHex(operateUserID),
				Content:       noteContent,
				UpdatedTime:   time.Now(),
			}

			if err := history.AddHistory(ctx, eachHistory); err != nil {
				return err
			}

			// 更新笔记图片
			data := note.MustData(ctx)
			noteImageService.UpdateNoteImages(userId, noteId, data.ImgSrc, noteContent)

			return m.history.Save(ctx, history)
		})
	})

	if err != nil {
		return false, err.Error(), 0
	}

	return false, "", 0
}

// ?????
// 这种方式太恶心, 改动很大
// 通过content修改笔记的imageIds列表
// src="http://localhost:9000/file/outputImage?fileId=541ae75499c37b6b79000005&noteId=541ae63c19807a4bb9000000"
func (this *NoteService) updateNoteImages(noteId string, content string) bool {
	return true
}

// 更新tags
// [ok] [del]
func (this *NoteService) UpdateTags(noteId string, userId string, tags []string) bool {
	ctx := context.Background()

	err := updateNote(this.note, ctx, repository.ID(noteId).WithUser(userId), func(note *model.Note) error {
		return note.SetTags(ctx, tags, userService.IncrUsn(userId))
	})
	if err != nil {
		log.G(ctx).WithError(err).Error("UpdateTags 失败")
		return false
	}

	return true
}

func (this *NoteService) ToBlog(userId, noteId string, isBlog, isTop bool) bool {
	ctx := context.Background()

	if isTop {
		isBlog = true
	}
	if !isBlog {
		isTop = false
	}

	err := updateNote(this.note, ctx, repository.ID(noteId).WithUser(userId), func(note *model.Note) error {
		usn := userService.IncrUsn(userId)
		if err := note.SetBlogStatus(ctx, isBlog, usn); err != nil {
			return err
		}

		return note.SetTopStatus(ctx, isTop, usn)
	})
	if err != nil {
		log.G(ctx).WithError(err).Error("ToBlog 失败")
		return false
	}

	// 重新计算tags
	go (func() {
		this.UpdateNoteContentIsBlog(noteId, userId, isBlog)

		blogService.ReCountBlogTags(userId)
	})()

	return true
}

// 移动note
// trash, 正常的都可以用
// 1. 要检查下notebookId是否是自己的
// 2. 要判断之前是否是blog, 如果不是, 那么notebook是否是blog?
func (this *NoteService) MoveNote(noteId, notebookId, userId string) info.Note {
	ctx := context.Background()
	if notebookService.IsMyNotebook(notebookId, userId) {
		note := this.GetNote(noteId, userId)
		preNotebookId := note.NotebookId.Hex()

		predicate := repository.ID(noteId).WithUser(userId).WithTrash(false)
		err := updateNote(this.note, ctx, predicate, func(note *model.Note) error {
			usn := userService.IncrUsn(userId)
			if err := note.SetTrash(ctx, false, usn); err != nil {
				return err
			}

			return note.SetBookID(ctx, notebookId, usn)
		})
		if err != nil {
			log.G(ctx).WithError(err).Error("UpdateTags 失败")
			return info.Note{}
		}

		// 更新blog状态
		this.updateToNotebookBlog(noteId, notebookId, userId)

		// recount notebooks' notes number
		notebookService.ReCountNotebookNumberNotes(notebookId)
		// 之前不是trash才统计, trash本不在统计中的
		if !note.IsTrash && preNotebookId != notebookId {
			notebookService.ReCountNotebookNumberNotes(preNotebookId)
		}

		return this.GetNote(noteId, userId)
	}

	return info.Note{}
}

// 如果自己的blog状态是true, 不用改变,
// 否则, 如果notebookId的blog是true, 则改为true之
// 返回blog状态
// move, copy时用
func (this *NoteService) updateToNotebookBlog(noteId, notebookId, userId string) bool {
	ctx := context.Background()
	if this.IsBlog(noteId) {
		return true
	}
	if !notebookService.IsBlog(notebookId) {
		return false
	}

	err := updateNote(this.note, ctx, repository.ID(noteId).WithUser(userId), func(note *model.Note) error {
		return note.SetBlogStatus(ctx, true, userService.IncrUsn(userId))
	})
	if err != nil {
		log.G(ctx).WithError(err).Error("UpdateTags 失败")
		return false
	}

	return true
}

// 判断是否是blog
func (this *NoteService) IsBlog(noteId string) bool {
	m := this
	ctx := context.Background()

	book, err := m.note.Find(ctx, repository.ID(noteId).WithSelect([]string{"IsBlog"}))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return false
	}

	if book == nil {
		return false
	}

	return book.MustData(ctx).IsBlog
}

// 复制note
// 正常的可以用
// 先查, 再新建
// 要检查下notebookId是否是自己的
func (this *NoteService) CopyNote(noteId, notebookId, userId string) info.Note {
	if notebookService.IsMyNotebook(notebookId, userId) {
		note := this.GetNote(noteId, userId)
		noteContent := this.GetNoteContent(noteId, userId)

		// 重新生成noteId
		note.NoteId = bson.NewObjectId()
		note.NotebookId = bson.ObjectIdHex(notebookId)

		noteContent.NoteId = note.NoteId
		note = this.AddNoteAndContent(note, noteContent, note.UserId)

		// 更新blog状态
		isBlog := this.updateToNotebookBlog(note.NoteId.Hex(), notebookId, userId)

		// recount
		notebookService.ReCountNotebookNumberNotes(notebookId)

		note.IsBlog = isBlog

		return note
	}

	return info.Note{}
}

// 复制别人的共享笔记给我
// 将别人可用的图片转为我的图片, 复制图片
func (this *NoteService) CopySharedNote(noteId, notebookId, fromUserId, myUserId string) info.Note {
	// 判断是否共享了给我
	// Log(notebookService.IsMyNotebook(notebookId, myUserId))
	if notebookService.IsMyNotebook(notebookId, myUserId) && shareService.HasReadPerm(fromUserId, myUserId, noteId) {
		note := this.GetNote(noteId, fromUserId)
		if note.NoteId == "" {
			return info.Note{}
		}
		noteContent := this.GetNoteContent(noteId, fromUserId)

		// 重新生成noteId
		note.NoteId = bson.NewObjectId()
		note.NotebookId = bson.ObjectIdHex(notebookId)
		note.UserId = bson.ObjectIdHex(myUserId)
		note.IsTop = false
		note.IsBlog = false // 别人的可能是blog

		note.ImgSrc = "" // 为什么清空, 因为图片需要复制, 先清空

		// content
		noteContent.NoteId = note.NoteId
		noteContent.UserId = note.UserId

		// 复制图片, 把note的图片都copy给我, 且修改noteContent图片路径
		noteContent.Content = noteImageService.CopyNoteImages(noteId, fromUserId, note.NoteId.Hex(), noteContent.Content, myUserId)

		// 复制附件
		attachService.CopyAttachs(noteId, note.NoteId.Hex(), myUserId)

		// 添加之
		note = this.AddNoteAndContent(note, noteContent, note.UserId)

		// 更新blog状态
		isBlog := this.updateToNotebookBlog(note.NoteId.Hex(), notebookId, myUserId)

		// recount
		notebookService.ReCountNotebookNumberNotes(notebookId)

		note.IsBlog = isBlog
		return note
	}

	return info.Note{}
}

// 通过noteId得到notebookId
// shareService call
// [ok]
func (this *NoteService) GetNotebookId(noteId string) bson.ObjectId {
	m := this
	ctx := context.Background()

	book, err := m.note.Find(ctx, repository.ID(noteId).WithSelect([]string{"NotebookId"}))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return ""
	}

	if book == nil {
		return ""
	}

	return book.MustData(ctx).NotebookId
}

//------------------
// 搜索Note, 博客使用了
func (this *NoteService) SearchNote(key, userId string, pageNumber, pageSize int, sortField string, isAsc, isBlog bool) (int, []info.Note) {
	ctx := context.Background()
	skipNum, sortFieldR := parsePageAndSort(pageNumber, pageSize, sortField, isAsc)

	predicate := repository.NoteSearchTitleAndDesc(key, isBlog).WithUser(userId)

	count, err := this.note.Count(ctx, predicate)
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return 0, nil
	}

	notes, err := this.note.FindAll(ctx, predicate.WithSort(sortFieldR).WithPage(skipNum, pageSize))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return 0, nil
	}

	resp := make([]info.Note, 0, len(notes))
	for _, v := range notes {
		resp = append(resp, info.Note(v.MustData(ctx)))
	}

	// 如果 < pageSize 那么搜索content, 且id不在这些id之间的
	if len(resp) < pageSize {
		resp = this.searchNoteFromContent(resp, userId, key, pageSize, sortFieldR, isBlog)
	}

	return count, resp
}

// 搜索noteContents, 补集pageSize个
func (this *NoteService) searchNoteFromContent(notes []info.Note, userId, key string, pageSize int, sortField string, isBlog bool) []info.Note {
	ctx := context.Background()
	var remain = pageSize - len(notes)
	noteIds := make([]string, len(notes))
	for i, note := range notes {
		noteIds[i] = note.NoteId.Hex()
	}

	predicate := repository.NoteSearchContent(noteIds, key, isBlog).WithUser(userId)
	contents, err := this.content.FindAll(ctx, predicate.WithSort(sortField).WithLimit(remain))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return nil
	}

	if len(contents) == 0 {
		return nil
	}

	// 收集ids
	noteIds2 := make([]bson.ObjectId, len(contents))
	for i, content := range contents {
		noteIds2[i] = content.MustData(ctx).NoteId
	}

	// 得到notes
	notes2 := this.ListNotesByNoteIds(noteIds2)

	// 合并之
	// 不能是删除的
	for _, n := range notes2 {
		if !n.IsDeleted && !n.IsTrash {
			// notes = append(notes, notes2...)
			notes = append(notes, n)
		}
	}
	return notes
}

//----------------
// tag搜索
func (this *NoteService) SearchNoteByTags(tags []string, userId string, pageNumber, pageSize int, sortField string, isAsc bool) (int, []info.Note) {
	ctx := context.Background()
	skipNum, sortFieldR := parsePageAndSort(pageNumber, pageSize, sortField, isAsc)

	predicate := repository.NoteSearchTags(tags).WithUser(userId)

	count, err := this.note.Count(ctx, predicate)
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return 0, nil
	}

	notes, err := this.note.FindAll(ctx, predicate.WithSort(sortFieldR).WithPage(skipNum, pageSize))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return 0, nil
	}

	resp := make([]info.Note, 0, len(notes))
	for _, v := range notes {
		resp = append(resp, info.Note(v.MustData(ctx)))
	}

	return count, resp
}

//------------
// 统计
func (this *NoteService) CountNote(userId string) int {
	ctx := context.Background()
	p := repository.All()
	if userId != "" {
		p = p.WithUser(userId)
	}

	count, err := this.note.Count(ctx, p)
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return 0
	}

	return count
}
func (this *NoteService) CountBlog(userId string) int {
	ctx := context.Background()
	p := repository.NoteBlog()
	if userId != "" {
		p = p.WithUser(userId)
	}

	count, err := this.note.Count(ctx, p.WithTrash(false).WithDeleted(false))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记失败")
		return 0
	}

	return count
}

// 通过标签来查询
func (this *NoteService) CountNoteByTag(userId string, tag string) int {
	ctx := context.Background()
	if tag == "" {
		return 0
	}

	count, err := this.note.Count(ctx, repository.NoteTags([]string{tag}).WithUser(userId))
	if err != nil {
		log.G(ctx).WithError(err).Error("查询笔记失败")
	}

	return count
}

// 删除tag
// 返回所有note的Usn
func (this *NoteService) UpdateNoteToDeleteTag(userId string, targetTag string) map[string]int {
	ctx := context.Background()

	ret := map[string]int{}
	err := updateNotes(this.note, ctx, repository.NoteTags([]string{targetTag}).WithUser(userId), func(notes []*model.Note) error {
		for _, note := range notes {
			data := note.MustData(ctx)
			tags := data.Tags
			if tags == nil {
				continue
			}
			for i, tag := range tags {
				if tag == targetTag {
					tags = append(tags[:i], tags[i+1:]...)
					break
				}
			}
			usn := userService.IncrUsn(userId)
			note.SetTags(ctx, tags, usn)
			ret[data.NoteId.Hex()] = usn
		}

		return nil
	})
	if err != nil {
		log.G(ctx).WithError(err).Error("删除tag失败")
	}

	return ret
}

// api

// 得到笔记的内容, 此时将笔记内的链接转成标准的Leanote Url
// 将笔记的图片, 附件链接转换成 site.url/file/getImage?fileId=xxx,  site.url/file/getAttach?fileId=xxxx
func (this *NoteService) FixContentBad(content string, isMarkdown bool) string {
	baseUrl := configService.GetSiteUrl()

	baseUrlPattern := baseUrl

	// 避免https的url
	if baseUrl[0:8] == "https://" {
		baseUrlPattern = strings.Replace(baseUrl, "https://", "https*://", 1)
	} else {
		baseUrlPattern = strings.Replace(baseUrl, "http://", "https*://", 1)
	}

	patterns := []map[string]string{
		map[string]string{"src": "src", "middle": "/file/outputImage", "param": "fileId", "to": "getImage?fileId="},
		map[string]string{"src": "href", "middle": "/attach/download", "param": "attachId", "to": "getAttach?fileId="},
		// 该链接已失效, 不再支持
		map[string]string{"src": "href", "middle": "/attach/downloadAll", "param": "noteId", "to": "getAllAttachs?noteId="},
	}

	for _, eachPattern := range patterns {

		if !isMarkdown {

			// 富文本处理

			// <img src="http://leanote.com/file/outputImage?fileId=5503537b38f4111dcb0000d1">
			// <a href="http://leanote.com/attach/download?attachId=5504243a38f4111dcb00017d"></a>

			var reg *regexp.Regexp
			if eachPattern["src"] == "src" {
				reg, _ = regexp.Compile("<img(?:[^>]+?)(" + eachPattern["src"] + `=['"]*` + baseUrlPattern + eachPattern["middle"] + `\?` + eachPattern["param"] + `=([a-z0-9A-Z]{24})["']*)[^>]*>`)
			} else {
				reg, _ = regexp.Compile("<a(?:[^>]+?)(" + eachPattern["src"] + `=['"]*` + baseUrlPattern + eachPattern["middle"] + `\?` + eachPattern["param"] + `=([a-z0-9A-Z]{24})["']*)[^>]*>`)
			}

			finds := reg.FindAllStringSubmatch(content, -1) // 查找所有的

			for _, eachFind := range finds {
				if len(eachFind) == 3 {
					// 这一行会非常慢!, content是全部的内容, 多次replace导致
					content = strings.Replace(content,
						eachFind[1],
						eachPattern["src"]+"=\""+baseUrl+"/api/file/"+eachPattern["to"]+eachFind[2]+"\"",
						1)
				}
			}
		} else {

			// markdown处理
			// ![](http://leanote.com/file/outputImage?fileId=5503537b38f4111dcb0000d1)
			// [selection 2.html](http://leanote.com/attach/download?attachId=5504262638f4111dcb00017f)
			// [all.tar.gz](http://leanote.com/attach/downloadAll?noteId=5503b57d59f81b4eb4000000)

			pre := "!"                        // 默认图片
			if eachPattern["src"] == "href" { // 是attach
				pre = ""
			}

			regImageMarkdown, _ := regexp.Compile(pre + `\[([^]]*?)\]\(` + baseUrlPattern + eachPattern["middle"] + `\?` + eachPattern["param"] + `=([a-z0-9A-Z]{24})\)`)
			findsImageMarkdown := regImageMarkdown.FindAllStringSubmatch(content, -1) // 查找所有的
			// [[![](http://leanote.com/file/outputImage?fileId=5503537b38f4111dcb0000d1) 5503537b38f4111dcb0000d1] [![你好啊, 我很好, 为什么?](http://leanote.com/file/outputImage?fileId=5503537b38f4111dcb0000d1) 5503537b38f4111dcb0000d1]]
			for _, eachFind := range findsImageMarkdown {
				// [![你好啊, 我很好, 为什么?](http://leanote.com/file/outputImage?fileId=5503537b38f4111dcb0000d1) 你好啊, 我很好, 为什么? 5503537b38f4111dcb0000d1]
				if len(eachFind) == 3 {
					content = strings.Replace(content, eachFind[0], pre+"["+eachFind[1]+"]("+baseUrl+"/api/file/"+eachPattern["to"]+eachFind[2]+")", 1)
				}
			}
		}
	}

	return content
}

// 得到笔记的内容, 此时将笔记内的链接转成标准的Leanote Url
// 将笔记的图片, 附件链接转换成 site.url/file/getImage?fileId=xxx,  site.url/file/getAttach?fileId=xxxx
// 性能更好, 5倍的差距
func (this *NoteService) FixContent(content string, isMarkdown bool) string {
	baseUrl := configService.GetSiteUrl()

	baseUrlPattern := baseUrl

	// 避免https的url
	if baseUrl[0:8] == "https://" {
		baseUrlPattern = strings.Replace(baseUrl, "https://", "https*://", 1)
	} else {
		baseUrlPattern = strings.Replace(baseUrl, "http://", "https*://", 1)
	}
	baseUrlPattern = "(?:" + baseUrlPattern + ")*"

	Log(baseUrlPattern)

	patterns := []map[string]string{
		map[string]string{"src": "src", "middle": "/api/file/getImage", "param": "fileId", "to": "getImage?fileId="},
		map[string]string{"src": "src", "middle": "/file/outputImage", "param": "fileId", "to": "getImage?fileId="},

		map[string]string{"src": "href", "middle": "/attach/download", "param": "attachId", "to": "getAttach?fileId="},
		map[string]string{"src": "href", "middle": "/api/file/getAtach", "param": "fileId", "to": "getAttach?fileId="},

		// 该链接已失效, 不再支持
		// map[string]string{"src": "href", "middle": "/attach/downloadAll", "param": "noteId", "to": "getAllAttachs?noteId="},
	}

	for _, eachPattern := range patterns {

		if !isMarkdown {

			// 富文本处理

			// <img src="http://leanote.com/file/outputImage?fileId=5503537b38f4111dcb0000d1">
			// <a href="http://leanote.com/attach/download?attachId=5504243a38f4111dcb00017d"></a>

			var reg *regexp.Regexp
			var reg2 *regexp.Regexp
			if eachPattern["src"] == "src" {
				reg, _ = regexp.Compile("<img(?:[^>]+?)(?:" + eachPattern["src"] + `=['"]*` + baseUrlPattern + eachPattern["middle"] + `\?` + eachPattern["param"] + `=(?:[a-z0-9A-Z]{24})["']*)[^>]*>`)
				reg2, _ = regexp.Compile("<img(?:[^>]+?)(" + eachPattern["src"] + `=['"]*` + baseUrlPattern + eachPattern["middle"] + `\?` + eachPattern["param"] + `=([a-z0-9A-Z]{24})["']*)[^>]*>`)
			} else {
				reg, _ = regexp.Compile("<a(?:[^>]+?)(?:" + eachPattern["src"] + `=['"]*` + baseUrlPattern + eachPattern["middle"] + `\?` + eachPattern["param"] + `=(?:[a-z0-9A-Z]{24})["']*)[^>]*>`)
				reg2, _ = regexp.Compile("<a(?:[^>]+?)(" + eachPattern["src"] + `=['"]*` + baseUrlPattern + eachPattern["middle"] + `\?` + eachPattern["param"] + `=([a-z0-9A-Z]{24})["']*)[^>]*>`)
			}

			// Log(reg2)

			content = reg.ReplaceAllStringFunc(content, func(str string) string {
				// str=这样的
				// <img src="http://localhost:9000/file/outputImage?fileId=563d706e99c37b48e0000001" alt="" data-mce-src="http://localhost:9000/file/outputImage?fileId=563d706e99c37b48e0000002">

				eachFind := reg2.FindStringSubmatch(str)
				str = strings.Replace(str,
					eachFind[1],
					eachPattern["src"]+"=\""+baseUrl+"/api/file/"+eachPattern["to"]+eachFind[2]+"\"",
					1)

				// fmt.Println(str)
				return str
			})
			/*
				finds := reg.FindAllStringSubmatch(content, -1) // 查找所有的

				for _, eachFind := range finds {
					if len(eachFind) == 3 {
						// 这一行会非常慢!, content是全部的内容, 多次replace导致
						content = strings.Replace(content,
							eachFind[1],
							eachPattern["src"]+"=\""+baseUrl+"/api/file/"+eachPattern["to"]+eachFind[2]+"\"",
							1)
					}
				}
			*/
		} else {

			// markdown处理
			// ![](http://leanote.com/file/outputImage?fileId=5503537b38f4111dcb0000d1)
			// [selection 2.html](http://leanote.com/attach/download?attachId=5504262638f4111dcb00017f)
			// [all.tar.gz](http://leanote.com/attach/downloadAll?noteId=5503b57d59f81b4eb4000000)

			pre := "!"                        // 默认图片
			if eachPattern["src"] == "href" { // 是attach
				pre = ""
			}

			regImageMarkdown, _ := regexp.Compile(pre + `\[(?:[^]]*?)\]\(` + baseUrlPattern + eachPattern["middle"] + `\?` + eachPattern["param"] + `=(?:[a-z0-9A-Z]{24})\)`)
			regImageMarkdown2, _ := regexp.Compile(pre + `\[([^]]*?)\]\(` + baseUrlPattern + eachPattern["middle"] + `\?` + eachPattern["param"] + `=([a-z0-9A-Z]{24})\)`)

			content = regImageMarkdown.ReplaceAllStringFunc(content, func(str string) string {
				// str=这样的
				// <img src="http://localhost:9000/file/outputImage?fileId=563d706e99c37b48e0000001" alt="" data-mce-src="http://localhost:9000/file/outputImage?fileId=563d706e99c37b48e0000002">

				eachFind := regImageMarkdown2.FindStringSubmatch(str)
				str = strings.Replace(str, eachFind[0], pre+"["+eachFind[1]+"]("+baseUrl+"/api/file/"+eachPattern["to"]+eachFind[2]+")", 1)

				// fmt.Println(str)
				return str
			})

			/*
				findsImageMarkdown := regImageMarkdown.FindAllStringSubmatch(content, -1) // 查找所有的
				// [[![](http://leanote.com/file/outputImage?fileId=5503537b38f4111dcb0000d1) 5503537b38f4111dcb0000d1] [![你好啊, 我很好, 为什么?](http://leanote.com/file/outputImage?fileId=5503537b38f4111dcb0000d1) 5503537b38f4111dcb0000d1]]
				for _, eachFind := range findsImageMarkdown {
					// [![你好啊, 我很好, 为什么?](http://leanote.com/file/outputImage?fileId=5503537b38f4111dcb0000d1) 你好啊, 我很好, 为什么? 5503537b38f4111dcb0000d1]
					if len(eachFind) == 3 {
						content = strings.Replace(content, eachFind[0], pre+"["+eachFind[1]+"]("+baseUrl+"/api/file/"+eachPattern["to"]+eachFind[2]+")", 1)
					}
				}
			*/
		}
	}

	return content
}
