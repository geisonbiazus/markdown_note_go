package notes

import (
	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/commands"
	"github.com/geisonbiazus/markdown_notes/notes/domain"
)

type CommandHandler struct {
	eventStore     cqrs.EventStore
	noteInteractor *domain.NoteInteractor
}

func NewCommandHandler(
	eventStore cqrs.EventStore, noteInteractor *domain.NoteInteractor,
) *CommandHandler {
	return &CommandHandler{eventStore, noteInteractor}
}

func (h *CommandHandler) CreateNote(command commands.CreateNoteCommand) {
	evts := h.noteInteractor.CreateNote(command.Title, command.Content)
	h.publishEvents(evts)
}

func (h *CommandHandler) publishEvents(evts []cqrs.Event) {
	for _, event := range evts {
		h.eventStore.AddEvent(event)
	}
}
