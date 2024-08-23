package repositories

import (
	"context"
	"database/sql"
)

type TaxBurdenRepository struct {
	Db *sql.DB
}

func (r *TaxBurdenRepository) GetTaxBurden(ctx context.Context, year int, currency string) (map[string]float64, error) {
	query := `
    SELECT 
        "Company",
        COALESCE(SUM("value"), 0) AS total_value
    FROM 
        prod.tax_burden
    WHERE 
        "year" = $1 AND
        "currency" = $2 
    GROUP BY 
        "Company";
    `

	rows, err := r.Db.QueryContext(ctx, query, year, currency)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]float64)
	for rows.Next() {
		var nameShortRu string
		var totalValue float64
		if err := rows.Scan(&nameShortRu, &totalValue); err != nil {
			return nil, err
		}
		result[nameShortRu] = totalValue
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
