package repositories

import (
	"context"
	"database/sql"
	models "tender/internal/models/dmart" // Ensure this is correct
)

type InvestmentOilProductionRepository struct {
	Db *sql.DB
}

func (r *InvestmentOilProductionRepository) GetInvestmentOilProductionSummary(ctx context.Context, year int, unit string) ([]models.InvestmentOilProductionSummary, error) {
	query := `
		SELECT
			'Общий фонд РК' AS coverage_scope,
			unit,
			SUM(value) AS total_value
		FROM 
			dmart.investment_oil_production
		WHERE 
			year = $1
			AND unit = $2
		GROUP BY
			unit

		UNION ALL

		SELECT
			CASE
				WHEN abd_scope THEN 'Покрытие АБД'
				ELSE 'Вне периметра'
			END AS coverage_scope,
			unit,
			SUM(value) AS total_value
		FROM 
			dmart.investment_oil_production
		WHERE 
			year = $1
			AND unit = $2
		GROUP BY
			CASE
				WHEN abd_scope THEN 'Покрытие АБД'
				ELSE 'Вне периметра'
			END,
			unit;
	`

	rows, err := r.Db.QueryContext(ctx, query, year, unit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.InvestmentOilProductionSummary
	for rows.Next() {
		var summary models.InvestmentOilProductionSummary
		if err := rows.Scan(&summary.CoverageScope, &summary.Unit, &summary.TotalValue); err != nil {
			return nil, err
		}
		results = append(results, summary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
