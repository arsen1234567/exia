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

func (s *InvestmentReviewForecastStepsService) GetRevenueForecastStepsSummary(ctx context.Context, unit, currency string) ([]models.RevenueByYear, error) {
	summary, err := s.Repo.GetRevenueInvestmentReviewForecastSteps(ctx, unit, currency)
	if err != nil {
		return nil, err
	}
	return summary, nil
}

func (s *InvestmentReviewForecastStepsService) GetCapExForecastStepsSummary(ctx context.Context, unit, currency string) ([]models.InvestmentReviewForecastData, error) {
	summary, err := s.Repo.GetCapExInvestmentReviewForecastSteps(ctx, unit, currency)
	if err != nil {
		return nil, err
	}
	return summary, nil
}

func (s *InvestmentReviewForecastStepsService) GetATFCFForecastStepsSummary(ctx context.Context, unit, currency string) ([]models.InvestmentReviewForecastData, error) {
	summary, err := s.Repo.GetATFCFInvestmentReviewForecastSteps(ctx, unit, currency)
	if err != nil {
		return nil, err
	}
	return summary, nil
}

func (s *InvestmentReviewForecastStepsService) GetOpExForecastStepsSummary(ctx context.Context, unit, currency string) ([]models.InvestmentReviewForecastData, error) {
	summary, err := s.Repo.GetOpExInvestmentReviewForecastSteps(ctx, unit, currency)
	if err != nil {
		return nil, err
	}
	return summary, nil
}

func (s *InvestmentReviewForecastStepsService) GetGovShareForecastStepsSummary(ctx context.Context, unit, currency string) ([]models.InvestmentReviewForecastData, error) {
	summary, err := s.Repo.GetGovShareInvestmentReviewForecastSteps(ctx, unit, currency)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
