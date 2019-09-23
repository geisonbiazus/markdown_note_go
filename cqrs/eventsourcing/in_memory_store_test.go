package eventsourcing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryStore(t *testing.T) {
	t.Run("Publishing one event", func(t *testing.T) {
		store := NewInMemoryStore()

		event := createEvent(1)

		store.Publish("stream", event)

		receivedEvents := []Event{}

		store.Stream("stream", 1, func(e Event) {
			receivedEvents = append(receivedEvents, e)
		})

		assert.Equal(t, []Event{event}, receivedEvents)
	})

	t.Run("Publishing more than event", func(t *testing.T) {
		store := NewInMemoryStore()

		event1 := createEvent(1)
		event2 := createEvent(2)

		store.Publish("stream", event1)
		store.Publish("stream", event2)

		receivedEvents := []Event{}

		store.Stream("stream", 1, func(e Event) {
			receivedEvents = append(receivedEvents, e)
		})

		assert.Equal(t, []Event{event1, event2}, receivedEvents)
	})

	t.Run("Publishing events after subscription", func(t *testing.T) {
		store := NewInMemoryStore()

		event1 := createEvent(1)
		event2 := createEvent(2)

		receivedEvents := []Event{}

		store.Stream("stream", 1, func(e Event) {
			receivedEvents = append(receivedEvents, e)
		})

		store.Publish("stream", event1)
		store.Publish("stream", event2)

		assert.Equal(t, []Event{event1, event2}, receivedEvents)
	})

	t.Run("Publishing events before and after subscription", func(t *testing.T) {
		store := NewInMemoryStore()

		event1 := createEvent(1)
		event2 := createEvent(2)

		receivedEvents := []Event{}

		store.Publish("stream", event1)

		store.Stream("stream", 1, func(e Event) {
			receivedEvents = append(receivedEvents, e)
		})

		store.Publish("stream", event2)

		assert.Equal(t, []Event{event1, event2}, receivedEvents)
	})

	t.Run("Publishing to more than one subscriber", func(t *testing.T) {
		store := NewInMemoryStore()

		event1 := createEvent(1)
		event2 := createEvent(2)

		receivedEvents1 := []Event{}
		receivedEvents2 := []Event{}

		store.Publish("stream", event1)

		store.Stream("stream", 1, func(e Event) {
			receivedEvents1 = append(receivedEvents1, e)
		})

		store.Stream("stream", 1, func(e Event) {
			receivedEvents2 = append(receivedEvents2, e)
		})

		store.Publish("stream", event2)

		assert.Equal(t, []Event{event1, event2}, receivedEvents1)
		assert.Equal(t, []Event{event1, event2}, receivedEvents2)
	})

	t.Run("Publishing events to different streams", func(t *testing.T) {
		store := NewInMemoryStore()

		event1 := createEvent(1)
		event2 := createEvent(1)
		event3 := createEvent(2)
		event4 := createEvent(2)

		eventsStream1 := []Event{}
		eventsStream2 := []Event{}

		store.Publish("stream1", event1)
		store.Publish("stream2", event2)

		store.Stream("stream1", 1, func(e Event) {
			eventsStream1 = append(eventsStream1, e)
		})

		store.Stream("stream2", 1, func(e Event) {
			eventsStream2 = append(eventsStream2, e)
		})

		store.Publish("stream1", event3)
		store.Publish("stream2", event4)

		assert.Equal(t, []Event{event1, event3}, eventsStream1)
		assert.Equal(t, []Event{event2, event4}, eventsStream2)
	})
}

func createEvent(version int) Event {
	id := uuid4()
	return Event{
		ID:      id,
		Version: uint64(version),
		Payload: map[string]interface{}{"key": id},
	}
}
