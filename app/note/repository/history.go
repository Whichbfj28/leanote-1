package repository

import (
	"context"

	"github.com/coocn-cn/leanote/app/note/model"
)

//xgo:generate gex implgen -destination git/$GOFILE -package git -impl_names HistoryRepository=content -source $GOFILE
//go:generate gex implgen -destination mongo/$GOFILE -package mongo -impl_names HistoryRepository=content -source $GOFILE
//go:generate gex mockgen -destination mock/$GOFILE -package mock -source $GOFILE

type HistoryRepository interface {
	// New is 初始化一个新的领域对象
	New(ctx context.Context, data model.HistoryData) *model.History

	// Find is 加载一个符合 Predicater 条件的领域对象
	Find(ctx context.Context, predicate Predicater) (*model.History, error)
	// Count is 获取所有符合 Predicater 条件的领域对象的个数
	Count(ctx context.Context, predicate Predicater) (int, error)
	// FindAll is 加载所有符合 Predicater 条件的领域对象
	FindAll(ctx context.Context, predicate Predicater) ([]*model.History, error)

	// Save is 保存领域对象到仓储
	Save(ctx context.Context, models ...*model.History) error

	// Delete is 删除仓储中的领域对象
	Delete(ctx context.Context, models ...*model.History) error
	// DeleteID is 通过ID删除仓储中的领域对象
	DeleteID(ctx context.Context, ids ...uint64) error
}

// HistoryNoteID is 查询条件 - 获取记事的历史纪录
func HistoryNoteID(noteID string) *PredicateBuild {
	data := map[string]string{
		"noteID": noteID,
	}

	return NewPredicate("HistoryNoteID", data)
}
