package services

import (
	"context" // Ensure this is correct
	"fmt"
	"sync"
	"tender/internal/caching"
	repositories "tender/internal/repositories/prod"
	"time"
)

type TaxBurdenService struct {
	Repo  *repositories.TaxBurdenRepository
	Cache caching.Cache
}

func (s *TaxBurdenService) GetTaxBurden(ctx context.Context, year int, currency string) (map[string]float64, error) {
	cacheKey := fmt.Sprintf("tax_burden_%d_%s", year, currency)

	if cachedData, found := s.Cache.Get(ctx, cacheKey); found {
		return cachedData.(map[string]float64), nil
	}

	yearRanges := []int{year}

	resultChan := make(chan map[string]float64, len(yearRanges))
	errorChan := make(chan error, len(yearRanges))

	var wg sync.WaitGroup

	for _, yr := range yearRanges {
		wg.Add(1)
		go func(year int) {
			defer wg.Done()
			results, err := s.Repo.GetTaxBurden(ctx, year, currency)
			if err != nil {
				errorChan <- err
				return
			}
			resultChan <- results
		}(yr)
	}

	// Wait for all Goroutines to finish
	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	// Collect results and handle errors
	finalResults := make(map[string]float64)
	var finalErr error

	for res := range resultChan {
		for key, value := range res {
			finalResults[key] = value
		}
	}

	if len(errorChan) > 0 {
		finalErr = <-errorChan
	}

	// Cache the result if no errors
	if finalErr == nil {
		s.Cache.Set(ctx, cacheKey, finalResults, 10*time.Minute)
	}

	return finalResults, finalErr
}
