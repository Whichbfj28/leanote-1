package repository

import (
	"context"

	"github.com/coocn-cn/leanote/app/note/model"
)

//go:generate gex implgen -destination git/$GOFILE -package git -impl_names HistoryRepository=history -source $GOFILE
//go:generate gex implgen -destination mongo/$GOFILE -package mongo -impl_names HistoryRepository=history -source $GOFILE
//go:generate gex mockgen -destination mock/$GOFILE -package mock -source $GOFILE

type HistoryRepository interface {
	// New is 初始化一个新的领域对象
	New(ctx context.Context, data model.HistoryData) *model.History

	// Load is 从仓储加载一个领域对象
	Load(ctx context.Context, id uint64) (*model.History, error)
	// Find is 加载一个符合 Predicater 条件的领域对象
	Find(ctx context.Context, predicates Predicater) (*model.History, error)
	// FindAll is 加载所有符合 Predicater 条件的领域对象
	FindAll(ctx context.Context, predicates Predicater) ([]*model.History, error)

	// Save is 保存领域对象到仓储
	Save(ctx context.Context, model *model.History) error

	// Delete is 删除仓储中的领域对象
	Delete(ctx context.Context, model *model.History) error
	// DeleteID is 通过ID删除仓储中的领域对象
	DeleteID(ctx context.Context, id uint64) error
}

// HistoryUserAndID is 查询条件 - 按用户和ID过滤
func HistoryUserAndID(userID string, id string) Predicater {
	data := map[string]string{
		"userID": userID,
		"id":     id,
	}

	return &basePredicate{name: "HistoryUserAndID", data: data}
}
