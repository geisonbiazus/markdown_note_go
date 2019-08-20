package notes

import (
	"testing"

	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/commands"
	"github.com/geisonbiazus/markdown_notes/notes/domain"
	"github.com/geisonbiazus/markdown_notes/notes/events"
	"github.com/geisonbiazus/markdown_notes/validations"
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
		t.Run("With valid arguments", func(t *testing.T) {
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

			t.Run("Returns the output containing the ID o the created node", func(t *testing.T) {
				noteID := "ID"
				f := setup()

				f.idGenerator.NextID = noteID

				output := f.handler.CreateNote(
					commands.CreateNoteCommand{Title: "Title", Content: "Content"},
				)

				assert.Equal(t, CreateNoteOutput{Valid: true, ID: noteID}, output)
			})
		})
		t.Run("With invalid arguments", func(t *testing.T) {
			t.Run("Returns the errors", func(t *testing.T) {
				f := setup()

				output := f.handler.CreateNote(
					commands.CreateNoteCommand{Title: "", Content: "Content"},
				)

				assert.Equal(t, CreateNoteOutput{
					Valid: false,
					Errors: []validations.Error{
						validations.Error{Field: "Title", Type: "REQUIRED"},
					},
				}, output)
			})
		})
	})

	t.Run("UpdateNote", func(t *testing.T) {
		t.Run("With valid arguments", func(t *testing.T) {
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

		t.Run("With invalid arguments", func(t *testing.T) {
			t.Run("Returns the errors", func(t *testing.T) {
				f := setup()

				createNoteOutput := f.handler.CreateNote(
					commands.CreateNoteCommand{Title: "Title", Content: "Content"},
				)

				output := f.handler.UpdateNote(
					commands.UpdateNoteCommand{
						ID:      createNoteOutput.ID,
						Title:   "",
						Content: "NewContent",
					},
				)

				assert.Equal(t, UpdateNoteOutput{
					Valid: false,
					Errors: []validations.Error{
						validations.Error{Field: "Title", Type: "REQUIRED"},
					},
				}, output)
			})
		})
	})
}
