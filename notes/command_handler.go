package notes

import (
	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/commands"
	"github.com/geisonbiazus/markdown_notes/notes/domain"
)

type CommandHandler struct {
	repo           *EventBasedNoteRepo
	eventStore     cqrs.EventStore
	noteInteractor *domain.NoteInteractor
}

func NewCommandHandler(
	eventStore cqrs.EventStore, noteInteractor *domain.NoteInteractor,
) *CommandHandler {
	return &CommandHandler{NewEventBasedNoteRepo(eventStore), eventStore, noteInteractor}
}

func (h *CommandHandler) CreateNote(cmd commands.CreateNoteCommand) {
	evts := h.noteInteractor.CreateNote(cmd.Title, cmd.Content)
	h.publishEvents(evts)
}

func (h *CommandHandler) UpdateNote(cmd commands.UpdateNoteCommand) {
	note := h.repo.GetNoteByID(cmd.ID)
	evts := h.noteInteractor.UpdateNote(note, cmd.Title, cmd.Content)
	h.publishEvents(evts)
}

func (h *CommandHandler) publishEvents(evts []cqrs.Event) {
	for _, event := range evts {
		h.eventStore.AddEvent(event)
	}
}
