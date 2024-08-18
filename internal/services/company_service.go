package services

import (
	"context"
	"tender/internal/models"
	"tender/internal/repositories"
)

type CompanyService struct {
	Repo *repositories.CompanyRepository
}

// CreateCompany creates a new company.
func (s *CompanyService) CreateCompany(ctx context.Context, company models.Company) (models.Company, error) {
	id, err := s.Repo.CreateCompany(ctx, company)
	if err != nil {
		return models.Company{}, err
	}

	company.ID = id
	return company, nil
}

// DeleteCompany deletes a company by ID.
func (s *CompanyService) DeleteCompany(ctx context.Context, id int) error {
	return s.Repo.DeleteCompany(ctx, id)
}

// UpdateCompany updates an existing company.
func (s *CompanyService) UpdateCompany(ctx context.Context, company models.Company) (models.Company, error) {
	return s.Repo.UpdateCompany(ctx, company)
}

// GetCompanyByID retrieves a company by ID.
func (s *CompanyService) GetCompanyByID(ctx context.Context, id int) (models.Company, error) {
	return s.Repo.GetCompanyByID(ctx, id)
}

// GetAllCompanies retrieves all companies.
func (s *CompanyService) GetAllCompanies(ctx context.Context) ([]models.Company, error) {
	return s.Repo.GetAllCompanies(ctx)
}
