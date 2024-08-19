package services

import (
	"context"
	models "tender/internal/models/dmart" // Ensure this is correct
	repositories_dmart "tender/internal/repositories/dmart"
)

type InvestmentOilProductionService struct {
	Repo *repositories_dmart.InvestmentOilProductionRepository
}

type InvestmentReservesService struct {
	Repo *repositories_dmart.InvestmentReservesRepository
}

// GetInvestmentOilProductionSummary retrieves a summary of investment oil production for a specific year.
func (s *InvestmentOilProductionService) GetInvestmentOilProductionSummary(ctx context.Context, year int, unit string) ([]models.InvestmentOilProductionSummary, error) {
	summary, err := s.Repo.GetInvestmentOilProductionSummary(ctx, year, unit)
	if err != nil {
		return nil, err
	}
	return summary, nil
}

func (s *InvestmentReservesService) GetInvestmentReservesSummary(ctx context.Context, year int, unit string) ([]models.InvestmentReservesSummary, error) {
	summary, err := s.Repo.GetInvestmentReservesSummary(ctx, year, unit)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
