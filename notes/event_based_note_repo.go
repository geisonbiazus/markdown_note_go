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

func (r *EventBasedNoteRepo) GetNoteByID(id string) domain.Note {
	evts, _ := r.eventStore.ReadEvents()

	note := domain.EmptyNote

	for _, evt := range evts {
		switch event := evt.(type) {
		case events.NoteCreatedEvent:
			note.ID = event.ID
			note.Title = event.Title
			note.Content = event.Content
		case events.NoteUpdatedEvent:
			note.Title = event.Title
			note.Content = event.Content
		}
	}

	return note
}
