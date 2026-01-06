package usecase

import (
	"context"

	"github.com/daniyar23/crm/internal/core/domain"
)

// интерфейс — usecase не знает, КАК реализован сервис
type CompanyService interface {
	CreateCompany(context.Context, *domain.Company) (*domain.Company, error)
	GetCompaniesByUser(context.Context, uint) ([]domain.Company, error)
	DeleteCompany(context.Context, uint) error
	DeleteCompaniesByUser(context.Context, uint) error
}

type CompanyUseCase struct {
	service CompanyService
}

func NewCompanyUseCase(service CompanyService) *CompanyUseCase {
	return &CompanyUseCase{service: service}
}

func (uc *CompanyUseCase) CreateCompany(
	ctx context.Context,
	company *domain.Company,
) (*domain.Company, error) {
	return uc.service.CreateCompany(ctx, company)
}

func (uc *CompanyUseCase) GetCompaniesByUser(
	ctx context.Context,
	userID uint,
) ([]domain.Company, error) {
	return uc.service.GetCompaniesByUser(ctx, userID)
}

func (uc *CompanyUseCase) DeleteCompany(
	ctx context.Context,
	id uint,
) error {
	return uc.service.DeleteCompany(ctx, id)
}

// используется слушателем события UserDeleted
func (uc *CompanyUseCase) DeleteCompaniesByUser(
	ctx context.Context,
	userID uint,
) error {
	return uc.service.DeleteCompaniesByUser(ctx, userID)
}
