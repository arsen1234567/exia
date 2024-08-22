package repositories

import (
	"context"
	"database/sql"
	// Ensure this is correct
)

type InvestPotentialMainRepository struct {
	Db *sql.DB
}

func (r *InvestPotentialMainRepository) GetInvestPotentialMain(ctx context.Context) (float64, error) {
	query := `
	SELECT COALESCE(SUM("Кратность запасов"), 0) AS total_reserve_multiple
	FROM dmart.invest_potential_main;
	`
	row := r.Db.QueryRowContext(ctx, query)

	var totalReserveMultiple float64
	if err := row.Scan(&totalReserveMultiple); err != nil {
		return 0, err
	}

	return totalReserveMultiple, nil
}

func (r *InvestPotentialMainRepository) GetSpecOpEx(ctx context.Context, currency, unit string) (float64, error) {
	query := `
	SELECT 
		CASE 
			WHEN SUM("OilProduction") = 0 THEN 0
			ELSE SUM("OperatingCosts") / SUM("OilProduction")
		END AS total_operational_expenses
	FROM 
		dmart.investment_review_forecast_steps
	WHERE 
		"Year" IN (2023, 2024, 2025, 2026, 2027, 2028)
		AND "scenario" = 'Forecast BBrent BCPI' AND
		"currency" = $1 AND
		"unit" = $2;
	`

	row := r.Db.QueryRowContext(ctx, query, currency, unit)

	var totalOperationalExpenses float64
	if err := row.Scan(&totalOperationalExpenses); err != nil {
		return 0, err
	}

	return totalOperationalExpenses, nil
}
