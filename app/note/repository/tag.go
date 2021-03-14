package repository

import (
	"context"

	"github.com/coocn-cn/leanote/app/note/model"
)

//go:generate gex implgen -destination mongo/$GOFILE -package mongo -impl_names TagRepository=tag -source $GOFILE
//go:generate gex mockgen -destination mock/$GOFILE -package mock -source $GOFILE

type TagRepository interface {
	// New is 初始化一个新的领域对象
	New(ctx context.Context, data model.TagData) *model.Tag

	// Find is 加载一个符合 Predicater 条件的领域对象
	Find(ctx context.Context, predicate Predicater) (*model.Tag, error)
	// Count is 获取所有符合 Predicater 条件的领域对象的个数
	Count(ctx context.Context, predicate Predicater) (int, error)
	// FindAll is 加载所有符合 Predicater 条件的领域对象
	FindAll(ctx context.Context, predicate Predicater) ([]*model.Tag, error)

	// Save is 保存领域对象到仓储
	Save(ctx context.Context, models ...*model.Tag) error

	// Delete is 删除仓储中的领域对象
	Delete(ctx context.Context, models ...*model.Tag) error
	// DeleteID is 通过ID删除仓储中的领域对象
	DeleteID(ctx context.Context, ids ...uint64) error
}

// TagTag is 查询条件 - 获取记事的历史纪录
func TagTag(tag string) *PredicateBuild {
	data := map[string]string{
		"tag": tag,
	}

	return NewPredicate("TagTag", data)
}
