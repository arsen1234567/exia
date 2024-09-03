package services

import (
	"context"
	"sync"
	models "tender/internal/models/prod" // Ensure this is correct
	repositories "tender/internal/repositories/prod"
)

type KgdTaxesProdService struct {
	Repo *repositories.KgdTaxesProdRepository
}

func (s *KgdTaxesProdService) GetKgdTaxesProdSummary(ctx context.Context, year int, currency string) ([]models.KgdTaxesProdSummary, error) {
	// Define the year ranges for parallel processing
	yearRanges := generateYearRanges(2015, year, 2) // Split into 2-year ranges for example

	resultChan := make(chan []models.KgdTaxesProdSummary, len(yearRanges))
	errorChan := make(chan error, len(yearRanges))

	var wg sync.WaitGroup

	// Execute queries concurrently
	for _, yr := range yearRanges {
		wg.Add(1)
		go func(startYear, endYear int) {
			defer wg.Done()
			results, err := s.Repo.GetKgdTaxesProdSummaryForYearRange(ctx, startYear, endYear, currency)
			if err != nil {
				errorChan <- err
				return
			}
			resultChan <- results
		}(yr.startYear, yr.endYear)
	}

	// Wait for all Goroutines to finish
	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	// Collect results
	finalResults := make([]models.KgdTaxesProdSummary, 0)
	for res := range resultChan {
		finalResults = append(finalResults, res...)
	}

	// Handle errors
	if len(errorChan) > 0 {
		return nil, <-errorChan
	}

	return finalResults, nil
}

// generateYearRanges splits the year range into smaller chunks
func generateYearRanges(startYear, endYear, interval int) []struct{ startYear, endYear int } {
	var ranges []struct{ startYear, endYear int }
	for y := startYear; y <= endYear; y += interval {
		ey := y + interval - 1
		if ey > endYear {
			ey = endYear
		}
		ranges = append(ranges, struct{ startYear, endYear int }{y, ey})
	}
	return ranges
}

func (s *KgdTaxesProdService) GetSummaAllTaxes(ctx context.Context, year int, currency, reporttype string) (map[string]float64, error) {
	totalSumSummary, err := s.Repo.GetSummaAllTaxes(ctx, year, currency, reporttype)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}
