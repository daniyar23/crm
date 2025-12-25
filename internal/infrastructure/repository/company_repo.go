package repository

import "github.com/daniyar23/crm/internal/domain"

type CompanyRepository interface {
	CreateCompany(company *domain.Company) (uint, error)
	GetCompanyByID(id uint) (*domain.Company, error)
	GetAllCompanies() ([]domain.Company, error)
	DeleteCompany(id uint) error
}
