package events

type NoteCreatedEvent struct {
	ID      string
	Title   string
	Content string
}

func (e NoteCreatedEvent) GetID() string {
	return e.ID
}

type NoteUpdatedEvent struct {
	ID      string
	Title   string
	Content string
}

func (e NoteUpdatedEvent) GetID() string {
	return e.ID
}
