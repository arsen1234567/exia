package repositories

import (
	"context"
	"database/sql"
	"fmt"
	models "tender/internal/models/dmart"
)

type InvestmentReviewForecastStepsRepository struct {
	Db *sql.DB
}

func (r *InvestmentReviewForecastStepsRepository) GetInvestmentReviewForecastSteps(ctx context.Context) (float64, error) {
	query := `
	SELECT 
	    SUM("GovShare") AS total_gov_share
	FROM 
	    dmart.investment_review_forecast_steps
	WHERE 
	    scenario IN ('Forecast BBrent BCPI')
	    AND "Year" IN (2023, 2024, 2025, 2026, 2027, 2028);
	`

	var totalGovShare sql.NullFloat64
	err := r.Db.QueryRowContext(ctx, query).Scan(&totalGovShare)
	if err != nil {
		return 0, err
	}

	if !totalGovShare.Valid {
		return 0, fmt.Errorf("totalGovShare is NULL")
	}

	return totalGovShare.Float64, nil
}

func (r *InvestmentReviewForecastStepsRepository) GetEbitdaToGrossRevenueRatio(ctx context.Context) (float64, error) {
	query := `
	SELECT 
	    SUM("EBITDA") / SUM("GrossRevenue") AS ebitda_to_gross_revenue_ratio
	FROM 
	    dmart.investment_review_forecast_steps
	WHERE 
	    scenario IN ('Forecast BBrent BCPI')
	    AND "Year" IN (2023, 2024, 2025, 2026, 2027, 2028);
	`

	var ebitdaToGrossRevenueRatio sql.NullFloat64
	err := r.Db.QueryRowContext(ctx, query).Scan(&ebitdaToGrossRevenueRatio)
	if err != nil {
		return 0, err
	}

	if !ebitdaToGrossRevenueRatio.Valid {
		return 0, fmt.Errorf("ebitdaToGrossRevenueRatio is NULL")
	}

	return ebitdaToGrossRevenueRatio.Float64, nil
}

func (r *InvestmentReviewForecastStepsRepository) GetCompaniesForecastSteps(ctx context.Context, currency, unit string) ([]models.InvestmentReviewForecastStepsSummary, error) {
	query := `
	SELECT 
	    "name_abbr", 
	    "Year",
	    SUM("OilProduction") AS OilProduction
	FROM 
	    dmart.investment_review_forecast_steps
	WHERE 
	    scenario IN ('Forecast BBrent BCPI') AND
		currency = $1 AND
		unit = $2
	GROUP BY 
	    "name_abbr", "Year"
	ORDER BY
	    "Year", "name_abbr"
	`

	rows, err := r.Db.QueryContext(ctx, query, currency, unit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.InvestmentReviewForecastStepsSummary
	for rows.Next() {
		var op models.InvestmentReviewForecastStepsSummary
		if err := rows.Scan(&op.NameAbbr, &op.Year, &op.OilProduction); err != nil {
			return nil, err
		}
		results = append(results, op)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
