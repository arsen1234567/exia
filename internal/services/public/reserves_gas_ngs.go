package services

import (
	"context"
	models "tender/internal/models/rawData"
	repositories "tender/internal/repositories/public"
)

type ReservesGasNgsService struct {
	Repo *repositories.ReservesGasNgsRepository
}

func (s *ReservesGasNgsService) GetDeposit(ctx context.Context) ([]models.NgsReservesGasSummary, error) {
	return s.Repo.GetReservesGasNgsDeposit(ctx)
}

func (s *ReservesGasNgsService) GetNumberOfCompanies(ctx context.Context) ([]models.NgsReservesGasSummary, error) {
	return s.Repo.GetReservesGasNgsNumberOfCompanies(ctx)
}

func (s *ReservesGasNgsService) GetTotalReserves(ctx context.Context, oilType string) ([]models.NgsReservesGasTotalReservesSummary, error) {
	return s.Repo.GetTotalReserves(ctx, oilType)
}

func (s *ReservesGasNgsService) GetProduction(ctx context.Context, oilType string) ([]models.NgsReservesGasTotalReservesSummary, error) {
	return s.Repo.GetProduction(ctx, oilType)
}

func (s *ReservesGasNgsService) GetNumberOfDepositsByRegion(ctx context.Context, year int) ([]models.DepositsByRegionSummary, error) {
	return s.Repo.GetNumberOfDepositsByRegion(ctx, year)
}

func (s *ReservesGasNgsService) GetTopCompaniesByReserves(ctx context.Context, year int, oilType string) ([]models.NgsReservesGasTopCompanies, error) {
	return s.Repo.GetTopCompaniesByReserves(ctx, oilType, year)
}
