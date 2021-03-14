package model

import (
	"context"
	"time"

	"github.com/coocn-cn/leanote/app/info"
)

type TagData info.NoteTag

type TagMutation interface {
	Data(context.Context) (TagData, error)
	SetField(name string, value interface{})
	SetFields(fields map[string]interface{})
}

type Tag struct {
	TagMutation
}

func NewTag(mut TagMutation) *Tag {
	return &Tag{TagMutation: mut}
}

func (m *Tag) MustData(ctx context.Context) TagData {
	data, err := m.Data(ctx)
	if err != nil {
		panic(err)
	}

	return data
}

func (m *Tag) SoftDelete(ctx context.Context, deleted bool, newUSN int) error {
	mut := m

	mut.SetField("Usn", newUSN)
	mut.SetField("IsDeleted", deleted)

	return nil
}

func (m *Tag) SetCount(ctx context.Context, counnt int) error {
	mut := m

	mut.SetField("Count", counnt)
	mut.SetField("UpdatedTime", time.Now())

	return nil
}
