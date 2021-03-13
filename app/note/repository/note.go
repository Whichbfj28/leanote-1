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

// NoteIDs is 查询条件 - 从仓储加载一个领域对象
func NoteID(id string) Predicater {
	data := map[string]string{
		"id": id,
	}

	return &basePredicate{name: "NoteID", data: data}
}

// NoteIDs is 查询条件 - 从仓储加载一个没有被删除的领域对象
func NoteIDAndNotDelete(id string) Predicater {
	data := map[string]string{
		"id": id,
	}

	return &basePredicate{name: "NoteIDAndNotDelete", data: data}
}

// NoteIDs is 查询条件 - 从仓储加载多个领域对象
func NoteIDs(ids []string) Predicater {
	data := map[string][]string{
		"ids": ids,
	}

	return &basePredicate{name: "NoteIDs", data: data}
}

// NoteBookID is 查询条件 - 获取笔记本里的所有笔记
func NoteBookID(id string) Predicater {
	data := map[string]string{
		"bookID": id,
	}

	return &basePredicate{name: "NoteBookID", data: data}
}
