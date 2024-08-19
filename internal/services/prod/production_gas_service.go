package services

import (
	"context"
	models "tender/internal/models/prod" // Ensure this is correct
	repositories "tender/internal/repositories/prod"
)

type ProductionGasService struct {
	Repo *repositories.ProductionGasRepository
}

func (s *ProductionGasService) GetGasProductionSummary(ctx context.Context, year int) ([]models.ProductionGasSummary, error) {
	summary, err := s.Repo.GetGasProductionSummary(ctx, year)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
