package notes

import (
	"testing"

	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/commands"
	"github.com/geisonbiazus/markdown_notes/notes/events"
	"github.com/stretchr/testify/assert"
)

func TestCommandHandler(t *testing.T) {
	t.Run("Creates a new note", func(t *testing.T) {
		noteID := "Note ID"
		idGenerator := NewFakeIdGenerator(noteID)
		store := cqrs.NewInMemoryEventStore()

		handler := NewCommandHandler(store, idGenerator)

		command := commands.NewCreateNoteCommand("Title", "Content")

		handler.CreateNote(command)

		assert.Equal(t, []cqrs.Event{
			events.NoteCreatedEvent{ID: noteID, Title: "Title", Content: "Content"},
		}, store.Events)
	})
}
