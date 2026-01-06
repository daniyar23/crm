package usecase

type EventBus struct {
	ch chan Event
}

func NewEventBus(buffer int) *EventBus {
	return &EventBus{
		ch: make(chan Event, buffer),
	}
}

func (b *EventBus) Publish(e Event) {
	b.ch <- e
}

func (b *EventBus) Subscribe() <-chan Event {
	return b.ch
}
