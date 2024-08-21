package repositories

import (
	"context"
	"database/sql"
	models "tender/internal/models/prod" // Ensure this is correct
)

type KgdTaxesProdRepository struct {
	Db *sql.DB
}

func (r *KgdTaxesProdRepository) GetKgdTaxesProdSummary(ctx context.Context, year int) ([]models.KgdTaxesProdSummary, error) {
	query := `
	SELECT 
    SUM(summa) AS total_value, 
    EXTRACT(YEAR FROM receipt_date) AS receipt_year
	FROM 
    prod.kgd_taxes_prod 
	WHERE 
    EXTRACT(YEAR FROM receipt_date) BETWEEN 2015 AND LEAST($1, 2023) 
	GROUP BY 
    receipt_year
	ORDER BY 
    receipt_year;
`

	rows, err := r.Db.QueryContext(ctx, query, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.KgdTaxesProdSummary
	for rows.Next() {
		var summary models.KgdTaxesProdSummary
		if err := rows.Scan(&summary.TotalValue, &summary.Year); err != nil {
			return nil, err
		}
		results = append(results, summary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *KgdTaxesProdRepository) GetSummaAllTaxes(ctx context.Context, year int, currency string) (map[string]float64, error) {
	query := `
    SELECT 
        "Company",
        COALESCE(SUM("summa"), 0) AS total_value
    FROM 
        prod.kgd_taxes_prod
    WHERE 
        EXTRACT(YEAR FROM "receipt_date") = $1 AND
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
		var companyName string
		var totalValue float64
		if err := rows.Scan(&companyName, &totalValue); err != nil {
			return nil, err
		}
		result[companyName] = totalValue
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
