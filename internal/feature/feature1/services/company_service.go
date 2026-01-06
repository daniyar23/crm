package services

import (
	"context"
	"errors"

	"github.com/daniyar23/crm/internal/core/domain"
	"github.com/daniyar23/crm/internal/feature/feature1/infrastructure/repository"
)

type CompanyService struct {
	repo repository.CompanyRepository
}

func NewCompanyService(repo repository.CompanyRepository) *CompanyService {
	return &CompanyService{repo: repo}
}

func (s *CompanyService) CreateCompany(ctx context.Context, company *domain.Company) (*domain.Company, error) {
	if company.Name == "" {
		return nil, errors.New("company name cannot be empty")
	}
	if company.UserID == 0 {
		return nil, errors.New("user id cannot be empty")
	}
	return s.repo.CreateCompany(ctx, company)
}

func (s *CompanyService) GetCompanyByID(ctx context.Context, id uint) (*domain.Company, error) {
	if id == 0 {
		return nil, errors.New("invalid company id")
	}
	return s.repo.GetCompanyByID(ctx, id)
}

func (s *CompanyService) GetAllCompanies(ctx context.Context) ([]domain.Company, error) {
	return s.repo.GetAllCompanies(ctx)
}

func (s *CompanyService) DeleteCompany(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid company id")
	}
	return s.repo.DeleteCompany(ctx, id)
}
