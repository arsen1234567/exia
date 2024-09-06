package services

import (
	"context"
	"fmt"
	"sync"
	"tender/internal/caching"
	models "tender/internal/models/rawData" // Убедитесь, что путь правильный
	repositories "tender/internal/repositories/rawData"
	"time" // Для использования времени истечения кэша
)

type KgdTaxesService struct {
	Repo  *repositories.KgdTaxesRepository
	Cache caching.Cache
}

func (s *KgdTaxesService) GetKgdTaxesSummary(ctx context.Context, year int, currency string) ([]models.KgdTaxesSummary, error) {

	cacheKey := fmt.Sprintf("kgd_taxes_%d_%s", year, currency)

	if cachedData, found := s.Cache.Get(ctx, cacheKey); found {
		return cachedData.([]models.KgdTaxesSummary), nil
	}

	binGroups := [][]string{
		{"000340002165", "050840002757"},
		{"970740000392", "020440001144"},
		{"060640006784", "101140017122"},
	}

	const maxGoroutines = 4
	sem := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup
	resultChan := make(chan []models.KgdTaxesSummary, len(binGroups))
	errorChan := make(chan error, len(binGroups))

	for _, bins := range binGroups {
		wg.Add(1)
		sem <- struct{}{}
		go func(bins []string) {
			defer wg.Done()
			defer func() { <-sem }()

			summary, err := s.Repo.GetKgdTaxesSummaryByBin(ctx, year, bins...)
			if err != nil {
				errorChan <- err
				return
			}
			resultChan <- summary
		}(bins)
	}

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	var finalResults []models.KgdTaxesSummary
	for res := range resultChan {
		finalResults = append(finalResults, res...)
	}

	// Если есть ошибки, возвращаем первую
	if len(errorChan) > 0 {
		return nil, <-errorChan
	}

	// Конвертация валюты в USD, если указана USD
	if currency == "USD" {
		for i := range finalResults {
			finalResults[i].TotalSum /= 456
		}
	}

	s.Cache.Set(ctx, cacheKey, finalResults, 10*time.Minute)

	return finalResults, nil
}
