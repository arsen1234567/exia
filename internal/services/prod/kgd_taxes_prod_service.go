package services

import (
	"context"
	models "tender/internal/models/prod" // Ensure this is correct
	repositories "tender/internal/repositories/prod"
)

type KgdTaxesProdService struct {
	Repo *repositories.KgdTaxesProdRepository
}

func (s *KgdTaxesProdService) GetKgdTaxesProdSummary(ctx context.Context, year int) ([]models.KgdTaxesProdSummary, error) {
	summary, err := s.Repo.GetKgdTaxesProdSummary(ctx, year)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
