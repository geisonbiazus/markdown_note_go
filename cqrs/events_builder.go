package cqrs

type EventsBuilder struct {
	Events []Event
}

func NewEventsBuilder() *EventsBuilder {
	return &EventsBuilder{[]Event{}}
}

func (b *EventsBuilder) Add(event Event) {
	b.Events = append(b.Events, event)
}
