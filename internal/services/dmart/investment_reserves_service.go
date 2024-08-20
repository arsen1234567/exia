package services

import (
	"context"
	models "tender/internal/models/dmart" // Ensure this is correct
	repositories_dmart "tender/internal/repositories/dmart"
)

type InvestmentReservesService struct {
	Repo *repositories_dmart.InvestmentReservesRepository
}

func (s *InvestmentReservesService) GetInvestmentReservesSummary(ctx context.Context, year int, unit string) ([]models.InvestmentReservesSummary, error) {
	summary, err := s.Repo.GetInvestmentReservesSummary(ctx, year, unit)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
