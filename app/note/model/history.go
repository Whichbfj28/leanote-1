package model

import (
	"github.com/coocn-cn/leanote/app/info"
)

const maxSize = 10

type HistoryData info.NoteContentHistory

type HistoryMutation interface {
	Data() HistoryData
	SetHistories([]info.EachHistory)
}

type History struct {
	old      HistoryData
	mutation HistoryMutation
}

func NewHistory(mut HistoryMutation) *History {
	return &History{
		old:      mut.Data(),
		mutation: mut,
	}
}

func (m *History) M() HistoryMutation {
	return m.mutation
}

func (m *History) Data() HistoryData {
	return m.M().Data()
}

func (m *History) AddHistory(history info.EachHistory) {
	// 判断是否超出 maxSize, 如果超出则pop最后一个, 再push之, 不用那么麻烦, 直接update吧, 虽然影响性能

	mut := m.M()
	old := m.Data().Histories

	histories := make([]info.EachHistory, 0, len(old)+1)
	histories = append(histories, history) // 在开头加了, 最近的在最前
	histories = append(histories, old...)

	if len(histories) >= maxSize {
		histories = histories[:maxSize]
	}

	mut.SetHistories(histories)
}
