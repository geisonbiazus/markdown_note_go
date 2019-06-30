package cqrs

type InMemoryEventStore struct {
	Events []Event
}

func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{
		Events: []Event{},
	}
}

func (s *InMemoryEventStore) AddEvent(event Event) error {
	s.Events = append(s.Events, event)
	return nil
}

func (s *InMemoryEventStore) ReadEvents() ([]Event, error) {
	return s.Events, nil
}
