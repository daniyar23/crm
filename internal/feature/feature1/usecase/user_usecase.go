package usecase

import (
	"context"
	"errors"

	"github.com/daniyar23/crm/internal/core/domain"
	"github.com/daniyar23/crm/internal/feature/feature1/usecase/events"
)

// UserService — интерфейс бизнес-логики.
// Реализация будет в service-слое.
type UserService interface {
	GetUserByID(ctx context.Context, id uint) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id uint) error
}

// UserUseCase — слой usecase (application layer).
// Отвечает за сценарии, а не за реализацию.
type UserUseCase struct {
	userService UserService
	eventBus    EventBus
}

// Конструктор
func NewUserUseCase(userService UserService, eventBus EventBus) *UserUseCase {
	return &UserUseCase{
		userService: userService,
		eventBus:    eventBus,
	}
}

// CreateUser — сценарий создания пользователя
func (u *UserUseCase) CreateUser(
	ctx context.Context,
	user *domain.User,
) (*domain.User, error) {

	if user.Email == "" || user.Name == "" {
		return nil, errors.New("email and name are required")
	}

	if _, err := u.userService.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID — сценарий получения пользователя
func (u *UserUseCase) GetUserByID(
	ctx context.Context,
	id uint,
) (*domain.User, error) {

	if id == 0 {
		return nil, errors.New("invalid user id")
	}

	return u.userService.GetUserByID(ctx, id)
}

// GetAllUsers — сценарий получения списка пользователей
func (u *UserUseCase) GetAllUsers(
	ctx context.Context,
) ([]domain.User, error) {

	return u.userService.GetAllUsers(ctx)
}

// DeleteUser — сценарий удаления пользователя
func (u *UserUseCase) DeleteUser(
	ctx context.Context,
	id uint,
) error {

	if id == 0 {
		return errors.New("invalid user id")
	}

	if err := u.userService.DeleteUser(ctx, id); err != nil {
		return err
	}

	u.eventBus.Publish(events.UserDeleted{
		UserID: id,
	})

	return nil
}
