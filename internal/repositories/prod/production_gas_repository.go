package repositories

import (
	"context"
	"database/sql"
	models "tender/internal/models/prod"
)

type ProductionGasRepository struct {
	Db *sql.DB
}

func (r *ProductionGasRepository) GetGasProductionSummary(ctx context.Context, year int) ([]models.ProductionGasSummary, error) {
	query := `
	WITH YearlyProduction AS (
		SELECT 
			year,
			SUM(production_associated + production_natural) AS total_production
		FROM 
			prod.production_gas
		WHERE 
			year BETWEEN 2021 AND CAST($1 AS INTEGER)
		GROUP BY 
			year
	)
	SELECT 
		year,
		total_production
	FROM 
		YearlyProduction
	ORDER BY 
		year;
	`

	rows, err := r.Db.QueryContext(ctx, query, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.ProductionGasSummary
	for rows.Next() {
		var summary models.ProductionGasSummary
		if err := rows.Scan(&summary.Year, &summary.TotalProduction); err != nil {
			return nil, err
		}
		results = append(results, summary)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
