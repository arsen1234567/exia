package repositories

import (
	"context"
	"database/sql"
	models "tender/internal/models/dmart" // Ensure this is correct
)

type DfoGgReportesRepository struct {
	Db *sql.DB
}

func (r *DfoGgReportesRepository) GetNetProfitSummary(ctx context.Context) ([]models.NetProfitSummary, error) {
	query := `
	SELECT 
		report_year AS year,
		SUM(period_end) AS net_profit
	FROM 
		dmart.dfo_qg_reports
	WHERE 
		item IN ('Прибыль за год (строка 200 + строка 201) относимая на:')
	GROUP BY 
		report_year
	ORDER BY 
		report_year ASC;
	`

	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.NetProfitSummary
	for rows.Next() {
		var summary models.NetProfitSummary
		if err := rows.Scan(&summary.Year, &summary.NetProfit); err != nil {
			return nil, err
		}
		results = append(results, summary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
