package model

import (
	"context"
	"time"

	"github.com/coocn-cn/leanote/app/info"
	"github.com/coocn-cn/leanote/pkg/errcode"
	"gopkg.in/mgo.v2/bson"
)

type BookData info.Notebook

type BookMutation interface {
	Data(context.Context) (BookData, error)
	SetField(name string, value interface{})
	SetFields(fields map[string]interface{})
}

type Book struct {
	BookMutation
}

func NewBook(mut BookMutation) *Book {
	return &Book{BookMutation: mut}
}

func (m *Book) MustData(ctx context.Context) BookData {
	data, err := m.Data(ctx)
	if err != nil {
		panic(err)
	}

	return data
}

// 查看是否有子notebook
// 先查看该notebookId下是否有notes, 没有则删除
func (m *Book) SoftDelete(ctx context.Context, newUSN int) error {
	mut := m

	mut.SetField("Usn", newUSN)
	mut.SetField("IsDeleted", true)

	return nil
}

func (m *Book) UpdateTitle(ctx context.Context, userId, title string, newUSN int) error {
	mut := m

	mut.SetField("Usn", newUSN)
	mut.SetField("Title", title)
	mut.SetField("UpdatedTime", time.Now())

	return nil
}

func (m *Book) UpdateNotebookApi(ctx context.Context, userId, title, parentNotebookId string, seq, usn, newUSN int) error {
	mut := m
	data, err := m.Data(ctx)
	if err != nil {
		return err
	}

	// 先判断usn是否和数据库的一样, 如果不一样, 则冲突, 不保存
	if data.Usn != usn {
		return errcode.DeadlineExceeded(ctx, "conflict", data.Usn, usn)
	}

	if err := m.UpdateTitle(ctx, userId, title, newUSN); err != nil {
		return err
	}

	mut.SetField("Seq", seq)
	if parentNotebookId != "" && bson.IsObjectIdHex(parentNotebookId) {
		mut.SetField("ParentNotebookId", bson.ObjectIdHex(parentNotebookId))
	} else {
		mut.SetField("ParentNotebookId", "")
	}

	return nil
}

// 修改排序权重
func (m *Book) SetSortWeight(ctx context.Context, weight int, newUSN int) error {
	mut := m

	mut.SetField("Seq", weight)
	mut.SetField("Usn", newUSN)

	return nil
}

func (m *Book) SetParent(ctx context.Context, parent *Book, newUSN int) error {
	mut := m

	mut.SetField("Usn", newUSN)
	if parent == nil {
		mut.SetField("ParentNotebookId", "")
	} else {
		data, err := parent.Data(ctx)
		if err != nil {
			return err
		}

		mut.SetField("ParentNotebookId", data.NotebookId)
	}

	// TODO: 修改关联关系对象
	// m.Edges.Parent = parent

	return nil
}

func (m *Book) RefreshNumberNotes(ctx context.Context, count int) error {
	mut := m

	mut.SetField("NumberNotes", count)

	return nil
}

func (m *Book) SetBlogStatus(ctx context.Context, blog bool, newUSN int) error {
	mut := m

	mut.SetField("Usn", newUSN)
	mut.SetField("IsBlog", blog)

	return nil
}
