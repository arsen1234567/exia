package services

import (
	"context"
	"sync"
	models "tender/internal/models/rawData" // Ensure this is correct
	repositories "tender/internal/repositories/rawData"
)

type KgdTaxesService struct {
	Repo *repositories.KgdTaxesRepository
}

func (s *KgdTaxesService) GetKgdTaxesSummary(ctx context.Context, year int, currency string) ([]models.KgdTaxesSummary, error) {
	binGroups := [][]string{
		{"000340002165", "050840002757"},
		{"970740000392", "020440001144"},
		{"060640006784", "101140017122"},
		// Add more bin groups here if necessary
	}

	var wg sync.WaitGroup
	resultChan := make(chan []models.KgdTaxesSummary, len(binGroups))
	errorChan := make(chan error, len(binGroups))

	// Launch Goroutines to fetch data in parallel
	for _, bins := range binGroups {
		wg.Add(1)
		go func(bins []string) {
			defer wg.Done()
			summary, err := s.Repo.GetKgdTaxesSummaryByBin(ctx, year, bins...)
			if err != nil {
				errorChan <- err
				return
			}
			resultChan <- summary
		}(bins)
	}

	// Wait for all Goroutines to finish
	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	// Collect results and handle errors
	var finalResults []models.KgdTaxesSummary
	for res := range resultChan {
		finalResults = append(finalResults, res...)
	}

	// If there were errors, return the first one
	if len(errorChan) > 0 {
		return nil, <-errorChan
	}

	// If the currency is USD, divide each summary value by 456
	if currency == "USD" {
		for i := range finalResults {
			finalResults[i].TotalSum = finalResults[i].TotalSum / 456
		}
	}

	return finalResults, nil
}
