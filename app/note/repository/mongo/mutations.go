// Code generated by ImplGen.
// Source: book_repository.go

package mongo

import (
	context "context"

	"github.com/coocn-cn/leanote/app/info"
	model "github.com/coocn-cn/leanote/app/note/model"
)

type bookData struct {
	*baseMutation
}

func newbookData(data model.BookData, db *info.Notebook, repo *book) *bookData {
	return &bookData{baseMutation: newbaseMutation(data, db)}
}

func (m *bookData) db() *info.Notebook {
	return m.baseMutation.db.(*info.Notebook)
}

func (m *bookData) Data(ctx context.Context) (data model.BookData, err error) {
	return data, m.baseMutation.Data(ctx, &data)
}

type historyMutation struct {
	*baseMutation
}

func newHistoryMutation(data model.HistoryData, db *info.NoteContentHistory) *historyMutation {
	return &historyMutation{baseMutation: newbaseMutation(data, db)}
}

func (m *historyMutation) db() *info.NoteContentHistory {
	return m.baseMutation.db.(*info.NoteContentHistory)
}

func (m *historyMutation) Data(ctx context.Context) (data model.HistoryData, err error) {
	return data, m.baseMutation.Data(ctx, &data)
}

type contentMutation struct {
	*baseMutation
}

func newContentMutation(data model.ContentData, db *info.NoteContent) *contentMutation {
	return &contentMutation{baseMutation: newbaseMutation(data, db)}
}

func (m *contentMutation) db() *info.NoteContent {
	return m.baseMutation.db.(*info.NoteContent)
}

func (m *contentMutation) Data(ctx context.Context) (data model.ContentData, err error) {
	return data, m.baseMutation.Data(ctx, &data)
}

type noteData struct {
	*baseMutation
}

func newnoteData(data model.NoteData, db *info.Note, repo *note) *noteData {
	return &noteData{baseMutation: newbaseMutation(data, db)}
}

func (m *noteData) db() *info.Note {
	return m.baseMutation.db.(*info.Note)
}

func (m *noteData) Data(ctx context.Context) (data model.NoteData, err error) {
	return data, m.baseMutation.Data(ctx, &data)
}

type tagMutation struct {
	*baseMutation
}

func newtagMutation(data model.TagData, db *info.NoteTag) *tagMutation {
	return &tagMutation{baseMutation: newbaseMutation(data, db)}
}

func (m *tagMutation) db() *info.NoteTag {
	return m.baseMutation.db.(*info.NoteTag)
}

func (m *tagMutation) Data(ctx context.Context) (data model.TagData, err error) {
	return data, m.baseMutation.Data(ctx, &data)
}