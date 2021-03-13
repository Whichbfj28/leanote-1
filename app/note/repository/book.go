package repository

import (
	"context"

	"github.com/coocn-cn/leanote/app/note/model"
)

//xgo:generate gex implgen -destination git/$GOFILE -package git -impl_names BookRepository=Book -source $GOFILE
//go:generate gex implgen -destination mongo/$GOFILE -package mongo -impl_names BookRepository=Book -source $GOFILE
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

// BookIDs is 查询条件 - 从仓储加载一个领域对象
func BookID(id string) Predicater {
	data := map[string]string{
		"id": id,
	}

	return &basePredicate{name: "BookID", data: data}
}

// BookIDs is 查询条件 - 从仓储加载多个领域对象
func BookIDs(ids []string) Predicater {
	data := map[string][]string{
		"ids": ids,
	}

	return &basePredicate{name: "BookIDs", data: data}
}

// BookIDs is 查询条件 - 加载多个ID
func BookUserAndIDs(userID string, ids []string) Predicater {
	data := map[string][]string{
		"userID": []string{userID},
		"ids":    ids,
	}

	return &basePredicate{name: "BookUserAndIDs", data: data}
}

// BookUserAndID is 查询条件 - 按用户和 ID 过滤
func BookUserAndID(userID string, id string) Predicater {
	data := map[string]string{
		"userID": userID,
		"id":     id,
	}

	return &basePredicate{name: "BookUserAndID", data: data}
}

// BookUserAndNotDelete is 查询条件 - 获取用户没有被删除的笔记本列表
func BookUserAndNotDelete(userID string) Predicater {
	data := map[string]string{
		"userID": userID,
	}

	return &basePredicate{name: "BookUserAndNotDelete", data: data}
}

// BookUserAndParentIDAndDelete is 查询条件 - 按用户, ParentID 和删除状态过滤
func BookUserAndParentIDAndDelete(userID, parentID string, delete bool) Predicater {
	data := map[string]interface{}{
		"userID":   userID,
		"parentID": parentID,
		"delete":   delete,
	}

	return &basePredicate{name: "BookUserAndParentIDAndDelete", data: data}
}

// BookUserAndBookIDAndTrashAndDelete is 查询条件 - 按用户, BookID, Trash状态和删除状态过滤
func BookUserAndBookIDAndTrashAndDelete(userID, bookID string, trash, delete bool) Predicater {
	data := map[string]interface{}{
		"userID": userID,
		"bookID": bookID,
		"trash":  trash,
		"delete": delete,
	}

	return &basePredicate{name: "BookUserAndBookIDAndTrashAndDelete", data: data}
}

// BookUserAndURLTitle is 查询条件 - 按用户和 URLTitle 过滤
func BookUserAndURLTitle(userID string, urlTitle string) Predicater {
	data := map[string]string{
		"userID":   userID,
		"urlTitle": urlTitle,
	}

	return &basePredicate{name: "BookUserAndURLTitle", data: data}
}

// BookUSNNextBooks is 查询条件 - 获取比指定Usn更新的记录
func BookUSNNextBooks(userID string, usn int, count int) Predicater {
	data := map[string]interface{}{
		"userID": userID,
		"usn":    usn,
		"limit":  count,
	}

	return &basePredicate{name: "BookUSNNextBooks", data: data}
}
