package notes

import (
	"github.com/geisonbiazus/markdown_notes/notes/commands"
	"github.com/geisonbiazus/markdown_notes/notes/domain"
	"github.com/geisonbiazus/markdown_notes/notes/events"
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
	if result := cmd.Validate(); result.Valid {
		evts := h.noteInteractor.CreateNote(cmd.Title, cmd.Content)
		h.repo.PublishEvents(evts)
		return CreateNoteOutput{
			ID:    evts[0].(events.NoteCreatedEvent).ID,
			Valid: true,
		}
	} else {
		return CreateNoteOutput{
			Valid:  result.Valid,
			Errors: result.Errors,
		}
	}
}

func (h *CommandHandler) UpdateNote(cmd commands.UpdateNoteCommand) {
	note := h.repo.GetNoteByID(cmd.ID)
	evts := h.noteInteractor.UpdateNote(note, cmd.Title, cmd.Content)
	h.repo.PublishEvents(evts)
}
