package commands

import (
	"testing"
)

func TestCreateNoteCommand(t *testing.T) {
	t.Run("Validations", func(t *testing.T) {
		t.Run("Valid Command", func(t *testing.T) {
			cmd := validCreateNoteCommand()
			assertValid(t, cmd.Validate())
		})

		t.Run("Title is required", func(t *testing.T) {
			cmd := validCreateNoteCommand()
			cmd.Title = ""

			assertValidationError(t, cmd.Validate(), "Title", "REQUIRED")
		})
	})
}

func validCreateNoteCommand() CreateNoteCommand {
	return CreateNoteCommand{Title: "Title", Content: "Content"}
}
