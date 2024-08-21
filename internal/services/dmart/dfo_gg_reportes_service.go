package services

import (
	"context"
	models "tender/internal/models/dmart" // Ensure this is correct
	repositories_dmart "tender/internal/repositories/dmart"
)

type DfoGgReportesService struct {
	Repo *repositories_dmart.DfoGgReportesRepository
}

func (s *DfoGgReportesService) GetNetProfitSummary(ctx context.Context) ([]models.NetProfitSummary, error) {
	summary, err := s.Repo.GetNetProfitSummary(ctx)
	if err != nil {
		return nil, err
	}
	return summary, nil
}

func (s *DfoGgReportesService) GetRevenueByCompanyAndYear(ctx context.Context, company string, year int) (float64, error) {
	return s.Repo.GetRevenueByCompanyAndYear(ctx, company, year)
}

func (s *DfoGgReportesService) GetCostOfGoodsWorksServicesSold(ctx context.Context, company string, year int) (float64, error) {
	totalCost, err := s.Repo.GetCostOfGoodsWorksServicesSold(ctx, company, year)
	if err != nil {
		return 0, err
	}
	return totalCost, nil
}

func (s *DfoGgReportesService) GetGrossProfit(ctx context.Context, company string, year int) (float64, error) {
	grossProfit, err := s.Repo.GetGrossProfit(ctx, company, year)
	if err != nil {
		return 0, err
	}
	return grossProfit, nil
}

func (s *DfoGgReportesService) GetCIT(ctx context.Context, company string, year int) (float64, error) {
	totalCIT, err := s.Repo.GetCIT(ctx, company, year)
	if err != nil {
		return 0, err
	}
	return totalCIT, nil
}
