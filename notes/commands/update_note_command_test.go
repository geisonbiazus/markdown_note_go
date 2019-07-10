package commands

import (
	"testing"

	"github.com/geisonbiazus/markdown_notes/validations/validationstest"
)

func TestUpdateNoteCommand(t *testing.T) {
	t.Run("Validations", func(t *testing.T) {
		t.Run("Valid Command", func(t *testing.T) {
			cmd := validUpdateNoteCommand()
			validationstest.AssertValid(t, cmd.Validate())
		})

		t.Run("required fields", func(t *testing.T) {
			cmd := UpdateNoteCommand{}

			validationstest.AssertValidationError(t, cmd.Validate(), "ID", "REQUIRED")
			validationstest.AssertValidationError(t, cmd.Validate(), "Title", "REQUIRED")
		})
	})
}

func validUpdateNoteCommand() UpdateNoteCommand {
	return UpdateNoteCommand{ID: "ID", Title: "Title", Content: "Content"}
}
