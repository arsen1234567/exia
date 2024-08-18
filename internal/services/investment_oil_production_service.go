package services

import (
	"context"
	models "tender/internal/models/dmart" // Ensure this is correct
	repositories "tender/internal/repositories/dmart"
)

type InvestmentOilProductionService struct {
	Repo *repositories.InvestmentOilProductionRepository
}

// GetInvestmentOilProductionSummary retrieves a summary of investment oil production for a specific year.
func (s *InvestmentOilProductionService) GetInvestmentOilProductionSummary(ctx context.Context, year int) ([]models.InvestmentOilProductionSummary, error) {
	summary, err := s.Repo.GetInvestmentOilProductionSummary(ctx, year)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
