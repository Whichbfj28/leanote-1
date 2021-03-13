package controllers

import (
	"context"

	"github.com/coocn-cn/leanote/app/note/repository"
	"github.com/coocn-cn/leanote/app/note/repository/mongo"
	"github.com/revel/revel"
)

type NoteContentHistory struct {
	BaseController
}

// 得到list
func (c NoteContentHistory) ListHistories(noteId string) revel.Result {
	ctx := context.Background()

	noteR := mongo.NewNote(nil)
	historyR := mongo.NewHistory(nil)

	note, err := noteR.Find(ctx, repository.NoteID(noteId))
	if err != nil {
		c.RenderError(err)
	}

	history, err := historyR.Find(ctx, repository.HistoryNoteID(note.MustData(ctx).NoteId.Hex()))
	if err != nil {
		c.RenderError(err)
	}

	if history == nil {
		return c.RenderJSON([]string{})
	}

	data, err := history.Data(ctx)
	if err != nil {
		c.RenderError(err)
	}

	return c.RenderJSON(data.Histories)
}
