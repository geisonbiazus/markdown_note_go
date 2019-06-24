package domain

import (
	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/events"
)

type NoteInteractor struct {
	idGen IDGenerator
}

func NewNoteInteractor(idGen IDGenerator) *NoteInteractor {
	return &NoteInteractor{idGen}
}

func (i *NoteInteractor) CreateNote(title, content string) []cqrs.Event {
	evts := []cqrs.Event{}
	evts = append(evts, events.NewNoteCreatedEvent(i.idGen.Generate(), title, content))
	return evts
}
