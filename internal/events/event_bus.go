package events

type EventBus interface {
	Publish(event any)
	Subscribe() <-chan any
}
