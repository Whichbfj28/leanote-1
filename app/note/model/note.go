package model

import (
	"context"
	"time"

	"github.com/coocn-cn/leanote/app/info"
	"github.com/coocn-cn/leanote/app/lea"
	"gopkg.in/mgo.v2/bson"
)

type NoteData info.Note

type NoteMutation interface {
	Data(context.Context) (NoteData, error)
	SetField(name string, value interface{})
	SetFields(fields map[string]interface{})
}

type Note struct {
	NoteMutation
}

func NewNote(mut NoteMutation) *Note {
	return &Note{
		NoteMutation: mut,
	}
}

func (m *Note) MustData(ctx context.Context) NoteData {
	data, err := m.Data(ctx)
	if err != nil {
		panic(err)
	}

	return data
}

func (m *Note) SetUSN(ctx context.Context, usn int) error {
	noteMut := m

	noteMut.SetField("Usn", usn)
	noteMut.SetField("UpdatedTime", time.Now())

	return nil
}

func (m *Note) SetTitle(ctx context.Context, operateUserID, title string, newUSN int) error {
	noteMut := m

	noteMut.SetField("Usn", newUSN)
	noteMut.SetField("Title", title)
	noteMut.SetField("UpdatedTime", time.Now())
	noteMut.SetField("UpdatedUserId", bson.ObjectIdHex(operateUserID))

	return nil
}

func (m *Note) SetTags(ctx context.Context, tags []string, newUSN int) error {
	noteMut := m

	noteMut.SetField("Usn", newUSN)
	noteMut.SetField("Tags", tags)

	return nil
}

func (m *Note) SetBookID(ctx context.Context, bookID string, newUSN int) error {
	noteMut := m

	noteMut.SetField("Usn", newUSN)
	noteMut.SetField("NotebookId", bookID)

	return nil
}

func (m *Note) SetTrash(ctx context.Context, trash bool, newUSN int) error {
	noteMut := m

	noteMut.SetField("Usn", newUSN)
	noteMut.SetField("IsTrash", trash)

	return nil
}

func (m *Note) SetBlogStatus(ctx context.Context, blog bool, newUSN int) error {
	noteMut := m

	noteMut.SetField("Usn", newUSN)
	noteMut.SetField("IsBlog", blog)
	if blog {
		noteMut.SetField("PublicTime", time.Now())
	} else {
		noteMut.SetField("HasSelfDefined", false)
	}

	return nil
}

func (m *Note) SetTopStatus(ctx context.Context, top bool, newUSN int) error {
	noteMut := m

	noteMut.SetField("Usn", newUSN)
	noteMut.SetField("IsTop", top)

	return nil
}

func (m *Note) Updete_needdelete_(ctx context.Context, needUpdate bson.M, operateUserID string, newUSN int) error {
	noteMut := m

	// 可以将时间传过来
	updatedTime, ok := needUpdate["UpdatedTime"].(time.Time)
	if ok {
		updatedTime = lea.FixUrlTime(updatedTime)
	} else {
		updatedTime = time.Now()
	}
	delete(needUpdate, "UpdatedTime") // 使用 SetField 更新

	noteMut.SetFields(needUpdate)

	noteMut.SetField("Usn", newUSN)
	noteMut.SetField("UpdatedTime", updatedTime)
	noteMut.SetField("UpdatedUserId", bson.ObjectIdHex(operateUserID))

	return nil
}
