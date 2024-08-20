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
