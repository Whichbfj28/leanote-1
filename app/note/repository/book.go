package repository

import (
	"context"

	"github.com/coocn-cn/leanote/app/note/model"
)

//xgo:generate gex implgen -destination git/$GOFILE -package git -impl_names BookRepository=book -source $GOFILE
//go:generate gex implgen -destination mongo/$GOFILE -package mongo -impl_names BookRepository=book -source $GOFILE
//go:generate gex mockgen -destination mock/$GOFILE -package mock -source $GOFILE

type BookRepository interface {
	// New is 初始化一个新的领域对象
	New(ctx context.Context, data model.BookData) *model.Book

	// Find is 加载一个符合 Predicater 条件的领域对象
	Find(ctx context.Context, predicate Predicater) (*model.Book, error)
	// Count is 获取所有符合 Predicater 条件的领域对象的个数
	Count(ctx context.Context, predicate Predicater) (int, error)
	// FindAll is 加载所有符合 Predicater 条件的领域对象
	FindAll(ctx context.Context, predicate Predicater) ([]*model.Book, error)

	// Save is 保存领域对象到仓储
	Save(ctx context.Context, models ...*model.Book) error

	// Delete is 删除仓储中的领域对象
	Delete(ctx context.Context, models ...*model.Book) error
	// DeleteID is 通过ID删除仓储中的领域对象
	DeleteID(ctx context.Context, ids ...string) error
}

// BookBookID is 查询条件 - 按用户, BookID, Trash状态和删除状态过滤
func BookBookID(bookID string) *PredicateBuild {
	data := map[string]interface{}{
		"bookID": bookID,
	}

	return NewPredicate("BookBookID", data)
}

// BookParentID is 查询条件 - 按用户, ParentID 和删除状态过滤
func BookParentID(parentID string) *PredicateBuild {
	data := map[string]interface{}{
		"parentID": parentID,
	}

	return NewPredicate("BookParentID", data)
}

// BookNexts is 查询条件 - 获取比指定Usn更新的记录
func BookNexts(usn int) *PredicateBuild {
	data := map[string]interface{}{
		"usn": usn,
	}

	return NewPredicate("BookNexts", data).WithSort("Usn")
}

// BookURLTitle is 查询条件 - 按用户和 URLTitle 过滤
func BookURLTitle(urlTitle string) *PredicateBuild {
	data := map[string]string{
		"urlTitle": urlTitle,
	}

	return NewPredicate("BookURLTitle", data)
}
