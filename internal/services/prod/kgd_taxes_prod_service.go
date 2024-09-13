package services

import (
	"context"
	"fmt"
	"sync"
	"tender/internal/caching"
	models "tender/internal/models/prod" // Ensure this is correct
	repositories "tender/internal/repositories/prod"
	"time"
)

type KgdTaxesProdService struct {
	Repo  *repositories.KgdTaxesProdRepository
	Cache caching.Cache
}

func (s *KgdTaxesProdService) GetKgdTaxesProdSummary(ctx context.Context, year int, currency string) ([]models.KgdTaxesProdSummary, error) {
	cacheKey := fmt.Sprintf("kgd_taxes_prod_%d_%s", year, currency)

	// Проверка кэша
	if cachedData, found := s.Cache.Get(ctx, cacheKey); found {
		return cachedData.([]models.KgdTaxesProdSummary), nil
	}

	yearRanges := generateYearRanges(2015, year, 2)
	var wg sync.WaitGroup
	const maxGoroutines = 20
	sem := make(chan struct{}, maxGoroutines)

	results := make([]models.KgdTaxesProdSummary, 0)
	errors := make([]error, 0)

	// Использование горутин для параллельной обработки запросов
	for _, yr := range yearRanges {
		wg.Add(1)
		go func(startYear, endYear int) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			res, err := s.Repo.GetKgdTaxesProdSummaryForYearRange(ctx, startYear, endYear, currency)
			if err != nil {
				errors = append(errors, err)
				return
			}
			results = append(results, res...)
		}(yr.startYear, yr.endYear)
	}

	wg.Wait()

	if len(errors) > 0 {
		return nil, errors[0]
	}

	// Кэширование результата на 10 минут
	s.Cache.Set(ctx, cacheKey, results, 10*time.Minute)

	return results, nil
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

func (s *KgdTaxesProdService) GetSummaAllTaxes(ctx context.Context, year int, currency, reporttype, language string) (map[string]float64, map[string]float64, error) {
	totalSumSummary, totalSumSummaryy, err := s.Repo.GetSummaAllTaxes(ctx, year, currency, reporttype, language)
	if err != nil {
		return nil, nil, err
	}
	return totalSumSummary, totalSumSummaryy, nil
}
