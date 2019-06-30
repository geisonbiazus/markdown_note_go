package notes

import (
	"testing"

	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/commands"
	"github.com/geisonbiazus/markdown_notes/notes/domain"
	"github.com/geisonbiazus/markdown_notes/notes/events"
	"github.com/stretchr/testify/assert"
)

func TestCommandHandler(t *testing.T) {
	type fixture struct {
		store       *cqrs.InMemoryEventStore
		idGenerator *domain.FakeIDGenerator
		handler     *CommandHandler
	}

	setup := func() *fixture {
		store := cqrs.NewInMemoryEventStore()
		idGenerator := domain.NewFakeIdGenerator("ID")

		handler := NewCommandHandler(
			NewEventBasedNoteRepo(store),
			domain.NewNoteInteractor(idGenerator),
		)
		return &fixture{
			store,
			idGenerator,
			handler,
		}
	}

	t.Run("CreateNote", func(t *testing.T) {
		t.Run("Creates a new note", func(t *testing.T) {
			noteID := "ID"
			f := setup()

			f.idGenerator.NextID = noteID

			f.handler.CreateNote(
				commands.CreateNoteCommand{Title: "Title", Content: "Content"},
			)

			assert.Equal(t, []cqrs.Event{
				events.NoteCreatedEvent{ID: noteID, Title: "Title", Content: "Content"},
			}, f.store.Events)
		})
	})

	t.Run("UpdateNote", func(t *testing.T) {
		t.Run("Updates a note", func(t *testing.T) {
			noteID := "NoteID"

			f := setup()
			f.idGenerator.NextID = noteID

			f.handler.CreateNote(
				commands.CreateNoteCommand{Title: "Title", Content: "Content"},
			)

			f.handler.UpdateNote(
				commands.UpdateNoteCommand{
					ID:      noteID,
					Title:   "NewTitle",
					Content: "NewContent",
				},
			)

			assert.Equal(t, []cqrs.Event{
				events.NoteCreatedEvent{ID: noteID, Title: "Title", Content: "Content"},
				events.NoteUpdatedEvent{ID: noteID, Title: "NewTitle", Content: "NewContent"},
			}, f.store.Events)
		})
	})
}
