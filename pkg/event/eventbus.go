package event

const (
	LinkVisitedEvent = "link_visited"
)

type Event struct {
	Type string
	Data any
}

type EventBus struct {
	bus chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{make(chan Event)}
}

func (b *EventBus) Publish(event Event) {
	b.bus <- event
}

func (b *EventBus) Subscribe() <-chan Event {
	return b.bus
}
