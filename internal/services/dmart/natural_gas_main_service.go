package services

import (
	"context"
	repositories_dmart "tender/internal/repositories/dmart"
)

type NaturalGasMainService struct {
	Repo *repositories_dmart.NaturalGasMainRepository
}

func (s *NaturalGasMainService) GetReserveRatio(ctx context.Context) (int64, error) {
	reserveRatio, err := s.Repo.GetReserveRatio(ctx)
	if err != nil {
		return 0, err
	}
	return reserveRatio, nil
}
