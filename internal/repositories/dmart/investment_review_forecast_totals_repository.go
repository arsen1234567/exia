package repositories

import (
	"context"
	"database/sql"
	"fmt"
)

type InvestmentReviewForecastTotalRepository struct {
	Db *sql.DB
}

func (r *InvestmentReviewForecastTotalRepository) GetInvestmentReviewForecastTotal(ctx context.Context) (float64, error) {
	query := `
	SELECT 
    SUM("NPV") AS total_npv
FROM 
    dmart.investment_review_forecast_totals
WHERE 
    scenario IN ('Forecast BBrent BCPI');
	`

	var GIRFTotal sql.NullFloat64
	err := r.Db.QueryRowContext(ctx, query).Scan(&GIRFTotal)
	if err != nil {
		return 0, err
	}

	if !GIRFTotal.Valid {
		return 0, fmt.Errorf("totalGovShare is NULL")
	}

	return GIRFTotal.Float64, nil
}
