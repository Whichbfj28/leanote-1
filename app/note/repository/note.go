package repository

import (
	"context"

	"github.com/coocn-cn/leanote/app/note/model"
)

//xgo:generate gex implgen -destination git/$GOFILE -package git -impl_names NoteRepository=note -source $GOFILE
//go:generate gex implgen -destination mongo/$GOFILE -package mongo -impl_names NoteRepository=note -source $GOFILE
//go:generate gex mockgen -destination mock/$GOFILE -package mock -source $GOFILE

type NoteRepository interface {
	// New is 初始化一个新的领域对象
	New(ctx context.Context, data model.NoteData) *model.Note

	// Find is 加载一个符合 Predicater 条件的领域对象
	Find(ctx context.Context, predicate Predicater) (*model.Note, error)
	// Count is 获取所有符合 Predicater 条件的领域对象的个数
	Count(ctx context.Context, predicate Predicater) (int, error)
	// FindAll is 加载所有符合 Predicater 条件的领域对象
	FindAll(ctx context.Context, predicate Predicater) ([]*model.Note, error)

	// Save is 保存领域对象到仓储
	Save(ctx context.Context, models ...*model.Note) error

	// Delete is 删除仓储中的领域对象
	Delete(ctx context.Context, models ...*model.Note) error
	// DeleteID is 通过ID删除仓储中的领域对象
	DeleteID(ctx context.Context, ids ...uint64) error
}

// NoteBookID is 查询条件 - 获取笔记本里的所有笔记
func NoteBookID(id string) *PredicateBuild {
	data := map[string]string{
		"bookID": id,
	}

	return NewPredicate("NoteBookID", data)
}

// NoteIDAndTags is 查询条件 - 获取拥有指定Tags的笔记
func NoteIDAndBlog(id string, blog bool) *PredicateBuild {
	data := map[string]interface{}{
		"id":   id,
		"blog": blog,
	}

	return NewPredicate("NoteIDAndBlog", data)
}

// NoteSrc is 查询条件 - 获取拥有指定Tags的笔记
func NoteSrc(src string) *PredicateBuild {
	data := map[string]interface{}{
		"src": src,
	}

	return NewPredicate("NoteSrc", data)
}

// NoteBlog is 查询条件 - 获取拥有指定Tags的笔记
func NoteBlog() *PredicateBuild {
	data := true

	return NewPredicate("NoteBlog", data)
}

// NoteTags is 查询条件 - 获取拥有指定Tags的笔记
func NoteTags(tags []string) *PredicateBuild {
	data := map[string]interface{}{
		"tags": tags,
	}

	return NewPredicate("NoteTags", data)
}

// NoteNexts is 查询条件 - 获取比指定Usn更新的记录
func NoteNexts(usn int) *PredicateBuild {
	data := map[string]interface{}{
		"usn": usn,
	}

	return NewPredicate("NoteNexts", data).WithSort("Usn")
}

// NoteBookIDAndBlog is 查询条件 - 获取笔记本里的所有笔记
func NoteBookIDAndBlog(bookID string, blog bool) *PredicateBuild {
	data := map[string]interface{}{
		"bookID": bookID,
		"blog":   blog,
	}

	return NewPredicate("NoteBookIDAndBlog", data)
}

// NoteSearchTags is 查询条件 - 获取拥有指定Tags的笔记
func NoteSearchTags(tags []string) *PredicateBuild {
	data := map[string]interface{}{
		"tags": tags,
	}

	return NewPredicate("NoteSearchTags", data)
}

// NoteSearchTitleAndDesc is 查询条件 - 获取拥有指定Tags的笔记
func NoteSearchTitleAndDesc(query string, searchBlog bool) *PredicateBuild {
	data := map[string]interface{}{
		"query": query,
		"blog":  searchBlog,
	}

	return NewPredicate("NoteSearchTitleAndDesc", data)
}

// NoteSearchContent is 查询条件 - 获取拥有指定Tags的笔记
func NoteSearchContent(excludeIDs []string, query string, searchBlog bool) *PredicateBuild {
	data := map[string]interface{}{
		"excludeIDs": excludeIDs,
		"query":      query,
		"blog":       searchBlog,
	}

	return NewPredicate("NoteSearchContent", data)
}
