package model

import (
	"context"
	"time"

	"github.com/coocn-cn/leanote/app/info"
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
