package services

import (
	"context"
	models "tender/internal/models/rawData" // Ensure this is correct
	repositories "tender/internal/repositories/rawData"
)

type NgsReservesGasService struct {
	Repo *repositories.NgsReservesGasRepository
}

func (s *NgsReservesGasService) GetRecoverableGasReservesSummary(ctx context.Context, year int) ([]models.RecoverableGasReservesSummary, error) {
	// Define channels for results and errors
	resultChan := make(chan []models.RecoverableGasReservesSummary)
	errorChan := make(chan error)

	// Define time ranges or other partitions (e.g., split the years)
	// For simplicity, here we just use two partitions: before and after 2000
	partitions := []struct {
		startYear int
		endYear   int
	}{
		{1990, 2000},
		{2001, year},
	}

	// Launch Goroutines to fetch data in parallel
	for _, partition := range partitions {
		go func(startYear, endYear int) {
			summary, err := s.Repo.GetRecoverableGasReservesByYearRange(ctx, startYear, endYear)
			if err != nil {
				errorChan <- err
				return
			}
			resultChan <- summary
		}(partition.startYear, partition.endYear)
	}

	// Collect results
	var finalResults []models.RecoverableGasReservesSummary
	for i := 0; i < len(partitions); i++ {
		select {
		case res := <-resultChan:
			finalResults = append(finalResults, res...)
		case err := <-errorChan:
			return nil, err
		}
	}

	return finalResults, nil
}
