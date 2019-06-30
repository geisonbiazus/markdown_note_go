package notes

import (
	"github.com/geisonbiazus/markdown_notes/notes/commands"
	"github.com/geisonbiazus/markdown_notes/notes/domain"
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

func (h *CommandHandler) CreateNote(cmd commands.CreateNoteCommand) {
	evts := h.noteInteractor.CreateNote(cmd.Title, cmd.Content)
	h.repo.PublishEvents(evts)
}

func (h *CommandHandler) UpdateNote(cmd commands.UpdateNoteCommand) {
	note := h.repo.GetNoteByID(cmd.ID)
	evts := h.noteInteractor.UpdateNote(note, cmd.Title, cmd.Content)
	h.repo.PublishEvents(evts)
}
