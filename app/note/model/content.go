package model

import (
	"context"
	"time"

	"github.com/coocn-cn/leanote/app/info"
	"github.com/coocn-cn/leanote/app/lea"
	"gopkg.in/mgo.v2/bson"
)

type ContentData info.NoteContent

type ContentMutation interface {
	Data(context.Context) (ContentData, error)
	SetField(name string, value interface{})
	SetFields(fields map[string]interface{})
}

type Content struct {
	ContentMutation
}

func NewContent(mut ContentMutation) *Content {
	return &Content{ContentMutation: mut}
}

func (m *Content) MustData(ctx context.Context) ContentData {
	data, err := m.Data(ctx)
	if err != nil {
		panic(err)
	}

	return data
}

func (m *Content) SetBlogStatus(ctx context.Context, blog bool) error {
	m.SetField("IsBlog", blog)

	return nil
}

func (m *Content) SetAbstract(ctx context.Context, abstract string) error {
	mut := m

	mut.SetFields(map[string]interface{}{
		"Abstract": abstract,
	})

	return nil
}

func (m *Content) SaveContent(ctx context.Context, operateUserID string, content string, updatedTime time.Time) error {
	mut := m

	updatedTime = lea.FixUrlTime(updatedTime)

	mut.SetFields(map[string]interface{}{
		"Content":       content,
		"UpdatedTime":   updatedTime,
		"UpdatedUserId": bson.ObjectIdHex(operateUserID),
	})

	return nil
}
