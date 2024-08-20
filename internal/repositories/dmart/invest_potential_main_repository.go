package repositories

import (
	"context"
	"database/sql"
	// Ensure this is correct
)

type InvestPotentialMainRepository struct {
	Db *sql.DB
}

func (r *InvestPotentialMainRepository) GetInvestPotentialMain(ctx context.Context) (float64, error) {
	query := `
	SELECT COALESCE(SUM("Кратность запасов"), 0) AS total_reserve_multiple
	FROM dmart.invest_potential_main;
	`
	row := r.Db.QueryRowContext(ctx, query)

	var totalReserveMultiple float64
	if err := row.Scan(&totalReserveMultiple); err != nil {
		return 0, err
	}

	return totalReserveMultiple, nil
}

func (r *InvestPotentialMainRepository) GetSpecOpEx(ctx context.Context) (float64, error) {
	query := `
	SELECT 
    SUM("Удельные операционные расходы") AS total_operational_expenses
FROM 
    dmart.invest_potential_main;
	`
	row := r.Db.QueryRowContext(ctx, query)

	var totalReserveMultiple float64
	if err := row.Scan(&totalReserveMultiple); err != nil {
		return 0, err
	}

	return totalReserveMultiple, nil
}
