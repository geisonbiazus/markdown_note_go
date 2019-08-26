package notes

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/domain"
	"github.com/geisonbiazus/markdown_notes/notes/events"
)

func TestEventBasedNoteRepo(t *testing.T) {
	t.Run("PublishEvents", func(t *testing.T) {
		t.Run("publishes the events to the store", func(t *testing.T) {
			store := cqrs.NewInMemoryEventStore()
			repo := NewEventBasedNoteRepo(store)

			evts := []cqrs.Event{
				events.NoteCreatedEvent{ID: "id1", Title: "title1", Content: "content1"},
				events.NoteCreatedEvent{ID: "id2", Title: "title2", Content: "content2"},
			}

			repo.PublishEvents(evts)

			assert.Equal(t, evts, store.Events)
		})

		t.Run("Returns the error when some error happen", func(t *testing.T) {
			err := errors.New("Error")
			store := NewErrorReturningEventStore(err)
			repo := NewEventBasedNoteRepo(store)

			evts := []cqrs.Event{
				events.NoteCreatedEvent{ID: "id1", Title: "title1", Content: "content1"},
				events.NoteCreatedEvent{ID: "id2", Title: "title2", Content: "content2"},
			}

			assert.Equal(t, err, repo.PublishEvents(evts))
		})
	})

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

		t.Run("Return the correct note when events from more than are note is stored", func(t *testing.T) {
			store := cqrs.NewInMemoryEventStore()
			repo := NewEventBasedNoteRepo(store)

			store.AddEvent(
				events.NoteCreatedEvent{ID: "1", Title: "title", Content: "content"},
			)

			store.AddEvent(
				events.NoteCreatedEvent{ID: "2", Title: "other title", Content: "other content"},
			)

			store.AddEvent(
				events.NoteUpdatedEvent{ID: "1", Title: "new title", Content: "new content"},
			)

			assert.Equal(t, domain.Note{ID: "1", Title: "new title", Content: "new content"}, repo.GetNoteByID("1"))
			assert.Equal(t, domain.Note{ID: "2", Title: "other title", Content: "other content"}, repo.GetNoteByID("2"))
		})
	})
}

type ErrorReturningEventStore struct {
	Error error
}

func NewErrorReturningEventStore(err error) *ErrorReturningEventStore {
	return &ErrorReturningEventStore{err}
}

func (s *ErrorReturningEventStore) AddEvent(event cqrs.Event) error {
	return s.Error
}

func (s *ErrorReturningEventStore) ReadAllEvents() ([]cqrs.Event, error) {
	return []cqrs.Event{}, s.Error
}

func (s *ErrorReturningEventStore) ReadEventsByID(id string) ([]cqrs.Event, error) {
	return []cqrs.Event{}, s.Error
}
