package usecase

type Event interface{}

type UserDeletedEvent struct {
	UserID uint
}
