package repositories

import (
	"context"
	"database/sql"
	models "tender/internal/models/rawData"

	"github.com/lib/pq"
)

type KgdTaxesRepository struct {
	Db *sql.DB
}

func (r *KgdTaxesRepository) GetKgdTaxesSummaryByBin(ctx context.Context, year int, bins ...string) ([]models.KgdTaxesSummary, error) {
	query := `
	SELECT 
		EXTRACT(YEAR FROM receipt_date) AS year,
		SUM(summa) AS total_sum
	FROM 
		raw_data.kgd_taxes
	WHERE 
		receipt_date < '2024-01-01'
		AND bin = ANY($2)
		AND pay_status = '2 - Разнесен'
		AND pay_type = 'Налог'
		AND EXTRACT(YEAR FROM receipt_date) <= CAST($1 AS INTEGER)
	GROUP BY 
		year
	ORDER BY 
		year;
	`

	rows, err := r.Db.QueryContext(ctx, query, year, pq.Array(bins))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.KgdTaxesSummary
	for rows.Next() {
		var summary models.KgdTaxesSummary
		if err := rows.Scan(&summary.Year, &summary.TotalSum); err != nil {
			return nil, err
		}
		results = append(results, summary)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
