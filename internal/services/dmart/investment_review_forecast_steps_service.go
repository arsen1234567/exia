package services

import (
	"context" // Ensure this is correct
	models "tender/internal/models/dmart"
	repositories_dmart "tender/internal/repositories/dmart"
)

type InvestmentReviewForecastStepsService struct {
	Repo *repositories_dmart.InvestmentReviewForecastStepsRepository
}

func (s *InvestmentReviewForecastStepsService) GetInvestmentReviewForecastSteps(ctx context.Context) (float64, error) {
	summary, err := s.Repo.GetInvestmentReviewForecastSteps(ctx)
	if err != nil {
		return 0, err
	}
	return summary, nil
}

func (s *InvestmentReviewForecastStepsService) GetEbitdaToGrossRevenueRatio(ctx context.Context) (float64, error) {
	summary, err := s.Repo.GetEbitdaToGrossRevenueRatio(ctx)
	if err != nil {
		return 0, err
	}
	return summary, nil
}

func (s *InvestmentReviewForecastStepsService) GetCompaniesForecastStepsSummary(ctx context.Context, currency, unit string) ([]models.InvestmentReviewForecastStepsSummary, error) {
	summary, err := s.Repo.GetCompaniesForecastSteps(ctx, currency, unit)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
