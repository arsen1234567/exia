package services

import (
	"context"
	rawData "tender/internal/models/rawData"
	repositories "tender/internal/repositories/prod"
)

type KgdTaxesService struct {
	Repo *repositories.KgdTaxesProdRepository
}

func (s *KgdTaxesService) GetKgdTaxesProdSummary(ctx context.Context, year int) ([]rawData.KgdTaxesProdSummary, error) {
	summary, err := s.Repo.GetKgdTaxesProd(ctx, year)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
