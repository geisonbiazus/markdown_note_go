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
		t.Run("Generates a NoteCreatedEvent", func(t *testing.T) {
			id := "ID"
			title := "Title"
			content := "Content"

			f := setup()
			f.idGenerator.NextID = id

			_, evts := f.interactor.CreateNote(title, content)

			assert.Equal(t,
				[]cqrs.Event{
					events.NoteCreatedEvent{ID: id, Title: title, Content: content},
				},
				evts,
			)
		})

		t.Run("Returns the created Note", func(t *testing.T) {
			id := "ID"
			title := "Title"
			content := "Content"

			f := setup()
			f.idGenerator.NextID = id

			note, _ := f.interactor.CreateNote(title, content)

			assert.Equal(t, Note{ID: id, Title: title, Content: content}, note)
		})
	})

	t.Run("UpdateNote", func(t *testing.T) {
		t.Run("Generates a NoteUpdatedEvent", func(t *testing.T) {
			f := setup()

			note := Note{ID: "id", Title: "Title", Content: "Content"}

			title := "New Title"
			content := "New Content"

			_, evts := f.interactor.UpdateNote(note, title, content)

			assert.Equal(t,
				[]cqrs.Event{
					events.NoteUpdatedEvent{ID: note.ID, Title: title, Content: content},
				},
				evts,
			)
		})

		t.Run("Returns the updated note", func(t *testing.T) {
			f := setup()

			note := Note{ID: "id", Title: "Title", Content: "Content"}

			title := "New Title"
			content := "New Content"

			updatedNote, _ := f.interactor.UpdateNote(note, title, content)

			assert.Equal(t, Note{ID: note.ID, Title: title, Content: content}, updatedNote)
		})
	})
}
