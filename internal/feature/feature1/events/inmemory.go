package events

type InMemoryBus struct {
	ch chan any
}

func NewInMemoryBus() *InMemoryBus {
	return &InMemoryBus{
		ch: make(chan any, 16),
	}
}

func (b *InMemoryBus) Publish(event any) {
	b.ch <- event
}

func (b *InMemoryBus) Subscribe() <-chan any {
	return b.ch
}
