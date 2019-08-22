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

type testCommandHandlerFixture struct {
	store       *cqrs.InMemoryEventStore
	idGenerator *domain.FakeIDGenerator
	handler     *CommandHandler
}

func TestCommandHandler(t *testing.T) {
	setup := func() *testCommandHandlerFixture {
		store := cqrs.NewInMemoryEventStore()
		idGenerator := domain.NewFakeIdGenerator("ID")

		handler := NewCommandHandler(
			NewEventBasedNoteRepo(store),
			domain.NewNoteInteractor(idGenerator),
		)
		return &testCommandHandlerFixture{
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

				createNote(f, "Title", "Content")

				assert.Equal(t, []cqrs.Event{
					events.NoteCreatedEvent{ID: noteID, Title: "Title", Content: "Content"},
				}, f.store.Events)
			})

			t.Run("Returns the output containing the ID of the created node", func(t *testing.T) {
				noteID := "ID"
				f := setup()

				f.idGenerator.NextID = noteID

				output := createNote(f, "Title", "Content")

				assert.Equal(t, CreateNoteOutput{Valid: true, ID: noteID}, output)
			})
		})
		t.Run("With invalid arguments", func(t *testing.T) {
			t.Run("Returns the errors", func(t *testing.T) {
				f := setup()

				output := createNote(f, "", "Content")

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
				f := setup()

				createNoteOutput := createNote(f, "Title", "Content")
				updateNote(f, createNoteOutput.ID, "NewTitle", "NewContent")

				assert.Equal(t, []cqrs.Event{
					events.NoteCreatedEvent{ID: createNoteOutput.ID, Title: "Title", Content: "Content"},
					events.NoteUpdatedEvent{ID: createNoteOutput.ID, Title: "NewTitle", Content: "NewContent"},
				}, f.store.Events)
			})

			t.Run("Returns the output containing the updated note ID", func(t *testing.T) {
				f := setup()

				createNoteOutput := createNote(f, "Title", "Content")
				updateNoteOutput := updateNote(f, createNoteOutput.ID, "NewTitle", "NewContent")

				assert.Equal(t, UpdateNoteOutput{Valid: true, ID: createNoteOutput.ID}, updateNoteOutput)
			})
		})

		t.Run("With invalid arguments", func(t *testing.T) {
			t.Run("Returns the errors", func(t *testing.T) {
				f := setup()

				createNoteOutput := createNote(f, "Title", "Content")
				updateNoteOutput := updateNote(f, createNoteOutput.ID, "", "NewContent")

				assert.Equal(t, UpdateNoteOutput{
					Valid: false,
					Errors: []validations.Error{
						validations.Error{Field: "Title", Type: "REQUIRED"},
					},
				}, updateNoteOutput)
			})

			t.Run("Does not publish events", func(t *testing.T) {
				f := setup()

				createNoteOutput := createNote(f, "Title", "Content")

				updateNote(f, createNoteOutput.ID, "", "NewContent")

				assert.NotContains(t, f.store.Events, events.NoteUpdatedEvent{ID: createNoteOutput.ID, Title: "", Content: "NewContent"})
			})
		})
	})
}

func createNote(f *testCommandHandlerFixture, title, content string) CreateNoteOutput {
	return f.handler.CreateNote(
		commands.CreateNoteCommand{Title: title, Content: content},
	)
}

func updateNote(f *testCommandHandlerFixture, id, title, content string) UpdateNoteOutput {
	return f.handler.UpdateNote(
		commands.UpdateNoteCommand{
			ID:      id,
			Title:   title,
			Content: content,
		},
	)
}
