package service

import (
	"context"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/coocn-cn/leanote/app/db"
	"github.com/coocn-cn/leanote/app/lea"
	"github.com/coocn-cn/leanote/app/note/model"
	"github.com/coocn-cn/leanote/app/note/repository"
	tag_model "github.com/coocn-cn/leanote/app/tag/model"
	tag_repo "github.com/coocn-cn/leanote/app/tag/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//----------------
// service 公用方法

// 分页, 排序处理
func parsePageAndSort(pageNumber, pageSize int, sortField string, isAsc bool) (skipNum int, sortFieldR string) {
	skipNum = (pageNumber - 1) * pageSize
	if sortField == "" {
		sortField = "UpdatedTime"
	}
	if !isAsc {
		sortFieldR = "-" + sortField
	} else {
		sortFieldR = sortField
	}
	return
}

// 将name=val的val进行encoding
func decodeValue(val string) string {
	v, _ := url.ParseQuery("a=" + val)
	return v.Get("a")
}

func encodeValue(val string) string {
	if val == "" {
		return val
	}
	v := url.Values{}
	v.Set("", val)
	return v.Encode()[1:]
}

// 添加笔记时通过title得到urlTitle
func fixUrlTitle(urlTitle string) string {
	if urlTitle != "" {
		// 把特殊字段给替换掉
		//		str := `life "%&()+,/:;<>=?@\|`
		reg, _ := regexp.Compile("/|#|\\$|!|\\^|\\*|'| |\"|%|&|\\(|\\)|\\+|\\,|/|:|;|<|>|=|\\?|@|\\||\\\\")
		urlTitle = reg.ReplaceAllString(urlTitle, "-")
		urlTitle = strings.Trim(urlTitle, "-") // 左右单独的-去掉
		// 把空格替换成-
		//		urlTitle = strings.Replace(urlTitle, " ", "-", -1)
		for strings.Index(urlTitle, "--") >= 0 { // 防止出现连续的--
			urlTitle = strings.Replace(urlTitle, "--", "-", -1)
		}
		return encodeValue(urlTitle)
	}
	return urlTitle
}

func getUniqueUrlTitle(userId string, urlTitle string, types string, padding int) string {
	urlTitle2 := urlTitle

	// 判断urlTitle是不是过长, 过长则截断, 300
	// 不然生成index有问题
	// it will not index a single field with more than 1024 bytes.
	// If you're indexing a field that is 2.5MB, it's not really indexing it, it's being skipped.
	if len(urlTitle2) > 320 {
		urlTitle2 = urlTitle2[:300] // 为什么要少些, 因为怕无限循环, 因为把padding截了
	}

	if padding > 1 {
		urlTitle2 = urlTitle + "-" + strconv.Itoa(padding)
	}
	userIdO := bson.ObjectIdHex(userId)

	var collection *mgo.Collection
	if types == "note" {
		collection = db.Notes
	} else if types == "notebook" {
		collection = db.Notebooks
	} else if types == "single" {
		collection = db.BlogSingles
	}
	for db.Has(collection, bson.M{"UserId": userIdO, "UrlTitle": urlTitle2}) { // 用户下唯一
		padding++
		urlTitle2 = urlTitle + "-" + strconv.Itoa(padding)
	}

	return urlTitle2
}

// 截取id 24位变成12位
// 先md5, 再取12位
func subIdHalf(id string) string {
	idMd5 := lea.Md5(id)
	return idMd5[:12]
}

// types == note,notebook,single
// id noteId, notebookId, singleId 当title没的时候才有用, 用它来替换
func GetUrTitle(userId string, title string, types string, id string) string {
	urlTitle := strings.Trim(title, " ")
	if urlTitle == "" {
		if id == "" {
			urlTitle = "Untitled-" + userId
		} else {
			urlTitle = subIdHalf(id)
		}
		// 不允许title是ObjectId
	} else if bson.IsObjectIdHex(title) {
		urlTitle = subIdHalf(id)
	}

	urlTitle = fixUrlTitle(urlTitle)
	return getUniqueUrlTitle(userId, urlTitle, types, 1)
}

func updateNote(repo repository.NoteRepository, ctx context.Context, predicate repository.Predicater, f func(*model.Note) error) error {
	model, err := repo.Find(ctx, predicate)
	if err != nil {
		return err
	}

	if err := f(model); err != nil {
		return err
	}

	if model == nil {
		return nil
	}

	return repo.Save(ctx, model)
}

func updateNotes(repo repository.NoteRepository, ctx context.Context, predicate repository.Predicater, f func([]*model.Note) error) error {
	models, err := repo.FindAll(ctx, predicate)
	if err != nil {
		return err
	}

	if err := f(models); err != nil {
		return err
	}

	if models == nil {
		return nil
	}

	return repo.Save(ctx, models...)
}

func updateContent(repo repository.ContentRepository, ctx context.Context, predicate repository.Predicater, f func(*model.Content) error) error {
	model, err := repo.Find(ctx, predicate)
	if err != nil {
		return err
	}

	if err := f(model); err != nil {
		return err
	}

	if model == nil {
		return nil
	}

	return repo.Save(ctx, model)
}

func updateContents(repo repository.ContentRepository, ctx context.Context, predicate repository.Predicater, f func([]*model.Content) error) error {
	models, err := repo.FindAll(ctx, predicate)
	if err != nil {
		return err
	}

	if err := f(models); err != nil {
		return err
	}

	if models == nil {
		return nil
	}

	return repo.Save(ctx, models...)
}

func updateHistory(repo repository.HistoryRepository, ctx context.Context, predicate repository.Predicater, f func(*model.History) error) error {
	model, err := repo.Find(ctx, predicate)
	if err != nil {
		return err
	}

	if err := f(model); err != nil {
		return err
	}

	if model == nil {
		return nil
	}

	return repo.Save(ctx, model)
}

func updateHistorys(repo repository.HistoryRepository, ctx context.Context, predicate repository.Predicater, f func([]*model.History) error) error {
	models, err := repo.FindAll(ctx, predicate)
	if err != nil {
		return err
	}

	if err := f(models); err != nil {
		return err
	}

	if models == nil {
		return nil
	}

	return repo.Save(ctx, models...)
}

func updateNoteTag(repo repository.TagRepository, ctx context.Context, predicate repository.Predicater, f func(*model.Tag) error) error {
	model, err := repo.Find(ctx, predicate)
	if err != nil {
		return err
	}

	if err := f(model); err != nil {
		return err
	}

	if model == nil {
		return nil
	}

	return repo.Save(ctx, model)
}

func updateNoteTags(repo repository.TagRepository, ctx context.Context, predicate repository.Predicater, f func([]*model.Tag) error) error {
	models, err := repo.FindAll(ctx, predicate)
	if err != nil {
		return err
	}

	if err := f(models); err != nil {
		return err
	}

	if models == nil {
		return nil
	}

	return repo.Save(ctx, models...)
}

func updateTag(repo tag_repo.TagRepository, ctx context.Context, predicate repository.Predicater, f func(*tag_model.Tag) error) error {
	model, err := repo.Find(ctx, predicate)
	if err != nil {
		return err
	}

	if err := f(model); err != nil {
		return err
	}

	if model == nil {
		return nil
	}

	return repo.Save(ctx, model)
}

func updateTags(repo tag_repo.TagRepository, ctx context.Context, predicate repository.Predicater, f func([]*tag_model.Tag) error) error {
	models, err := repo.FindAll(ctx, predicate)
	if err != nil {
		return err
	}

	if err := f(models); err != nil {
		return err
	}

	if models == nil {
		return nil
	}

	return repo.Save(ctx, models...)
}
