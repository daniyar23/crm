package repository

import (
	"context"

	"github.com/daniyar23/crm/internal/core/domain"
)

type CompanyRepository interface {
	CreateCompany(ctx context.Context, company *domain.Company) (*domain.Company, error)

	GetCompaniesByUser(ctx context.Context, userID uint) ([]domain.Company, error)
	DeleteCompaniesByUser(ctx context.Context, userID uint) error

	DeleteCompany(ctx context.Context, id uint) error
}
