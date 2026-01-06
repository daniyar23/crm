package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/daniyar23/crm/internal/core/domain"
)

type CompanyPostgresRepository struct {
	db *sql.DB
}

func NewCompanyPostgresRepository(db *sql.DB) *CompanyPostgresRepository {
	return &CompanyPostgresRepository{db: db}
}
func (r *CompanyPostgresRepository) CreateCompany(
	ctx context.Context,
	company *domain.Company,
) (*domain.Company, error) {

	query := `
		INSERT INTO companies (name, user_id)
		VALUES ($1, $2)
		RETURNING id
	`

	var id uint
	err := r.db.QueryRowContext(
		ctx,
		query,
		company.Name,
		company.UserID,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &domain.Company{
		ID:     id,
		Name:   company.Name,
		UserID: company.UserID,
	}, nil
}
func (r *CompanyPostgresRepository) GetCompanyByID(
	ctx context.Context,
	id uint,
) (*domain.Company, error) {

	query := `
		SELECT id, name, user_id
		FROM companies
		WHERE id = $1
	`

	var company domain.Company

	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&company.ID, &company.Name, &company.UserID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return &company, nil
}
func (r *CompanyPostgresRepository) GetAllCompanies(
	ctx context.Context,
) ([]domain.Company, error) {

	query := `
		SELECT id, name, user_id
		FROM companies
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []domain.Company

	for rows.Next() {
		var c domain.Company
		err := rows.Scan(&c.ID, &c.Name, &c.UserID)
		if err != nil {
			return nil, err
		}
		companies = append(companies, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
func (r *CompanyPostgresRepository) DeleteCompany(
	ctx context.Context,
	id uint,
) error {

	query := `
		DELETE FROM companies
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
