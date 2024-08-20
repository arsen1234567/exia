package services

import (
	"context"
	models "tender/internal/models/rawData"
	repositories "tender/internal/repositories/rawData"
)

type StGasBalanceService struct {
	Repo *repositories.StGasBalanceRepository
}

func (s *StGasBalanceService) GetGasBalance(ctx context.Context, year int64, unit string) ([]models.GasBalanceSummary, error) {
	summaries, err := s.Repo.GetGasBalance(ctx, year)
	if err != nil {
		return nil, err
	}

	for i := range summaries {
		if unit == "ft" {
			// Convert to cubic feet
			summaries[i].TotalGasProduction *= 35.3147
			summaries[i].Import *= 35.3147
			summaries[i].Export *= 35.3147
			summaries[i].PersonalNeeds *= 35.3147
			summaries[i].ReinjectionIntoReservoir *= 35.3147
			summaries[i].GasSales *= 35.3147
		}
		// Calculate the difference
		summaries[i].Difference = summaries[i].TotalGasProduction + summaries[i].Import - summaries[i].Export - summaries[i].PersonalNeeds - summaries[i].ReinjectionIntoReservoir - summaries[i].GasSales
	}

	return summaries, nil
}
