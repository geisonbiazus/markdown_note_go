package cqrs

type Event interface{}

type EventStore interface {
	AddEvent(event Event) error
	ReadEvents() ([]Event, error)
}
