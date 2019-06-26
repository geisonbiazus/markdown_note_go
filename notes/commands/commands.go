package commands

type CreateNoteCommand struct {
	Title   string
	Content string
}

type UpdateNoteCommand struct {
	ID      string
	Title   string
	Content string
}
