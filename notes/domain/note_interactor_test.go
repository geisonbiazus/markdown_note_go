package domain

import (
	"testing"

	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/events"
	"github.com/stretchr/testify/assert"
)

func TestNoteInteractor(t *testing.T) {
	type fixture struct {
		idGenerator *FakeIDGenerator
		interactor  *NoteInteractor
	}

	setup := func() *fixture {
		idGenerator := NewFakeIdGenerator("ID")
		interactor := NewNoteInteractor(idGenerator)
		return &fixture{
			idGenerator,
			interactor,
		}
	}

	t.Run("CreateNote", func(t *testing.T) {
		t.Run("generates a NoteCreatedEvent", func(t *testing.T) {
			id := "ID"
			title := "Title"
			content := "Content"

			f := setup()
			f.idGenerator.NextID = id

			assert.Equal(t,
				[]cqrs.Event{
					events.NoteCreatedEvent{ID: id, Title: title, Content: content},
				},
				f.interactor.CreateNote(title, content),
			)
		})
	})

	t.Run("UpdateNote", func(t *testing.T) {
		t.Run("generates a NoteUpdatedEvent", func(t *testing.T) {
			f := setup()

			note := Note{ID: "id", Title: "Title", Content: "Content"}

			title := "New Title"
			content := "New Content"

			assert.Equal(t,
				[]cqrs.Event{
					events.NoteUpdatedEvent{ID: note.ID, Title: title, Content: content},
				},
				f.interactor.UpdateNote(note, title, content),
			)
		})
	})
}
