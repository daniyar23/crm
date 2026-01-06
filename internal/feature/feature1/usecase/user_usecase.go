package usecase

import (
	"context"
	"errors"

	"github.com/daniyar23/crm/internal/core/domain"
)

// UserService — интерфейс бизнес-логики.
// Реализация будет в service-слое.
type UserService interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id uint) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
	Delete(ctx context.Context, id uint) error
}

// UserUseCase — слой usecase (application layer).
// Отвечает за сценарии, а не за реализацию.
type UserUseCase struct {
	userService UserService
}

// Конструктор
func NewUserUseCase(userService UserService) *UserUseCase {
	return &UserUseCase{
		userService: userService,
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

	if err := u.userService.Create(ctx, user); err != nil {
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

	return u.userService.GetByID(ctx, id)
}

// GetAllUsers — сценарий получения списка пользователей
func (u *UserUseCase) GetAllUsers(
	ctx context.Context,
) ([]*domain.User, error) {

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

	return u.userService.Delete(ctx, id)
}
