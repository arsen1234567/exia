package repositories

import (
	"context"
	"database/sql"
)

type SpecificTaxesRepository struct {
	Db *sql.DB
}

func (r *SpecificTaxesRepository) GetSpecificTaxes(ctx context.Context, productionunit string, year int) (map[string]float64, error) {
	query := `
    SELECT 
        "name_short_ru",
        COALESCE(SUM("specific_tax"), 0) AS total_specific_tax
    FROM 
        prod.specific_taxes
    WHERE 
        "ProductionUnit" = $1 AND
        "year" = $2 
    GROUP BY 
        "name_short_ru";
    `

	rows, err := r.Db.QueryContext(ctx, query, productionunit, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]float64)
	for rows.Next() {
		var nameShortRu string
		var totalSpecificTax float64
		if err := rows.Scan(&nameShortRu, &totalSpecificTax); err != nil {
			return nil, err
		}
		result[nameShortRu] = totalSpecificTax
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
