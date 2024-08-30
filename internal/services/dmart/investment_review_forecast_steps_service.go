package services

import (
	"context" // Ensure this is correct
	models "tender/internal/models/dmart"
	repositories_dmart "tender/internal/repositories/dmart"
)

type InvestmentReviewForecastStepsService struct {
	Repo *repositories_dmart.InvestmentReviewForecastStepsRepository
}

func (s *InvestmentReviewForecastStepsService) GetInvestmentReviewForecastSteps(ctx context.Context, currency string) (float64, error) {
	summary, err := s.Repo.GetInvestmentReviewForecastSteps(ctx, currency)
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

func (s *InvestmentReviewForecastStepsService) GetCompaniesForecastStepsSummary(ctx context.Context, unit string) ([]models.InvestmentReviewForecastStepsSummary, error) {
	summary, err := s.Repo.GetCompaniesForecastSteps(ctx, unit)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
