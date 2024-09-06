package services

import (
	"context"
	"fmt"
	"sync"
	"tender/internal/caching"
	models "tender/internal/models/dmart" // Ensure this is correct
	repositories_dmart "tender/internal/repositories/dmart"
	"time"
)

type InvestmentOilProductionService struct {
	Repo  *repositories_dmart.InvestmentOilProductionRepository
	Cache caching.Cache
}

// GetInvestmentOilProductionSummary retrieves a summary of investment oil production for a specific year.
func (s *InvestmentOilProductionService) GetInvestmentOilProductionSummary(ctx context.Context, year int, unit string) ([]models.InvestmentOilProductionSummary, error) {

	cacheKey := fmt.Sprintf("investment_oil_production_%d_%s", year, unit)

	if cachedData, found := s.Cache.Get(ctx, cacheKey); found {
		return cachedData.([]models.InvestmentOilProductionSummary), nil
	}

	var wg sync.WaitGroup
	resultChan := make(chan []models.InvestmentOilProductionSummary, 1)
	errorChan := make(chan error, 1)

	wg.Add(1)

	go func() {

		defer wg.Done()
		summary, err := s.Repo.GetInvestmentOilProductionSummary(ctx, year, unit)

		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- summary
	}()

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	var finalResults []models.InvestmentOilProductionSummary
	for res := range resultChan {
		finalResults = append(finalResults, res...)
	}

	if len(errorChan) > 0 {
		return nil, <-errorChan
	}

	s.Cache.Set(ctx, cacheKey, finalResults, 10*time.Minute)

	return finalResults, nil

}
