package events

type NoteCreatedEvent struct {
	ID      string
	Title   string
	Content string
}

func NewNoteCreatedEvent(id, title, content string) NoteCreatedEvent {
	return NoteCreatedEvent{
		ID:      id,
		Title:   title,
		Content: content,
	}
}
