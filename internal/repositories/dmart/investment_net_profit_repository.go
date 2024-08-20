package repositories

import (
	"context"
	"database/sql"
	models "tender/internal/models/dmart" // Ensure this is correct
)

type InvestmentNetProfitRepository struct {
	Db *sql.DB
}

func (r *InvestmentNetProfitRepository) GetInvestmentNetProfitSummary(ctx context.Context, year int) ([]models.InvestmentNetProfitSummary, error) {
	query := `
	SELECT 
    "Year",
    SUM("NetProfit") AS total_sum
FROM 
    dmart.investment_net_profit
WHERE 
    "Year" BETWEEN 2008 AND $1
GROUP BY 
    "Year"
ORDER BY 
    "Year";
`
	rows, err := r.Db.QueryContext(ctx, query, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.InvestmentNetProfitSummary
	for rows.Next() {
		var summary models.InvestmentNetProfitSummary
		if err := rows.Scan(&summary.Year, &summary.TotalSum); err != nil {
			return nil, err
		}
		results = append(results, summary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
