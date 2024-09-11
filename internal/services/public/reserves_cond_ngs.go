package services

import (
	"context"
	models "tender/internal/models/rawData"
	repositories "tender/internal/repositories/public"
)

type ReservesCondNgsService struct {
	Repo *repositories.ReservesCondNgsRepository
}

func (s *ReservesCondNgsService) GetDeposit(ctx context.Context) ([]models.NgsReservesCondSummary, error) {
	return s.Repo.GetReservesCondNgsDeposit(ctx)
}

func (s *ReservesCondNgsService) GetNumberOfCompanies(ctx context.Context) ([]models.NgsReservesCondSummary, error) {
	return s.Repo.GetReservesCondNgsNumberOfCompanies(ctx)
}

func (s *ReservesCondNgsService) GetTotalReserves(ctx context.Context, oilType string) ([]models.NgsReservesCondTotalReservesSummary, error) {
	return s.Repo.GetTotalReserves(ctx, oilType)
}

func (s *ReservesCondNgsService) GetProduction(ctx context.Context, oilType string) ([]models.NgsReservesCondTotalReservesSummary, error) {
	return s.Repo.GetProduction(ctx, oilType)
}

func (s *ReservesCondNgsService) GetNumberOfDepositsByRegion(ctx context.Context, year int) ([]models.DepositsByRegionSummary, error) {
	return s.Repo.GetNumberOfDepositsByRegion(ctx, year)
}


func (s *ReservesCondNgsService) GetTopCompaniesByReserves(ctx context.Context, year int, oilType string) ([]models.NgsReservesCondTopCompanies, error) {
	return s.Repo.GetTopCompaniesByReserves(ctx, oilType, year)
}
