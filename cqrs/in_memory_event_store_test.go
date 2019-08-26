package cqrs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryEventStore(t *testing.T) {
	t.Run("ReadEventsByID", func(t *testing.T) {
		t.Run("Returns empty when no event is added", func(t *testing.T) {
			store := NewInMemoryEventStore()
			evts, err := store.ReadEventsByID("id")

			assert.Equal(t, []Event{}, evts)
			assert.Equal(t, nil, err)
		})

		t.Run("Returns all events when all events belongs to the same ID", func(t *testing.T) {
			store := NewInMemoryEventStore()

			store.AddEvent(FakeEvent{id: "id", value: "value 1"})
			store.AddEvent(FakeEvent{id: "id", value: "value 2"})

			evts, err := store.ReadEventsByID("id")

			assert.Equal(t, []Event{
				FakeEvent{id: "id", value: "value 1"},
				FakeEvent{id: "id", value: "value 2"},
			}, evts)
			assert.Equal(t, nil, err)
		})

		t.Run("Returns events filtered by ID when the events belongs to more than one ID", func(t *testing.T) {
			store := NewInMemoryEventStore()

			store.AddEvent(FakeEvent{id: "1", value: "value 1"})
			store.AddEvent(FakeEvent{id: "1", value: "value 2"})
			store.AddEvent(FakeEvent{id: "2", value: "value 3"})

			evts, err := store.ReadEventsByID("1")

			assert.Equal(t, []Event{
				FakeEvent{id: "1", value: "value 1"},
				FakeEvent{id: "1", value: "value 2"},
			}, evts)
			assert.Equal(t, nil, err)

			evts, _ = store.ReadEventsByID("2")
			assert.Equal(t, []Event{
				FakeEvent{id: "2", value: "value 3"},
			}, evts)

			evts, _ = store.ReadEventsByID("3")
			assert.Equal(t, []Event{}, evts)
		})
	})
}

type FakeEvent struct {
	id    string
	value string
}

func (e FakeEvent) GetID() string {
	return e.id
}
