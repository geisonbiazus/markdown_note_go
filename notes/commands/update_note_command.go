package commands

import "github.com/geisonbiazus/markdown_notes/validations"

type UpdateNoteCommand struct {
	ID      string
	Title   string
	Content string
}

func (c UpdateNoteCommand) Validate() validations.Result {
	v := validations.NewValidator()
	v.ValidateRequired("ID", c.ID)
	v.ValidateRequired("Title", c.Title)

	return v.Result()
}
