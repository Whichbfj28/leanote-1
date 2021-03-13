package model

import (
	"context"

	"github.com/coocn-cn/leanote/app/info"
)

const maxSize = 10

type HistoryData info.NoteContentHistory

type HistoryMutation interface {
	Data(context.Context) (HistoryData, error)
	SetField(name string, value interface{})
	SetFields(fields map[string]interface{})
}

type History struct {
	HistoryMutation
}

func NewHistory(mut HistoryMutation) *History {
	return &History{HistoryMutation: mut}
}

func (m *History) MustData(ctx context.Context) HistoryData {
	data, err := m.Data(ctx)
	if err != nil {
		panic(err)
	}

	return data
}

// 判断是否超出 maxSize, 如果超出则pop最后一个, 再push之, 不用那么麻烦, 直接update吧, 虽然影响性能
func (m *History) AddHistory(ctx context.Context, history info.EachHistory) error {
	mut := m
	data, err := m.Data(ctx)
	if err != nil {
		return err
	}
	old := data.Histories

	histories := make([]info.EachHistory, 0, len(old)+1)
	histories = append(histories, history) // 在开头加了, 最近的在最前
	histories = append(histories, old...)

	if len(histories) >= maxSize {
		histories = histories[:maxSize]
	}

	mut.SetField("Histories", histories)

	return nil
}
