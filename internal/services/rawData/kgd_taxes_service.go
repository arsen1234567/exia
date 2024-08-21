package services

import (
	"context"
	models "tender/internal/models/rawData" // Ensure this is correct
	repositories "tender/internal/repositories/rawData"
)

type KgdTaxesService struct {
	Repo *repositories.KgdTaxesRepository
}

func (s *KgdTaxesService) GetKgdTaxesSummary(ctx context.Context, year int, currency string) ([]models.KgdTaxesSummary, error) {
	summary, err := s.Repo.GetKgdTaxesSummary(ctx, year, currency)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
