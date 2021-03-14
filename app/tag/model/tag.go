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

// 判断是否超出 maxSize, 如果超出则pop最后一个, 再push之, 不用那么麻烦, 直接update吧, 虽然影响性能
func (m *Tag) AddTags(ctx context.Context, tags []string) error {
	mut := m
	data, err := m.Data(ctx)
	if err != nil {
		return err
	}

	mut.SetField("Tags", append(data.Tags, tags...))

	return nil
}

func (m *Tag) SoftDelete(ctx context.Context, deleted bool, newUSN int) error {
	mut := m

	mut.SetField("Usn", newUSN)
	mut.SetField("IsDeleted", deleted)

	return nil
}
