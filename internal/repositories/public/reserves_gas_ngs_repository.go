package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	models "tender/internal/models/rawData"
)

type ReservesGasNgsRepository struct {
	Db *sql.DB
}

func (r *ReservesGasNgsRepository) GetReservesGasNgsDeposit(ctx context.Context) ([]models.NgsReservesGasSummary, error) {
	query := `
		SELECT EXTRACT(YEAR FROM "Год") AS year, COUNT(DISTINCT "Месторождения в алфавитном справ.") AS count
		FROM raw_data.ngs_reserves_gas
		WHERE "Год" >= '1990-01-01' AND "Год" <= '2023-01-01'
		GROUP BY year
		ORDER BY year;
	`

	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var results []models.NgsReservesGasSummary
	for rows.Next() {
		var result models.NgsReservesGasSummary
		if err := rows.Scan(&result.Year, &result.AlphRegionCount); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *ReservesGasNgsRepository) GetReservesGasNgsNumberOfCompanies(ctx context.Context) ([]models.NgsReservesGasSummary, error) {
	query := `
		SELECT EXTRACT(YEAR FROM "Год") AS year, COUNT(DISTINCT "Недропользователь") AS count
		FROM raw_data.ngs_reserves_gas
		WHERE "Год" >= '1990-01-01' AND "Год" <= '2023-01-01'
		GROUP BY year
		ORDER BY year;
	`

	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var results []models.NgsReservesGasSummary
	for rows.Next() {
		var result models.NgsReservesGasSummary
		if err := rows.Scan(&result.Year, &result.AlphRegionCount); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *ReservesGasNgsRepository) GetTotalReserves(ctx context.Context, oilType string) ([]models.NgsReservesGasTotalReservesSummary, error) {
	query := `
		SELECT EXTRACT(YEAR FROM "Год") AS year, 
		       SUM("Балансовые запасы на начало(А+В+С1)" + "Балансовые запасы на начало(С2)") AS total_reserves_sum
		FROM raw_data.ngs_reserves_gas
		WHERE "Год" >= '1990-01-01' AND "Год" <= '2023-01-01'
		  AND "Тип" = $1
		GROUP BY year
		ORDER BY year;
	`

	rows, err := r.Db.QueryContext(ctx, query, oilType)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var results []models.NgsReservesGasTotalReservesSummary
	for rows.Next() {
		var result models.NgsReservesGasTotalReservesSummary
		if err := rows.Scan(&result.Year, &result.TotalReservesSum); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *ReservesGasNgsRepository) GetProduction(ctx context.Context, oilType string) ([]models.NgsReservesGasTotalReservesSummary, error) {
	query := `
		SELECT EXTRACT(YEAR FROM "Год") AS year, 
		       SUM("Изм. Бал. Зап. добычи(А+В+С1)") AS total_reserves_sum
		FROM raw_data.ngs_reserves_gas
		WHERE "Год" >= '1990-01-01' AND "Год" <= '2023-01-01'
		  AND "Тип" = $1
		GROUP BY year
		ORDER BY year;
	`

	rows, err := r.Db.QueryContext(ctx, query, oilType)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var results []models.NgsReservesGasTotalReservesSummary
	for rows.Next() {
		var result models.NgsReservesGasTotalReservesSummary
		if err := rows.Scan(&result.Year, &result.TotalReservesSum); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *ReservesGasNgsRepository) GetNumberOfDepositsByRegion(ctx context.Context, year int) ([]models.DepositsByRegionSummary, error) {
	query := `
        SELECT "Регион", COUNT(DISTINCT "Месторождения в алфавитном справ.") 
        FROM raw_data.ngs_reserves_gas
        WHERE EXTRACT(YEAR FROM "Год") = $1
        GROUP BY "Регион"
        ORDER BY "Регион";
    `

	rows, err := r.Db.QueryContext(ctx, query, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.DepositsByRegionSummary
	for rows.Next() {
		var regionCount models.DepositsByRegionSummary
		if err := rows.Scan(&regionCount.Region, &regionCount.FieldCount); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		results = append(results, regionCount)
	}

	return results, nil
}

func (r *ReservesGasNgsRepository) GetTopCompaniesByReserves(ctx context.Context, oilType string, year int) ([]models.NgsReservesGasTopCompanies, error) {
	query := `
		SELECT 
			"Недропользователь" AS company_name,
			SUM("Балансовые запасы на конец(А+В+С1)" + "Балансовые запасы на конец(С2)") AS total_reserves
		FROM 
		    raw_data.ngs_reserves_gas
		WHERE 
			DATE_TRUNC('year', "Год") = TO_TIMESTAMP($1::TEXT, 'YYYY')
			AND "Тип" = $2
		GROUP BY 
			"Недропользователь"
		ORDER BY 
			total_reserves DESC;
	`

	rows, err := r.Db.QueryContext(ctx, query, year, oilType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.NgsReservesGasTopCompanies
	for rows.Next() {
		var summary models.NgsReservesGasTopCompanies
		if err := rows.Scan(&summary.CompanyName, &summary.TotalReserves); err != nil {
			return nil, err
		}
		results = append(results, summary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
