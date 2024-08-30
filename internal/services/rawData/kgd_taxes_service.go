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
	// Fetch the summary from the repository
	summary, err := s.Repo.GetKgdTaxesSummary(ctx, year)
	if err != nil {
		return nil, err
	}

	// If the currency is USD, divide each summary value by 456
	if currency == "USD" {
		for i := range summary {
			summary[i].TotalSum = summary[i].TotalSum / 456
		}
	}

	return summary, nil
}
