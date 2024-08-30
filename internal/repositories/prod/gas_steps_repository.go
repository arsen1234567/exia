package repositories

import (
	"context"
	"database/sql"
	"errors"
)

type GasStepsRepository struct {
	Db *sql.DB
}

func (r *GasStepsRepository) GetAmountOfPredictedTaxes(ctx context.Context, currency string) (float64, error) {
	query := `
    SELECT 
        SUM(kpn_total) AS total_kpn
    FROM 
        prod.gas_steps 
    WHERE 
        username = 'a.sagynayev@kazenergy.com' AND
		currency = $1
    `

	var totalKpn sql.NullFloat64
	err := r.Db.QueryRowContext(ctx, query, currency).Scan(&totalKpn)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	if !totalKpn.Valid {
		return 0, errors.New("invalid result")
	}

	return totalKpn.Float64, nil
}

func (r *GasStepsRepository) GetEBITDAmargin(ctx context.Context) (float64, error) {
	query := `
    SELECT 
        (SUM(ebit) + SUM(depreciation_amortization) + SUM(qg_share)) / SUM(revenue_total) AS result
    FROM 
        prod.gas_steps
    WHERE 
        username = 'a.sagynayev@kazenergy.com'
    `

	var ebitdaMargin sql.NullFloat64
	err := r.Db.QueryRowContext(ctx, query).Scan(&ebitdaMargin)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	if !ebitdaMargin.Valid {
		return 0, errors.New("invalid result")
	}

	return ebitdaMargin.Float64, nil
}
