package services

import (
	"context"
	repositories "tender/internal/repositories/prod"
)

type GasStepsService struct {
	Repo *repositories.GasStepsRepository
}

// GetTotalKpnByUsername retrieves the total KPN for a given username using the repository
func (s *GasStepsService) GetAmountOfPredictedTaxes(ctx context.Context, currency string) (float64, error) {
	return s.Repo.GetAmountOfPredictedTaxes(ctx, currency)
}

func (s *GasStepsService) GetEBITDAmargin(ctx context.Context) (float64, error) {
	return s.Repo.GetEBITDAmargin(ctx)
}
