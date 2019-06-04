package domain

import (
	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/events"
)

type Note struct {
	ID      string
	Title   string
	Content string
}

func CreateNote(id, title, content string) (Note, []cqrs.Event) {
	note := Note{
		ID:      id,
		Title:   title,
		Content: content,
	}

	evts := []cqrs.Event{}
	evts = append(evts, events.NewNoteCreatedEvent(id, title, content))

	return note, evts
}
