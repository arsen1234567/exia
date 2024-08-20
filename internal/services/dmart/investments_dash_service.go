package services

import (
	"context" // Ensure this is correct
	repositories_dmart "tender/internal/repositories/dmart"
)

type InvestmentsDashService struct {
	Repo *repositories_dmart.InvestmentsDashRepository
}

func (s *InvestmentsDashService) GetInvestmentsDash(ctx context.Context) (float64, error) {
	summary, err := s.Repo.GetInvestmentsDash(ctx)
	if err != nil {
		return 0, err
	}
	return summary, nil
}
