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
<<<<<<< HEAD
		SUM("Балансовые запасы на конец(А+В+С1)" + "Балансовые запасы на конец(С2)") AS total_sum,
	CASE
	  WHEN abd_scope THEN 'Покрытие АБД'
	  WHEN "Недропользователь" = 'ОБЩИЙ ФОНД РК' THEN 'Общий фонд РК'
	  ELSE 'Вне периметра'
	END AS category
  	FROM dmart.investment_reserves
  	WHERE "Тип" = 'Извлекаемые';
=======
		'Покрытие АБД' AS category,
		SUM("Балансовые запасы на конец(А+В+С1)" + "Балансовые запасы на конец(С2)") AS total_sum
	FROM dmart.investment_reserves
	WHERE "Тип" = 'Извлекаемые' AND abd_scope = true AND "year" = $1 AND "unit" = $2
	
	UNION ALL
	
	SELECT
		'Вне периметра' AS category,
		SUM("Балансовые запасы на конец(А+В+С1)" + "Балансовые запасы на конец(С2)") AS total_sum
	FROM dmart.investment_reserves
	WHERE "Тип" = 'Извлекаемые' AND abd_scope = false AND "Недропользователь" != 'ОБЩИЙ ФОНД РК' AND "year" = $1 AND "unit" = $2
	
	UNION ALL
	
	SELECT
		'Общий фонд РК' AS category,
		SUM("Балансовые запасы на конец(А+В+С1)" + "Балансовые запасы на конец(С2)") AS total_sum
	FROM dmart.investment_reserves
	WHERE "Тип" = 'Извлекаемые' AND "Недропользователь" = 'ОБЩИЙ ФОНД РК' AND "year" = $1 AND "unit" = $2;
>>>>>>> f1fc91e06cf9fb949406a24e3831868cdc4216bd
	`

	rows, err := r.Db.QueryContext(ctx, query, year, unit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.InvestmentReservesSummary
	for rows.Next() {
		var summary models.InvestmentReservesSummary
		var totalValue sql.NullFloat64 // Use sql.NullFloat64 to handle NULL values
		if err := rows.Scan(&summary.CoverageScope, &totalValue); err != nil {
			return nil, err
		}
		if totalValue.Valid {
			summary.TotalValue = &totalValue.Float64
		} else {
			summary.TotalValue = nil
		}
		results = append(results, summary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
