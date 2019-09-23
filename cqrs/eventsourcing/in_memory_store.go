package eventsourcing

type Stream struct {
	events    []Event
	callbacks []func(Event)
}

func NewStream() *Stream {
	return &Stream{events: []Event{}, callbacks: []func(Event){}}
}

func (s *Stream) Publish(event Event) error {
	s.events = append(s.events, event)

	if s.callbacks != nil {
		for _, callback := range s.callbacks {
			callback(event)
		}
	}
	return nil
}

func (s *Stream) Stream(fromVersion uint64, callback func(Event)) error {
	s.callbacks = append(s.callbacks, callback)

	if s.callbacks != nil {
		for _, evt := range s.events {
			callback(evt)
		}
	}

	return nil
}

type InMemoryStore struct {
	streams map[string]*Stream
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{streams: make(map[string]*Stream)}
}

func (s *InMemoryStore) Publish(streamName string, event Event) error {
	return s.getStream(streamName).Publish(event)
}

func (s *InMemoryStore) Stream(streamName string, fromVersion uint64, callback func(Event)) error {
	return s.getStream(streamName).Stream(fromVersion, callback)
}

func (s *InMemoryStore) getStream(name string) *Stream {
	stream, exists := s.streams[name]
	if !exists {
		stream = NewStream()
		s.streams[name] = stream
	}
	return stream
}
