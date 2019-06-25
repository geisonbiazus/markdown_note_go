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
	builder := cqrs.NewEventsBuilder()
	builder.Add(events.NoteCreatedEvent{i.idGen.Generate(), title, content})
	return builder.Events
}

func (i *NoteInteractor) UpdateNote(note Note, title, content string) []cqrs.Event {
	builder := cqrs.NewEventsBuilder()
	builder.Add(events.NoteUpdatedEvent{note.ID, title, content})
	return builder.Events
}
