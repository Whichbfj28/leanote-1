// Code generated by ImplGen.
// Source: history_repository.go

package mongo

import (
	"github.com/coocn-cn/leanote/app/info"
	model "github.com/coocn-cn/leanote/app/note/model"
)

type historyMutation struct {
	model.HistoryData

	db  *info.NoteContentHistory
	old *model.HistoryData
}

func newHistoryMutation(data model.HistoryData, db *info.NoteContentHistory) *historyMutation {
	m := &historyMutation{
		db:          db,
		old:         nil,
		HistoryData: data,
	}

	if db != nil {
		old := data
		m.old = &old
	}

	return m
}

func (m *historyMutation) Data() model.HistoryData {
	return m.HistoryData
}

func (m *historyMutation) SetHistories(histories []info.EachHistory) {
	m.Histories = histories
}
