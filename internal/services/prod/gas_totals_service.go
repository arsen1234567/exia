package services

import (
	"context"
	repositories "tender/internal/repositories/prod" // Ensure this import path is correct
)

type GasTotalsService struct {
	Repo *repositories.GasTotalsRepository
}

// GetNPVplusTV retrieves the sum of NPV plus Terminal Value for a specific user
func (s *GasTotalsService) GetNPVplusTV(ctx context.Context, currency string) (float64, error) {
	totalNPVplusTV, err := s.Repo.GetNPVplusTV(ctx, currency)
	if err != nil {
		return 0, err
	}
	return totalNPVplusTV, nil
}
