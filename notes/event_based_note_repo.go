package notes

import (
	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/domain"
	"github.com/geisonbiazus/markdown_notes/notes/events"
)

type EventBasedNoteRepo struct {
	eventStore cqrs.EventStore
}

func NewEventBasedNoteRepo(eventStore cqrs.EventStore) *EventBasedNoteRepo {
	return &EventBasedNoteRepo{eventStore}
}

func (r *EventBasedNoteRepo) GetByID(id string) domain.Note {
	store := r.eventStore.(*cqrs.InMemoryEventStore)
	evt := store.Events[0].(events.NoteCreatedEvent)

	return domain.Note{ID: evt.ID, Title: evt.Title, Content: evt.Content}
}
