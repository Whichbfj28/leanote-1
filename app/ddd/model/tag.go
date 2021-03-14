package model

import (
	"context"

	"github.com/coocn-cn/leanote/app/info"
)

type TagData info.Tag

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
