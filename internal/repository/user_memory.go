package repository

import (
	"errors"
	"sync"

	"github.com/daniyar23/crm/internal/domain"
)

type UserMemoryRepository struct {
	mu     sync.Mutex
	users  map[uint]domain.User
	nextID uint
}

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		users:  make(map[uint]domain.User),
		nextID: 1,
	}
}

func (r *UserMemoryRepository) CreateUser(user domain.User) (domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.nextID++

	r.users[user.ID] = user
	return user, nil
}

func (r *UserMemoryRepository) GetUserByID(id uint) (domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.users[id]
	if !ok {
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r *UserMemoryRepository) GetAllUsers() ([]domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]domain.User, 0, len(r.users))
	for _, u := range r.users {
		result = append(result, u)
	}

	return result, nil
}

func (r *UserMemoryRepository) DeleteUser(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}
