package notes

import (
	"github.com/geisonbiazus/markdown_notes/notes/commands"
	"github.com/geisonbiazus/markdown_notes/notes/domain"
	"github.com/geisonbiazus/markdown_notes/validations"
)

type CommandHandler struct {
	repo           *EventBasedNoteRepo
	noteInteractor *domain.NoteInteractor
}

func NewCommandHandler(
	repo *EventBasedNoteRepo, noteInteractor *domain.NoteInteractor,
) *CommandHandler {
	return &CommandHandler{repo, noteInteractor}
}

type CreateNoteOutput struct {
	ID     string
	Valid  bool
	Errors []validations.Error
}

func (h *CommandHandler) CreateNote(cmd commands.CreateNoteCommand) CreateNoteOutput {
	result := cmd.Validate()

	if result.Valid {
		return h.createNote(cmd)
	}

	return h.invalidCreateNoteOutput(result)
}

func (h *CommandHandler) createNote(cmd commands.CreateNoteCommand) CreateNoteOutput {
	note, evts := h.noteInteractor.CreateNote(cmd.Title, cmd.Content)
	h.repo.PublishEvents(evts)

	return CreateNoteOutput{ID: note.ID, Valid: true}
}

func (h *CommandHandler) invalidCreateNoteOutput(result validations.Result) CreateNoteOutput {
	return CreateNoteOutput{Valid: result.Valid, Errors: result.Errors}
}

type UpdateNoteOutput struct {
	ID     string
	Valid  bool
	Errors []validations.Error
}

func (h *CommandHandler) UpdateNote(cmd commands.UpdateNoteCommand) UpdateNoteOutput {
	result := cmd.Validate()

	note := h.repo.GetNoteByID(cmd.ID)
	_, evts := h.noteInteractor.UpdateNote(note, cmd.Title, cmd.Content)
	h.repo.PublishEvents(evts)

	return UpdateNoteOutput{Valid: result.Valid, Errors: result.Errors}
}
