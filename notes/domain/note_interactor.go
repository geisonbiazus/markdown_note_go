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
	builder.Add(events.NoteCreatedEvent{ID: i.idGen.Generate(), Title: title, Content: content})
	return builder.Events
}

func (i *NoteInteractor) UpdateNote(note Note, title, content string) []cqrs.Event {
	builder := cqrs.NewEventsBuilder()
	builder.Add(events.NoteUpdatedEvent{ID: note.ID, Title: title, Content: content})
	return builder.Events
}
