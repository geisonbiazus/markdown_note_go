package notes

import (
	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/commands"
	"github.com/geisonbiazus/markdown_notes/notes/domain"
)

type CommandHandler struct {
	eventStore  cqrs.EventStore
	idGenerator IDGenerator
}

func NewCommandHandler(eventStore cqrs.EventStore, idGenerator IDGenerator) *CommandHandler {
	return &CommandHandler{eventStore: eventStore, idGenerator: idGenerator}
}

func (h *CommandHandler) CreateNote(command commands.CreateNoteCommand) {
	_, evts := domain.CreateNote(h.idGenerator.Generate(), command.Title, command.Content)

	for _, event := range evts {
		h.eventStore.AddEvent(event)
	}
}
