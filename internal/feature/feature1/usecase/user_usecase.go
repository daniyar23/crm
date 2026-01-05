package usecase

import "context"

// UserService — контракт бизнес-операций с пользователем.
// Реализация живёт в services-слое.
type UserService interface {
	Create(ctx context.Context, email, name string) (uint, error)
}

// UserUseCase — оркестратор бизнес-флоу.
type UserUseCase struct {
	userService UserService
}

// NewUserUseCase — конструктор. Создаётся в main (composition root).
func NewUserUseCase(userService UserService) *UserUseCase {
	return &UserUseCase{
		userService: userService,
	}
}

// CreateUser — usecase создания пользователя.
// Никаких транспортных деталей.
func (uc *UserUseCase) CreateUser(
	ctx context.Context,
	input CreateUserInput,
) (uint, error) {

	// TODO:
	// 1. валидация input
	// 2. вызов uc.userService.Create
	// 3. публикация события (через channel) — позже

	return uc.userService.Create(ctx, input.Email, input.Name)
}
