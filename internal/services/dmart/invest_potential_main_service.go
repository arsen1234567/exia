package services

import (
	"context" // Ensure this is correct
	repositories_dmart "tender/internal/repositories/dmart"
)

type InvestPotentialMainService struct {
	Repo *repositories_dmart.InvestPotentialMainRepository
}

func (s *InvestPotentialMainService) GetInvestPotentialMain(ctx context.Context) (float64, error) {
	summary, err := s.Repo.GetInvestPotentialMain(ctx)
	if err != nil {
		return 0, err
	}
	return summary, nil
}

func (s *InvestPotentialMainService) GetSpecOpEx(ctx context.Context) (float64, error) {
	summary, err := s.Repo.GetSpecOpEx(ctx)
	if err != nil {
		return 0, err
	}
	return summary, nil
}
