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

type InvestmentReservesService struct {
	Repo  *repositories_dmart.InvestmentReservesRepository
	Cache caching.Cache
}

func (s *InvestmentReservesService) GetInvestmentReservesSummary(ctx context.Context, year int, unit string) ([]models.InvestmentReservesSummary, error) {
	cacheKey := fmt.Sprintf("investment_reserves_summary_%d_%s", year, unit)

	// Проверка кэша
	if cachedData, found := s.Cache.Get(ctx, cacheKey); found {
		return cachedData.([]models.InvestmentReservesSummary), nil
	}

	// Подготовка к параллельному запросу
	var wg sync.WaitGroup
	const maxGoroutines = 4
	sem := make(chan struct{}, maxGoroutines)

	results := make([]models.InvestmentReservesSummary, 0)
	errors := make([]error, 0)

	// Запуск горутин для параллельного получения данных
	wg.Add(1)
	go func() {
		defer wg.Done()
		sem <- struct{}{}
		defer func() { <-sem }()

		summary, err := s.Repo.GetInvestmentReservesSummary(ctx, year, unit)
		if err != nil {
			errors = append(errors, err)
			return
		}
		results = append(results, summary...)
	}()

	// Ожидание завершения всех горутин
	wg.Wait()

	if len(errors) > 0 {
		return nil, errors[0]
	}

	// Кэширование результата на 10 минут
	s.Cache.Set(ctx, cacheKey, results, 10*time.Minute)

	return results, nil
}
