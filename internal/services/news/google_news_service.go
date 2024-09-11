package services

import (
	"context"
	models "tender/internal/models/news"
	repositories "tender/internal/repositories/news"
	"time"
)

type GoogleNewsService struct {
	Repo *repositories.GoogleNewsRepository
}

// GetTotalKpnByUsername retrieves the total KPN for a given username using the repository
func (s *GoogleNewsService) GetGoogleNews(ctx context.Context, period string) ([]models.GoogleNewsSummary, error) {
	return s.Repo.GetCompanyCount(ctx, period)
}

func (s *GoogleNewsService) GetMediaCount(ctx context.Context, period string) ([]models.GoogleNewsSummary, error) {
	return s.Repo.GetMediaCount(ctx, period)
}

func (s *GoogleNewsService) GetMediaNews(ctx context.Context, period string) ([]models.GoogleNews, error) {
	return s.Repo.GetMediaNews(ctx, period)
}

func (s *GoogleNewsService) GetTonal(ctx context.Context, period string) (map[string]int, error) {
	return s.Repo.GetTonal(ctx, period)
}

func (s *GoogleNewsService) GetSentimentCountsByDay(ctx context.Context, period string) (map[time.Time]map[string]int, error) {
	return s.Repo.GetSentimentCountsByDay(ctx, period)
}

func (s *GoogleNewsService) GetSentimentMap(ctx context.Context, period, sentiment string) (map[time.Time]int, error) {
	return s.Repo.GetSentimentMap(ctx, period, sentiment)
}

func (s *GoogleNewsService) GetTopCompanyCount(ctx context.Context, period string) ([]models.GoogleNewsSummary, error) {
	return s.Repo.GetTopCompanyCount(ctx, period)
}

func (s *GoogleNewsService) GetTopCompanyCountDict(ctx context.Context, period string) ([]models.GoogleNewsSummary, error) {
	return s.Repo.GetTopCompanyCount(ctx, period)
}

func (s *GoogleNewsService) GetSourceCount(ctx context.Context, period string) ([]models.GoogleNewsSummary, error) {
	return s.Repo.GetSourceCount(ctx, period)
}

func (s *GoogleNewsService) GetNewsByPeriod(period string) ([]models.NewsSummary, error) {
	startDate, endDate := CalculatePeriod(period)
	return s.Repo.GetNewsByPeriod(startDate, endDate)
}

// Вспомогательная функция для расчета интервала
func CalculatePeriod(period string) (time.Time, time.Time) {
	now := time.Now()
	var startDate, endDate time.Time

	switch period {
	case "last_month":
		startDate = now.AddDate(0, -1, 0)
		endDate = now
	case "last_week":
		startDate = now.AddDate(0, 0, -7)
		endDate = now
	case "last_year":
		startDate = now.AddDate(-1, 0, 0)
		endDate = now
	default:
		startDate = now.AddDate(0, 0, -30) // Default to last 30 days
		endDate = now
	}

	return startDate, endDate
}
