package services

import (
	"context"
	models "tender/internal/models/public"
	rawModels "tender/internal/models/rawData"
	repositories "tender/internal/repositories/public"
)

type ReservesOilNgsService struct {
	Repo *repositories.ReservesOilNgsRepository
}

func (s *ReservesOilNgsService) GetDeposit(ctx context.Context) ([]models.NgsReservesOilSummary, error) {
	return s.Repo.GetReservesOilNgsDeposit(ctx)
}

func (s *ReservesOilNgsService) GetNumberOfCompanies(ctx context.Context) ([]models.NgsReservesOilSummary, error) {
	return s.Repo.GetReservesOilNgsNumberOfCompanies(ctx)
}

func (s *ReservesOilNgsService) GetTotalProduction(ctx context.Context, oilType string) ([]rawModels.ReservesOilNgsTotalProductionSummary, error) {
	return s.Repo.GetReservesOilNgsTotalProduction(ctx, oilType)
}

func (s *ReservesOilNgsService) GetNumberOfDepositsByRegion(ctx context.Context, year int) ([]rawModels.DepositsByRegionSummary, error) {
	return s.Repo.GetNumberOfDepositsByRegion(ctx, year)
}

func (s *ReservesOilNgsService) GetTopCompaniesByReserves(ctx context.Context, year int, oilType string) ([]models.NgsReservesOilTopCompanies, error) {
	return s.Repo.GetTopCompaniesByReserves(ctx, oilType, year)
}
