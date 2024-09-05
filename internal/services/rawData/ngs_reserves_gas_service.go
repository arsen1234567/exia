package services

import (
	"context"
	"fmt"
	"time"

	"tender/internal/caching"
	models "tender/internal/models/rawData" // Ensure this is correct
	repositories "tender/internal/repositories/rawData"
)

type NgsReservesGasService struct {
	Repo  *repositories.NgsReservesGasRepository
	Cache caching.Cache
}

func (s *NgsReservesGasService) GetRecoverableGasReservesSummary(ctx context.Context, year int) ([]models.RecoverableGasReservesSummary, error) {
	// Generate cache key based on the year
	cacheKey := fmt.Sprintf("recoverable_gas_reserves_%d", year)

	// Check if the result is in the cache
	if cachedData, found := s.Cache.Get(ctx, cacheKey); found {
		return cachedData.([]models.RecoverableGasReservesSummary), nil
	}

	// Fetch data if not in cache
	resultChan := make(chan []models.RecoverableGasReservesSummary)
	errorChan := make(chan error)

	partitions := []struct {
		startYear int
		endYear   int
	}{
		{1990, 2000},
		{2001, year},
	}

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

	var finalResults []models.RecoverableGasReservesSummary
	for i := 0; i < len(partitions); i++ {
		select {
		case res := <-resultChan:
			finalResults = append(finalResults, res...)
		case err := <-errorChan:
			return nil, err
		}
	}

	// Store the result in the cache
	s.Cache.Set(ctx, cacheKey, finalResults, 10*time.Minute)

	return finalResults, nil
}
