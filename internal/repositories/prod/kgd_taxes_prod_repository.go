package repositories

import (
	"context"
	"database/sql"
	models "tender/internal/models/prod" // Ensure this is correct
)

type KgdTaxesProdRepository struct {
	Db *sql.DB
}

func (r *KgdTaxesProdRepository) GetKgdTaxesProdSummaryForYearRange(ctx context.Context, startYear, endYear int, currency string) ([]models.KgdTaxesProdSummary, error) {
	query := `
	SELECT 
		SUM(summa) AS total_value, 
		EXTRACT(YEAR FROM receipt_date) AS receipt_year
	FROM 
		prod.kgd_taxes_prod 
	WHERE 
		EXTRACT(YEAR FROM receipt_date) BETWEEN $1 AND $2 AND
		currency = $3
	GROUP BY 
		receipt_year
	ORDER BY 
		receipt_year;
	`

	rows, err := r.Db.QueryContext(ctx, query, startYear, endYear, currency)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]models.KgdTaxesProdSummary, 0)
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

func (r *KgdTaxesProdRepository) GetSummaAllTaxes(ctx context.Context, year int, currency, reporttype string) (map[string]float64, error) {
	query := `
    SELECT 
        "name_short_en",
        COALESCE(SUM("TotalTaxes"), 0) AS total_value
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
