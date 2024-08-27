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
        "name_short_en",
        COALESCE(AVG("TotalTaxes" / NULLIF("Production", 0)), 0) AS average_specific_tax
    FROM 
        dmart.investments_dash
    WHERE 
        "report_year" = $1 AND
        "currencyunit" = $2 AND
        "report_type" = $3
    GROUP BY 
        "name_short_en";
`
	rows, err := r.Db.QueryContext(ctx, query, year, currency, reporttype)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]float64)
	for rows.Next() {
		var nameShortRu string
		var totalSpecificTax sql.NullFloat64 // Use sql.NullFloat64 to handle potential NULLs
		if err := rows.Scan(&nameShortRu, &totalSpecificTax); err != nil {
			return nil, err
		}
		if totalSpecificTax.Valid { // Check if the value is valid
			result[nameShortRu] = totalSpecificTax.Float64
		} else {
			result[nameShortRu] = 0.0 // Handle NULLs as 0.0
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
