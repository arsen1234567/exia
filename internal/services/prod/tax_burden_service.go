package services

import (
	"context" // Ensure this is correct
	repositories "tender/internal/repositories/prod"
)

type TaxBurdenService struct {
	Repo *repositories.TaxBurdenRepository
}

func (s *TaxBurdenService) GetTaxBurden(ctx context.Context, year int, currency string) (map[string]float64, error) {
	totalSumSummary, err := s.Repo.GetTaxBurden(ctx, year, currency)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}
