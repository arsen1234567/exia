package repositories

import (
	"context"
	"database/sql"
)

type NaturalGasMainRepository struct {
	Db *sql.DB
}

func (r *NaturalGasMainRepository) GetReserveRatio(ctx context.Context) (int64, error) {
	query := `
	SELECT 
		"Кратность запасов"
	FROM 
		dmart.natural_gas_main ngm 
	LIMIT 1;
	`

	var reserveRatio int64
	err := r.Db.QueryRowContext(ctx, query).Scan(&reserveRatio)
	if err != nil {
		return 0, err
	}

	return reserveRatio, nil
}
