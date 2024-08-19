package repositories

import (
	"context"
	"database/sql"
	models "tender/internal/models/dmart" // Ensure this is correct
)

type InvestmentReservesRepository struct {
	Db *sql.DB
}

func (r *InvestmentReservesRepository) GetInvestmentReservesSummary(ctx context.Context, year int, unit string) ([]models.InvestmentReservesSummary, error) {
	query := `
	SELECT
	SUM("Балансовые запасы на конец(А+В+С1)" + "Балансовые запасы на конец(С2)") AS total_sum,
	CASE
	  WHEN abd_scope THEN 'Покрытие АБД'
	  WHEN "Недропользователь" = 'ОБЩИЙ ФОНД РК' THEN 'Общий фонд РК'
	  ELSE 'Вне периметра'
	END AS category
  FROM dmart.investment_reserves
  WHERE "Тип" = 'Извлекаемые';
	`

	rows, err := r.Db.QueryContext(ctx, query, year, unit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.InvestmentReservesSummary
	for rows.Next() {
		var summary models.InvestmentReservesSummary
		if err := rows.Scan(&summary.CoverageScope, &summary.TotalValue); err != nil {
			return nil, err
		}
		results = append(results, summary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
