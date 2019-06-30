package domain

import (
	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/events"
)

var EmptyNote = Note{}

type Note struct {
	ID      string
	Title   string
	Content string
}

func NewNote() Note {
	return Note{}
}

func (n *Note) ApplyEvents(evts []cqrs.Event) {
	for _, event := range evts {
		switch evt := event.(type) {
		case events.NoteCreatedEvent:
			n.applyNoteCreatedEvent(evt)
		case events.NoteUpdatedEvent:
			n.applyNoteUpdatedEvent(evt)
		}
	}
}

func (n *Note) applyNoteCreatedEvent(evt events.NoteCreatedEvent) {
	n.ID = evt.ID
	n.Title = evt.Title
	n.Content = evt.Content
}

func (n *Note) applyNoteUpdatedEvent(evt events.NoteUpdatedEvent) {
	n.Title = evt.Title
	n.Content = evt.Content
}
