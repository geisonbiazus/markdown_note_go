package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/events"
)

func TestNote(t *testing.T) {
	t.Run("ApplyEvents", func(t *testing.T) {
		t.Run("NonRecognizedEvent", func(t *testing.T) {
			note := EmptyNote

			evts := []cqrs.Event{FakeEvent{"id"}}

			(&note).ApplyEvents(evts)

			assert.Equal(t, EmptyNote, note)
		})

		t.Run("NoteCreatedEvent", func(t *testing.T) {
			note := EmptyNote

			evts := []cqrs.Event{
				events.NoteCreatedEvent{ID: "id", Title: "title", Content: "content"},
			}

			(&note).ApplyEvents(evts)

			assert.Equal(t,
				Note{ID: "id", Title: "title", Content: "content"},
				note,
			)
		})

		t.Run("NoteUpdatedEvent", func(t *testing.T) {
			note := Note{ID: "id", Title: "title", Content: "content"}

			evts := []cqrs.Event{
				events.NoteUpdatedEvent{ID: "id", Title: "newTitle", Content: "newContent"},
			}

			(&note).ApplyEvents(evts)

			assert.Equal(t,
				Note{ID: "id", Title: "newTitle", Content: "newContent"},
				note,
			)
		})

		t.Run("Multipl events", func(t *testing.T) {
			note := EmptyNote

			evts := []cqrs.Event{
				events.NoteCreatedEvent{ID: "id", Title: "title", Content: "content"},
				events.NoteUpdatedEvent{ID: "id", Title: "newTitle", Content: "content"},
				events.NoteUpdatedEvent{ID: "id", Title: "newTitle", Content: "newContent"},
			}

			(&note).ApplyEvents(evts)

			assert.Equal(t,
				Note{ID: "id", Title: "newTitle", Content: "newContent"},
				note,
			)
		})
	})
}

type FakeEvent struct {
	ID string
}
