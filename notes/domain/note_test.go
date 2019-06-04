package domain

import (
	"testing"

	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/events"
	"github.com/stretchr/testify/assert"
)

func TestNote(t *testing.T) {
	t.Run("CreateNote", func(t *testing.T) {
		id := "ID"
		title := "Title"
		content := "Content"
		note, evts := CreateNote(id, title, content)

		assert.Equal(t,
			Note{ID: id, Title: title, Content: content},
			note,
		)
		assert.Equal(t, []cqrs.Event{
			events.NoteCreatedEvent{
				ID: id, Title: title, Content: content,
			},
		}, evts)
	})
}
