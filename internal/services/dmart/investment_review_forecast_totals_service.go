package services

import (
	"context" // Ensure this is correct
	repositories_dmart "tender/internal/repositories/dmart"
)

type InvestmentReviewForecastTotalService struct {
	Repo *repositories_dmart.InvestmentReviewForecastTotalRepository
}

func (s *InvestmentReviewForecastTotalService) GetInvestmentReviewForecastTotal(ctx context.Context) (float64, error) {
	summary, err := s.Repo.GetInvestmentReviewForecastTotal(ctx)
	if err != nil {
		return 0, err
	}
	return summary, nil
}
