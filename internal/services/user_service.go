package services

// Service может работать без HTTP.
// Service может работать без PostgreSQL.
// Но Service НЕ может работать без Domain и Repository

import (
	"context"
	"errors"

	"github.com/daniyar23/crm/internal/domain"
	"github.com/daniyar23/crm/internal/infrastructure/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if user.Email == "" {
		return nil, errors.New("email is required")
	}
	if user.Name == "" {
		return nil, errors.New("name is required")
	}

	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	if id == 0 {
		return nil, errors.New("invalid user id")
	}

	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.GetAllUsers(ctx)
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid user id")
	}
	return s.repo.DeleteUser(ctx, id)
}
