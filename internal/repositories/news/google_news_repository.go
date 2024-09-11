package repositories

import (
	"context"
	"database/sql"
	"fmt"
	models "tender/internal/models/news"
	"time"
)

type GoogleNewsRepository struct {
	Db *sql.DB
}

func (r *GoogleNewsRepository) GetCompanyCount(ctx context.Context, period string) ([]models.GoogleNewsSummary, error) {
	var startDate, endDate time.Time
	now := time.Now()

	switch period {
	case "last_week":
		startDate = now.AddDate(0, 0, -7).Truncate(24 * time.Hour)
		endDate = now.Truncate(24 * time.Hour)
	case "last_month":
		startDate = now.AddDate(0, -1, 0).Truncate(24 * time.Hour)
		endDate = now.Truncate(24 * time.Hour)
	case "last_quarter":
		startDate = now.AddDate(0, -3, 0).Truncate(24 * time.Hour)
		endDate = now.Truncate(24 * time.Hour)
	case "last_year":
		startDate = now.AddDate(-1, 0, 0).Truncate(24 * time.Hour)
		endDate = now.Truncate(24 * time.Hour)
	default:
		return nil, fmt.Errorf("invalid period: %s", period)
	}

	query := `
    SELECT 
        keyword AS company_name,
        COUNT(*) AS company_count
    FROM 
        news.google_news
    WHERE 
        date_news >= $1
        AND date_news < $2
    GROUP BY 
        keyword
    ORDER BY 
        keyword ASC;
    `

	rows, err := r.Db.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.GoogleNewsSummary
	for rows.Next() {
		var summary models.GoogleNewsSummary
		if err := rows.Scan(&summary.CompanyName, &summary.CompanyCount); err != nil {
			return nil, err
		}
		summary.StartDate = startDate
		summary.EndDate = endDate
		results = append(results, summary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *GoogleNewsRepository) GetMediaCount(ctx context.Context, period string) ([]models.GoogleNewsSummary, error) {
	var startDate, endDate time.Time
	now := time.Now()

	switch period {
	case "last_week":
		startDate = now.AddDate(0, 0, -7)
		endDate = now
	case "last_month":
		startDate = now.AddDate(0, -1, 0)
		endDate = now
	case "last_quarter":
		startDate = now.AddDate(0, -3, 0)
		endDate = now
	case "last_year":
		startDate = now.AddDate(-1, 0, 0)
		endDate = now
	default:
		return nil, fmt.Errorf("invalid period: %s", period)
	}

	query := `
    SELECT 
        source AS company_name,
        COUNT(*) AS company_count
    FROM 
        news.google_news
    WHERE 
        date_news >= $1
        AND date_news <= $2
    GROUP BY 
        source
    ORDER BY 
	source ASC;
    `

	rows, err := r.Db.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.GoogleNewsSummary
	for rows.Next() {
		var summary models.GoogleNewsSummary
		if err := rows.Scan(&summary.CompanyName, &summary.CompanyCount); err != nil {
			return nil, err
		}
		summary.StartDate = startDate
		summary.EndDate = endDate
		results = append(results, summary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *GoogleNewsRepository) GetMediaNews(ctx context.Context, period string) ([]models.GoogleNews, error) {
	var startDate, endDate time.Time
	now := time.Now()

	switch period {
	case "last_week":
		startDate = now.AddDate(0, 0, -7)
		endDate = now
	case "last_month":
		startDate = now.AddDate(0, -1, 0)
		endDate = now
	case "last_quarter":
		startDate = now.AddDate(0, -3, 0)
		endDate = now
	case "last_year":
		startDate = now.AddDate(-1, 0, 0)
		endDate = now
	default:
		return nil, fmt.Errorf("invalid period: %s", period)
	}

	query := `
    SELECT 
        date_news AS data_news,
        source AS company_name,
        title AS title,
        dict_sent_wrd AS dict_sent_wrd
    FROM 
        news.google_news
    WHERE 
        date_news >= $1
        AND date_news <= $2
    ORDER BY 
        source ASC;
    `

	rows, err := r.Db.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.GoogleNews
	for rows.Next() {
		var news models.GoogleNews
		if err := rows.Scan(&news.Date_news, &news.Source, &news.Title, &news.Dict_sent_wrd); err != nil {
			return nil, err
		}
		news.StartDate = startDate
		news.EndDate = endDate
		results = append(results, news)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *GoogleNewsRepository) GetTonal(ctx context.Context, period string) (map[string]int, error) {
	var startDate, endDate time.Time
	now := time.Now()

	switch period {
	case "last_week":
		startDate = now.AddDate(0, 0, -7)
		endDate = now
	case "last_month":
		startDate = now.AddDate(0, -1, 0)
		endDate = now
	case "last_quarter":
		startDate = now.AddDate(0, -3, 0)
		endDate = now
	case "last_year":
		startDate = now.AddDate(-1, 0, 0)
		endDate = now
	default:
		return nil, fmt.Errorf("invalid period: %s", period)
	}

	query := `
    SELECT 
        SUM(CASE WHEN dict_sent_wrd = 'позитивная' THEN 1 ELSE 0 END) AS positive_count,
        SUM(CASE WHEN dict_sent_wrd = 'негативная' THEN 1 ELSE 0 END) AS negative_count,
        SUM(CASE WHEN dict_sent_wrd = 'нейтральная' THEN 1 ELSE 0 END) AS neutral_count
    FROM 
        news.google_news
    WHERE 
        date_news >= $1
        AND date_news <= $2;
    `

	row := r.Db.QueryRowContext(ctx, query, startDate, endDate)

	var positiveCount, negativeCount, neutralCount int
	if err := row.Scan(&positiveCount, &negativeCount, &neutralCount); err != nil {
		return nil, err
	}

	// Create a map to store the counts for each sentiment
	sentimentCounts := map[string]int{
		"positive": positiveCount,
		"negative": negativeCount,
		"neutral":  neutralCount,
	}

	return sentimentCounts, nil
}

func (r *GoogleNewsRepository) GetSentimentCountsByDay(ctx context.Context, period string) (map[time.Time]map[string]int, error) {
	now := time.Now()
	var startDate, endDate time.Time

	switch period {
	case "last_week":
		startDate = now.AddDate(0, 0, -7)
		endDate = now
	case "last_month":
		startDate = now.AddDate(0, -1, 0)
		endDate = now
	case "last_quarter":
		startDate = now.AddDate(0, -3, 0)
		endDate = now
	case "last_year":
		startDate = now.AddDate(-1, 0, 0)
		endDate = now
	default:
		return nil, fmt.Errorf("invalid period: %s", period)
	}

	// // Adjusting to the exact start and end dates
	// if period == "last_month" {
	// 	startDate = time.Date(startDate.Year(), startDate.Month(), 1, 0, 0, time.UTC, time.UTC, time.UTC)
	// 	endDate = time.Date(endDate.Year(), endDate.Month(), 1, 0, 0, time.UTC, time.UTC, time.UTC).AddDate(0, 1, -1)
	// } else if period == "last_week" {
	// 	startDate = startDate.Truncate(24 * time.Hour)
	// 	endDate = endDate.Truncate(24 * time.Hour)
	// }

	query := `
    SELECT
        date_news AS date,
        dict_sent_wrd AS sentiment,
        COUNT(*) AS count
    FROM
        news.google_news
    WHERE
        date_news >= $1
        AND date_news <= $2
    GROUP BY
        date_news, dict_sent_wrd
    ORDER BY
        date_news ASC, dict_sent_wrd;
    `

	rows, err := r.Db.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make(map[time.Time]map[string]int)
	for rows.Next() {
		var date time.Time
		var sentiment string
		var count int

		if err := rows.Scan(&date, &sentiment, &count); err != nil {
			return nil, err
		}

		if _, exists := results[date]; !exists {
			results[date] = make(map[string]int)
		}
		results[date][sentiment] = count
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *GoogleNewsRepository) GetSentimentMap(ctx context.Context, period string, sentiment string) (map[time.Time]int, error) {
	now := time.Now()
	var startDate, endDate time.Time

	switch period {
	case "last_week":
		startDate = now.AddDate(0, 0, -7)
		endDate = now
	case "last_month":
		startDate = now.AddDate(0, -1, 0)
		endDate = now
	case "last_quarter":
		startDate = now.AddDate(0, -3, 0)
		endDate = now
	case "last_year":
		startDate = now.AddDate(-1, 0, 0)
		endDate = now
	default:
		return nil, fmt.Errorf("invalid period: %s", period)
	}

	// // Adjusting to the exact start and end dates for specific periods
	// if period == "last_month" {
	// 	startDate = time.Date(startDate.Year(), startDate.Month(), 1, 0, 0, time.UTC, time.UTC, time.UTC)
	// 	endDate = time.Date(endDate.Year(), endDate.Month(), 1, 0, 0, time.UTC, time.UTC, time.UTC).AddDate(0, 1, -1)
	// } else if period == "last_week" {
	// 	startDate = startDate.Truncate(24 * time.Hour)
	// 	endDate = endDate.Truncate(24 * time.Hour)
	// }

	query := `
    SELECT
        date_news AS date,
        COUNT(*) AS count
    FROM
        news.google_news
    WHERE
        date_news >= $1
        AND date_news <= $2
        AND dict_sent_wrd = $3
    GROUP BY
        date_news
    ORDER BY
        date_news ASC;
    `

	rows, err := r.Db.QueryContext(ctx, query, startDate, endDate, sentiment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make(map[time.Time]int)
	for rows.Next() {
		var date time.Time
		var count int

		if err := rows.Scan(&date, &count); err != nil {
			return nil, err
		}

		results[date] = count
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *GoogleNewsRepository) GetTopCompanyCount(ctx context.Context, period string) ([]models.GoogleNewsSummary, error) {
	var startDate, endDate time.Time
	now := time.Now()

	switch period {
	case "last_week":
		startDate = now.AddDate(0, 0, -7)
		endDate = now
	case "last_month":
		startDate = now.AddDate(0, -1, 0)
		endDate = now
	case "last_quarter":
		startDate = now.AddDate(0, -3, 0)
		endDate = now
	case "last_year":
		startDate = now.AddDate(-1, 0, 0)
		endDate = now
	default:
		return nil, fmt.Errorf("invalid period: %s", period)
	}

	query := `
    SELECT 
        keyword AS company_name,
        COUNT(*) AS company_count
    FROM 
        news.google_news
    WHERE 
        date_news >= $1
        AND date_news <= $2
    GROUP BY 
        keyword
    ORDER BY 
        company_count DESC;
    `

	rows, err := r.Db.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.GoogleNewsSummary
	for rows.Next() {
		var summary models.GoogleNewsSummary
		if err := rows.Scan(&summary.CompanyName, &summary.CompanyCount); err != nil {
			return nil, err
		}
		summary.StartDate = startDate
		summary.EndDate = endDate
		results = append(results, summary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *GoogleNewsRepository) GetTopCompanyCountDict(ctx context.Context, period string) ([]models.GoogleNewsSummary, error) {
	var startDate, endDate time.Time
	now := time.Now()

	switch period {
	case "last_week":
		startDate = now.AddDate(0, 0, -7)
		endDate = now
	case "last_month":
		startDate = now.AddDate(0, -1, 0)
		endDate = now
	case "last_quarter":
		startDate = now.AddDate(0, -3, 0)
		endDate = now
	case "last_year":
		startDate = now.AddDate(-1, 0, 0)
		endDate = now
	default:
		return nil, fmt.Errorf("invalid period: %s", period)
	}

	query := `
    SELECT 
        keyword AS company_name,
        SUM(CASE WHEN dict_sent_wrd = 'позитивная' THEN 1 ELSE 0 END) AS positive_count,
        SUM(CASE WHEN dict_sent_wrd = 'негативная' THEN 1 ELSE 0 END) AS negative_count,
        SUM(CASE WHEN dict_sent_wrd = 'нейтральная' THEN 1 ELSE 0 END) AS neutral_count
    FROM 
        news.google_news
    WHERE 
        date_news >= $1
        AND date_news <= $2
    GROUP BY 
        keyword
    ORDER BY 
        company_name ASC;
    `

	rows, err := r.Db.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.GoogleNewsSummary
	for rows.Next() {
		var summary models.GoogleNewsSummary
		if err := rows.Scan(&summary.CompanyName, &summary.PositiveCount, &summary.NegativeCount, &summary.NeutralCount); err != nil {
			return nil, err
		}
		summary.StartDate = startDate
		summary.EndDate = endDate
		results = append(results, summary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *GoogleNewsRepository) GetSourceCount(ctx context.Context, period string) ([]models.GoogleNewsSummary, error) {
	var startDate, endDate time.Time
	now := time.Now()

	switch period {
	case "last_week":
		startDate = now.AddDate(0, 0, -7)
		endDate = now
	case "last_month":
		startDate = now.AddDate(0, -1, 0)
		endDate = now
	case "last_quarter":
		startDate = now.AddDate(0, -3, 0)
		endDate = now
	case "last_year":
		startDate = now.AddDate(-1, 0, 0)
		endDate = now
	default:
		return nil, fmt.Errorf("invalid period: %s", period)
	}

	query := `
    SELECT 
        keyword AS company_name,
        source AS news_source,
        COUNT(*) AS mention_count
    FROM 
        news.google_news
    WHERE 
        date_news >= $1
        AND date_news <= $2
    GROUP BY 
        keyword, source
    ORDER BY 
        company_name, news_source;
    `

	rows, err := r.Db.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.GoogleNewsSummary
	for rows.Next() {
		var summary models.GoogleNewsSummary
		if err := rows.Scan(&summary.CompanyName, &summary.NewsSource, &summary.CompanyCount); err != nil {
			return nil, err
		}
		summary.StartDate = startDate
		summary.EndDate = endDate
		results = append(results, summary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *GoogleNewsRepository) GetNewsByPeriod(startDate, endDate time.Time) ([]models.NewsSummary, error) {
	query := `SELECT keyword, snippet, source, dict_sent_wrd FROM news.google_news WHERE date_news BETWEEN $1 AND $2`

	rows, err := r.Db.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var news []models.NewsSummary
	for rows.Next() {
		var gn models.NewsSummary
		if err := rows.Scan(&gn.Keyword, &gn.Snippet, &gn.Source, &gn.Sentiment); err != nil {
			return nil, err
		}
		news = append(news, gn)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return news, nil
}
