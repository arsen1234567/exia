package services

import (
	"context"
	models "tender/internal/models/dmart" // Ensure this is correct
	repositories_dmart "tender/internal/repositories/dmart"
)

type InvestmentNetProfitService struct {
	Repo *repositories_dmart.InvestmentNetProfitRepository
}

// GetInvestmentOilProductionSummary retrieves a summary of investment oil production for a specific year.
func (s *InvestmentNetProfitService) GetInvestmentNetProfitSummary(ctx context.Context, year int) ([]models.InvestmentNetProfitSummary, error) {
	summary, err := s.Repo.GetInvestmentNetProfitSummary(ctx, year)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
