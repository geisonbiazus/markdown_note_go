package notes

import (
	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/domain"
)

type EventBasedNoteRepo struct {
	eventStore cqrs.EventStore
}

func NewEventBasedNoteRepo(eventStore cqrs.EventStore) *EventBasedNoteRepo {
	return &EventBasedNoteRepo{eventStore}
}

func (r *EventBasedNoteRepo) PublishEvents(evts []cqrs.Event) error {
	for _, evt := range evts {
		err := r.eventStore.AddEvent(evt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *EventBasedNoteRepo) GetNoteByID(id string) domain.Note {
	evts, _ := r.eventStore.ReadEvents()

	note := domain.NewNote()
	(&note).ApplyEvents(evts)

	return note
}
