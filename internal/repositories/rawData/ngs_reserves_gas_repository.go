package repositories

import (
	"context"
	"database/sql"
	"fmt"
	models "tender/internal/models/rawData"
)

type NgsReservesGasRepository struct {
	Db *sql.DB
}

// GetRecoverableGasReserves retrieves the total recoverable gas reserves for each year
func (r *NgsReservesGasRepository) GetRecoverableGasReserves(ctx context.Context, startYear, endYear int) ([]models.RecoverableGasReservesSummary, error) {
	query := `
	SELECT 
		EXTRACT(YEAR FROM n."Год") AS year,
		SUM(n."Балансовые запасы на конец(А+В+С1)") + SUM(n."Бал. запасы утв. ГКЗ (С2)") AS total_reserves
	FROM 
		raw_data.ngs_reserves_gas n
	RIGHT JOIN 
		dict.companies_1 c ON c.bin::text = n.bin
	WHERE 
		EXTRACT(YEAR FROM n."Год") BETWEEN $1 AND $2
		AND n."Тип" = 'Извлекаемые'
	GROUP BY 
		EXTRACT(YEAR FROM n."Год")
	ORDER BY 
		year;
	`

	rows, err := r.Db.QueryContext(ctx, query, startYear, endYear)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var results []models.RecoverableGasReservesSummary
	for rows.Next() {
		var summary models.RecoverableGasReservesSummary
		if err := rows.Scan(&summary.Year, &summary.TotalReserves); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, summary)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return results, nil
}
