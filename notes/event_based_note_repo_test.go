package notes

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/domain"
	"github.com/geisonbiazus/markdown_notes/notes/events"
)

func TestEventBasedNoteRepo(t *testing.T) {
	t.Run("GetNoteByID", func(t *testing.T) {
		t.Run("Returns an empty note when the note doesn't exist", func(t *testing.T) {
			store := cqrs.NewInMemoryEventStore()
			repo := NewEventBasedNoteRepo(store)

			assert.Equal(t, domain.EmptyNote, repo.GetNoteByID("id"))
		})

		t.Run("Returns a note when there is a note created event", func(t *testing.T) {
			store := cqrs.NewInMemoryEventStore()
			repo := NewEventBasedNoteRepo(store)

			store.AddEvent(events.NoteCreatedEvent{
				ID: "id", Title: "title", Content: "content",
			})

			assert.Equal(t, domain.Note{
				ID: "id", Title: "title", Content: "content",
			}, repo.GetNoteByID("id"))
		})

		t.Run("Returns the last version of the note when there are more than one event", func(t *testing.T) {
			store := cqrs.NewInMemoryEventStore()
			repo := NewEventBasedNoteRepo(store)

			store.AddEvent(
				events.NoteCreatedEvent{ID: "id", Title: "title", Content: "content"},
			)
			store.AddEvent(
				events.NoteUpdatedEvent{ID: "id", Title: "new title", Content: "new content"},
			)

			assert.Equal(t, domain.Note{
				ID: "id", Title: "new title", Content: "new content",
			}, repo.GetNoteByID("id"))
		})
	})
}
