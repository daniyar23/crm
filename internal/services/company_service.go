package services

import (
	"errors"

	"github.com/daniyar23/crm/internal/domain"
	"github.com/daniyar23/crm/internal/infrastructure/repository"
)

type CompanyService struct {
	repo repository.CompanyRepository
}

func NewCompanyService(repo repository.CompanyRepository) *CompanyService {
	return &CompanyService{repo: repo}
}

func (s *CompanyService) Create(company *domain.Company) (uint, error) {
	if company.Name == "" {
		return 0, errors.New("company name cannot be empty")
	}
	if company.User_ID == 0 {
		return 0, errors.New("user id cannot be empty")
	}
	return s.repo.CreateCompany(company)
}

func (s *CompanyService) GetByID(id uint) (*domain.Company, error) {
	if id == 0 {
		return nil, errors.New("invalid company id")
	}
	return s.repo.GetCompanyByID(id)
}

func (s *CompanyService) GetAll() ([]domain.Company, error) {
	return s.repo.GetAllCompanies()
}

func (s *CompanyService) Delete(id uint) error {
	if id == 0 {
		return errors.New("invalid company id")
	}
	return s.repo.DeleteCompany(id)
}
