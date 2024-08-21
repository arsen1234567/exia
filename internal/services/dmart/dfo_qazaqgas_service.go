package services

import (
	"context"
	models "tender/internal/models/dmart"
	repositories "tender/internal/repositories/dmart"
)

type DfoQazaqgasService struct {
	Repo *repositories.DfoQazaqgasRepository
}

// GetRevenueByServicesAndCompanyAndYear retrieves the service revenue by company and year
func (s *DfoQazaqgasService) GetRevenueByServicesAndCompanyAndYear(ctx context.Context, company string, year int) (models.RevenueByServicesSummary, error) {
	return s.Repo.GetRevenueByServicesAndCompanyAndYear(ctx, company, year)
}

// GetRevenueByGeographyAndCompanyAndYear retrieves the revenue by geography and company
func (s *DfoQazaqgasService) GetRevenueByGeographyAndCompanyAndYear(ctx context.Context, company string, year int) (models.RevenueByGeographySummary, error) {
	return s.Repo.GetRevenueByGeographyAndCompanyAndYear(ctx, company, year)
}

// GetCostItemsByCompanyAndYear retrieves the cost items by company and year
func (s *DfoQazaqgasService) GetCostItemsByCompanyAndYear(ctx context.Context, company string, year int) (models.CostItemsSummary, error) {
	return s.Repo.GetCostItemsByCompanyAndYear(ctx, company, year)
}
