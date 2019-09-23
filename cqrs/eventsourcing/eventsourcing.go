package eventsourcing

type Event struct {
	ID      string                 `json:id`
	Version uint64                 `json:version`
	Payload map[string]interface{} `json:payload`
}

type Store interface {
	Publish(streamName string, event Event) error
	Stream(streamName string, fromVersion uint64) chan Event
}
