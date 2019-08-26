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

func (s *InMemoryEventStore) ReadAllEvents() ([]Event, error) {
	return s.Events, nil
}

func (s *InMemoryEventStore) ReadEventsByID(id string) ([]Event, error) {
	evts := []Event{}

	for _, e := range s.Events {
		if e.GetID() == id {
			evts = append(evts, e)
		}
	}

	return evts, nil
}
