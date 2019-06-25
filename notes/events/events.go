package events

type NoteCreatedEvent struct {
	ID      string
	Title   string
	Content string
}

type NoteUpdatedEvent struct {
	ID      string
	Title   string
	Content string
}
