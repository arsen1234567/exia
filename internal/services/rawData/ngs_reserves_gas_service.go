package services

import (
	"context"
	models "tender/internal/models/rawData" // Ensure this is correct
	repositories "tender/internal/repositories/rawData"
)

type NgsReservesGasService struct {
	Repo *repositories.NgsReservesGasRepository
}

// GetRecoverableGasReservesSummary retrieves the summary of recoverable gas reserves for the specified year range
func (s *NgsReservesGasService) GetRecoverableGasReservesSummary(ctx context.Context, startYear, endYear int) ([]models.RecoverableGasReservesSummary, error) {
	summary, err := s.Repo.GetRecoverableGasReserves(ctx, startYear, endYear)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
