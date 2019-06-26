package notes

import (
	"github.com/geisonbiazus/markdown_notes/cqrs"
	"github.com/geisonbiazus/markdown_notes/notes/commands"
	"github.com/geisonbiazus/markdown_notes/notes/domain"
	"github.com/geisonbiazus/markdown_notes/notes/events"
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

func (h *CommandHandler) CreateNote(cmd commands.CreateNoteCommand) {
	evts := h.noteInteractor.CreateNote(cmd.Title, cmd.Content)
	h.publishEvents(evts)
}

func (h *CommandHandler) UpdateNote(cmd commands.UpdateNoteCommand) {
	note := h.loadNote(cmd.ID)
	evts := h.noteInteractor.UpdateNote(note, cmd.Title, cmd.Content)
	h.publishEvents(evts)
}

func (h *CommandHandler) publishEvents(evts []cqrs.Event) {
	for _, event := range evts {
		h.eventStore.AddEvent(event)
	}
}

func (h *CommandHandler) loadNote(id string) domain.Note {
	store := h.eventStore.(*cqrs.InMemoryEventStore)
	evt := store.Events[0].(events.NoteCreatedEvent)

	return domain.Note{ID: evt.ID, Title: evt.Title, Content: evt.Content}
}
