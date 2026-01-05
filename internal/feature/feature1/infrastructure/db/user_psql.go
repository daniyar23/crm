package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/daniyar23/crm/internal/core/domain"
)

type UserPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) CreateUser(
	ctx context.Context,
	user *domain.User,
) (*domain.User, error) {

	var id uint

	err := r.db.QueryRowContext(
		ctx,
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
		user.Name,
		user.Email,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:    id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
func (r *UserPostgresRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
func (r *UserPostgresRepository) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	var u domain.User
	row := r.db.QueryRowContext(ctx, "SELECT id, name, email FROM users WHERE id = $1", id)

	err := row.Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}

	return &u, nil
}
func (r *UserPostgresRepository) DeleteUser(ctx context.Context, id uint) error {
	result, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	return nil
}
