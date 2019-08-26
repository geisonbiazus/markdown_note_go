package cqrs

type Event interface {
	GetID() string
}

type EventStore interface {
	AddEvent(event Event) error
	ReadAllEvents() ([]Event, error)
	ReadEventsByID(id string) ([]Event, error)
}
