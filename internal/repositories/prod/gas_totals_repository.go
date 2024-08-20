package repositories

import (
	"context"
	"database/sql"
)

type GasTotalsRepository struct {
	Db *sql.DB
}

// GetNPVplusTV retrieves the sum of npv_plus_terminal_val for a given username
func (r *GasTotalsRepository) GetNPVplusTV(ctx context.Context) (float64, error) {
	query := `
		SELECT 
			SUM(npv_plus_terminal_value)
		FROM 
			prod.gas_totals 
		WHERE 
			username IN ('a.sagynayev@kazenergy.com');
	`

	var totalNPVplusTV float64
	err := r.Db.QueryRowContext(ctx, query).Scan(&totalNPVplusTV)
	if err != nil {
		return 0, err
	}

	return totalNPVplusTV, nil
}
