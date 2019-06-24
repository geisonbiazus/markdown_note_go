package domain

import (
	"testing"

	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/events"
	"github.com/stretchr/testify/assert"
)

func TestNoteInteractor(t *testing.T) {
	t.Run("createNote", func(t *testing.T) {
		t.Run("returns a createdNoteEvent", func(t *testing.T) {
			id := "ID"
			title := "Title"
			content := "Content"

			interactor := NewNoteInteractor(NewFakeIdGenerator(id))

			assert.Equal(t,
				[]cqrs.Event{
					events.NoteCreatedEvent{ID: id, Title: title, Content: content},
				},
				interactor.CreateNote(title, content),
			)
		})
	})
}
