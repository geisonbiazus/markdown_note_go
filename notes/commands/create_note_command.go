package commands

type CreateNoteCommand struct {
	Title   string
	Content string
}

func (c CreateNoteCommand) Validate() ValidationResult {
	v := NewValidator()
	v.ValidateRequired("Title", c.Title)

	return v.Result()
}
