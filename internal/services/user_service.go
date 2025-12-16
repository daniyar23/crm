package services

// Service может работать без HTTP.
// Service может работать без PostgreSQL.
// Но Service НЕ может работать без Domain и Repository
import (
	"errors"

	"github.com/daniyar23/crm/internal/domain"
	"github.com/daniyar23/crm/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user domain.User) (domain.User, error) {
	if user.Email == "" {
		return domain.User{}, errors.New("email is required")
	}
	if user.Name == "" {
		return domain.User{}, errors.New("name is required")
	}

	return s.repo.CreateUser(user)
}

func (s *UserService) GetByID(id uint) (domain.User, error) {
	if id == 0 {
		return domain.User{}, errors.New("invalid user id")
	}

	return s.repo.GetUserByID(id)
}

func (s *UserService) GetAll() ([]domain.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) Delete(id uint) error {
	if id == 0 {
		return errors.New("invalid user id")
	}
	return s.repo.DeleteUser(id)
}
