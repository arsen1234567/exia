package repositories

import (
	"context"
	"database/sql"
	models "tender/internal/models/rawData"
)

type KgdTaxesRepository struct {
	Db *sql.DB
}

func (r *KgdTaxesRepository) GetKgdTaxesSummary(ctx context.Context, year int, currency string) ([]models.KgdTaxesSummary, error) {
	query := `
	SELECT 
		EXTRACT(YEAR FROM receipt_date) AS year,
		SUM(summa) AS total_sum
	FROM 
		raw_data.kgd_taxes
	WHERE 
		receipt_date < '2024-01-01'
		AND bin IN ('000340002165','050840002757', 
		            '970740000392','020440001144',
		            '060640006784','101140017122',
		            '110140008803','150441026419',
		            '050140002494','190240017187',
		            '030940002310','060140015913',
		            '201240027449','031040006125',
		            '080240013062')
		AND pay_status = '2 - Разнесен'
		AND pay_type = 'Налог'
		AND EXTRACT(YEAR FROM receipt_date) <= CAST($1 AS INTEGER)
		AND currency = $2
	GROUP BY 
		year
	ORDER BY 
		year;
	`

	rows, err := r.Db.QueryContext(ctx, query, year, currency)
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
