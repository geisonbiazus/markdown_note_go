package commands

import (
	"testing"

	"github.com/geisonbiazus/markdown_notes/validations/validationstest"
)

func TestCreateNoteCommand(t *testing.T) {
	t.Run("Validations", func(t *testing.T) {
		t.Run("Valid Command", func(t *testing.T) {
			cmd := validCreateNoteCommand()
			validationstest.AssertValid(t, cmd.Validate())
		})

		t.Run("Title is required", func(t *testing.T) {
			cmd := validCreateNoteCommand()
			cmd.Title = ""

			validationstest.AssertValidationError(t, cmd.Validate(), "Title", "REQUIRED")
		})
	})
}

func validCreateNoteCommand() CreateNoteCommand {
	return CreateNoteCommand{Title: "Title", Content: "Content"}
}
