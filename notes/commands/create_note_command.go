package commands

import "github.com/geisonbiazus/markdown_notes/validations"

type CreateNoteCommand struct {
	Title   string
	Content string
}

func (c CreateNoteCommand) Validate() validations.Result {
	v := validations.NewValidator()
	v.ValidateRequired("Title", c.Title)

	return v.Result()
}
