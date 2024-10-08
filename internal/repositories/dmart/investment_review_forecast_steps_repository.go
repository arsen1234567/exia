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

func (r *InvestmentReviewForecastStepsRepository) GetInvestmentReviewForecastSteps(ctx context.Context, currency string) (float64, error) {
	query := `
	SELECT 
	    SUM("GovShare") AS total_gov_share
	FROM 
	    dmart.investment_review_forecast_steps
	WHERE 
	    scenario IN ('Forecast BBrent BCPI')
	    AND "Year" IN (2023, 2024, 2025, 2026, 2027, 2028)
		AND currency = $1;
	`

	var totalGovShare sql.NullFloat64
	err := r.Db.QueryRowContext(ctx, query, currency).Scan(&totalGovShare)
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

func (r *InvestmentReviewForecastStepsRepository) GetCompaniesForecastSteps(ctx context.Context, unit, language string) ([]models.InvestmentReviewForecastStepsSummary, error) {

	query := fmt.Sprintf(`
		SELECT 
			%s AS NameAbbr, 
			"Year",
			SUM("OilProduction") AS OilProduction
		FROM 
			dmart.investment_review_forecast_steps
		WHERE 
			scenario IN ('Forecast BBrent BCPI') AND
			unit = $1
		GROUP BY 
			%s, "Year"
		ORDER BY
			"Year", %s
	`, language, language, language)

	// Выполняем запрос
	rows, err := r.Db.QueryContext(ctx, query, unit)
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

func (r *InvestmentReviewForecastStepsRepository) GetRevenueInvestmentReviewForecastSteps(ctx context.Context, unit, currency string) ([]models.RevenueByYear, error) {
	query := `
	SELECT 
	    "name_abbr",
	    "Year",
	    SUM("GrossRevenue") AS total_gross_revenue
	FROM 
	    dmart.investment_review_forecast_steps
	WHERE 
	    scenario IN ('Forecast BBrent BCPI')
	    AND unit = $1
	    AND currency = $2
	GROUP BY 
	    "name_abbr", "Year"
	ORDER BY 
	    "name_abbr", "Year";
	`

	rows, err := r.Db.QueryContext(ctx, query, unit, currency)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.RevenueByYear
	for rows.Next() {
		var rByCompanyAndYear models.RevenueByYear
		if err := rows.Scan(&rByCompanyAndYear.CompanyName, &rByCompanyAndYear.Year, &rByCompanyAndYear.TotalGrossRevenue); err != nil {
			return nil, err
		}
		results = append(results, rByCompanyAndYear)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *InvestmentReviewForecastStepsRepository) GetCapExInvestmentReviewForecastSteps(ctx context.Context, unit, currency string) ([]models.InvestmentReviewForecastData, error) {
	query := `
	SELECT 
	    "name_abbr",
	    "Year",
	    SUM("CapEx") AS cap_ex
	FROM 
	    dmart.investment_review_forecast_steps
	WHERE 
	    scenario IN ('Forecast BBrent BCPI')
	    AND unit = $1
	    AND currency = $2
	GROUP BY 
	    "name_abbr", "Year"
	ORDER BY 
	    "name_abbr", "Year";
	`

	rows, err := r.Db.QueryContext(ctx, query, unit, currency)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.InvestmentReviewForecastData
	for rows.Next() {
		var rByCompanyAndYear models.InvestmentReviewForecastData
		if err := rows.Scan(&rByCompanyAndYear.CompanyName, &rByCompanyAndYear.Year, &rByCompanyAndYear.Data); err != nil {
			return nil, err
		}
		results = append(results, rByCompanyAndYear)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *InvestmentReviewForecastStepsRepository) GetATFCFInvestmentReviewForecastSteps(ctx context.Context, unit, currency string) ([]models.InvestmentReviewForecastData, error) {
	query := `
	SELECT 
	    "name_abbr",
	    "Year",
	    SUM("ATFCF") AS cap_ex
	FROM 
	    dmart.investment_review_forecast_steps
	WHERE 
	    scenario IN ('Forecast BBrent BCPI')
	    AND unit = $1
	    AND currency = $2
	GROUP BY 
	    "name_abbr", "Year"
	ORDER BY 
	    "name_abbr", "Year";
	`

	rows, err := r.Db.QueryContext(ctx, query, unit, currency)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.InvestmentReviewForecastData
	for rows.Next() {
		var rByCompanyAndYear models.InvestmentReviewForecastData
		if err := rows.Scan(&rByCompanyAndYear.CompanyName, &rByCompanyAndYear.Year, &rByCompanyAndYear.Data); err != nil {
			return nil, err
		}
		results = append(results, rByCompanyAndYear)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *InvestmentReviewForecastStepsRepository) GetOpExInvestmentReviewForecastSteps(ctx context.Context, unit, currency string) ([]models.InvestmentReviewForecastData, error) {
	query := `
	SELECT 
	    "name_abbr",
	    "Year",
	    SUM("OperatingCosts") AS op_ex
	FROM 
	    dmart.investment_review_forecast_steps
	WHERE 
	    scenario IN ('Forecast BBrent BCPI')
	    AND unit = $1
	    AND currency = $2
	GROUP BY 
	    "name_abbr", "Year"
	ORDER BY 
	    "name_abbr", "Year";
	`

	rows, err := r.Db.QueryContext(ctx, query, unit, currency)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.InvestmentReviewForecastData
	for rows.Next() {
		var rByCompanyAndYear models.InvestmentReviewForecastData
		if err := rows.Scan(&rByCompanyAndYear.CompanyName, &rByCompanyAndYear.Year, &rByCompanyAndYear.Data); err != nil {
			return nil, err
		}
		results = append(results, rByCompanyAndYear)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *InvestmentReviewForecastStepsRepository) GetGovShareInvestmentReviewForecastSteps(ctx context.Context, unit, currency string) ([]models.InvestmentReviewForecastData, error) {
	query := `
	SELECT 
	    "name_abbr",
	    "Year",
	    SUM("GovShare") AS gov_share
	FROM 
	    dmart.investment_review_forecast_steps
	WHERE 
	    scenario IN ('Forecast BBrent BCPI')
	    AND unit = $1
	    AND currency = $2
	GROUP BY 
	    "name_abbr", "Year"
	ORDER BY 
	    "name_abbr", "Year";
	`

	rows, err := r.Db.QueryContext(ctx, query, unit, currency)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.InvestmentReviewForecastData
	for rows.Next() {
		var rByCompanyAndYear models.InvestmentReviewForecastData
		if err := rows.Scan(&rByCompanyAndYear.CompanyName, &rByCompanyAndYear.Year, &rByCompanyAndYear.Data); err != nil {
			return nil, err
		}
		results = append(results, rByCompanyAndYear)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
