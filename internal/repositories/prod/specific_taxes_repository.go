package repositories

import (
	"context"
	"database/sql"
)

type SpecificTaxesRepository struct {
	Db *sql.DB
}

func (r *SpecificTaxesRepository) GetSpecificTaxes(ctx context.Context, year int, currency, reporttype string) (map[string]float64, error) {
	query := `
    SELECT 
        "name_short_ru",
        COALESCE(AVG("TotalTaxes" / NULLIF("Production", 0)), 0) AS average_specific_tax
    FROM 
        dmart.investments_dash
    WHERE 
        "report_year" = $1 AND
        "currencyunit" = $2 AND
        "report_type" = $3
    GROUP BY 
        "name_short_ru";
`
	rows, err := r.Db.QueryContext(ctx, query, year, currency, reporttype)
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
