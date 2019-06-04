package commands

type CreateNoteCommand struct {
	Title   string
	Content string
}

func NewCreateNoteCommand(title, content string) CreateNoteCommand {
	return CreateNoteCommand{Title: title, Content: content}
}
