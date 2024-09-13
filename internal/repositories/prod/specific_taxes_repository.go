package repositories

import (
	"context"
	"database/sql"
	"fmt"
)

type SpecificTaxesRepository struct {
	Db *sql.DB
}

func (r *SpecificTaxesRepository) GetSpecificTaxes(ctx context.Context, year int, currency, reporttype, language string) (map[string]float64, map[string]float64, error) {
	query := fmt.Sprintf(`
    SELECT 
        %s,
        COALESCE(AVG("TotalTaxes" / NULLIF("Production", 0)), 0) AS average_specific_tax
    FROM 
        dmart.investments_dash
    WHERE 
        "report_year" = $1 AND
        "currencyunit" = $2 AND
        "report_type" = $3
    GROUP BY 
	%s;
`, language, language)

	query2 := `
SELECT 
	name_abbr,
	COALESCE(AVG("TotalTaxes" / NULLIF("Production", 0)), 0) AS average_specific_tax
FROM 
	dmart.investments_dash
WHERE 
	"report_year" = $1 AND
	"currencyunit" = $2 AND
	"report_type" = $3
GROUP BY 
name_abbr;
`

	rows, err := r.Db.QueryContext(ctx, query, year, currency, reporttype)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	result := make(map[string]float64)
	for rows.Next() {
		var nameShortRu string
		var totalSpecificTax sql.NullFloat64 // Use sql.NullFloat64 to handle potential NULLs
		if err := rows.Scan(&nameShortRu, &totalSpecificTax); err != nil {
			return nil, nil, err
		}
		if totalSpecificTax.Valid { // Check if the value is valid
			result[nameShortRu] = totalSpecificTax.Float64
		} else {
			result[nameShortRu] = 0.0 // Handle NULLs as 0.0
		}
	}

	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	rows2, err := r.Db.QueryContext(ctx, query2, year, currency, reporttype)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	result2 := make(map[string]float64)
	for rows2.Next() {
		var nameShortRu string
		var totalSpecificTax sql.NullFloat64 // Use sql.NullFloat64 to handle potential NULLs
		if err := rows2.Scan(&nameShortRu, &totalSpecificTax); err != nil {
			return nil, nil, err
		}
		if totalSpecificTax.Valid { // Check if the value is valid
			result2[nameShortRu] = totalSpecificTax.Float64
		} else {
			result2[nameShortRu] = 0.0 // Handle NULLs as 0.0
		}
	}

	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return result, result2, nil
}
