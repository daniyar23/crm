package repository

import (
	"context"

	"github.com/daniyar23/crm/internal/core/domain"
)

type CompanyRepository interface {
	CreateCompany(ctx context.Context, company *domain.Company) (uint, error)
	GetCompanyByID(ctx context.Context, id uint) (*domain.Company, error)
	GetAllCompanies(ctx context.Context) ([]domain.Company, error)
	DeleteCompany(ctx context.Context, id uint) error
}
