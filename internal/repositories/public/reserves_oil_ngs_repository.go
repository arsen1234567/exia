package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	models "tender/internal/models/public"
	rawModels "tender/internal/models/rawData"
)

type ReservesOilNgsRepository struct {
	Db *sql.DB
}

func (r *ReservesOilNgsRepository) GetReservesOilNgsDeposit(ctx context.Context) ([]models.NgsReservesOilSummary, error) {
	query := `
		SELECT EXTRACT(YEAR FROM "Год") AS year, COUNT(DISTINCT "Месторождения в алфавитном справ.") AS count
		FROM public.reserves_oil_ngs
		WHERE "Год" >= '1990-01-01' AND "Год" <= '2023-01-01'
		GROUP BY year
		ORDER BY year;
	`

	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var results []models.NgsReservesOilSummary
	for rows.Next() {
		var result models.NgsReservesOilSummary
		if err := rows.Scan(&result.Year, &result.AlphRegionCount); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *ReservesOilNgsRepository) GetReservesOilNgsNumberOfCompanies(ctx context.Context) ([]models.NgsReservesOilSummary, error) {
	query := `
		SELECT EXTRACT(YEAR FROM "Год") AS year, COUNT(DISTINCT "Недропользователь") AS count
		FROM public.reserves_oil_ngs
		WHERE "Год" >= '1990-01-01' AND "Год" <= '2023-01-01'
		GROUP BY year
		ORDER BY year;
	`

	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var results []models.NgsReservesOilSummary
	for rows.Next() {
		var result models.NgsReservesOilSummary
		if err := rows.Scan(&result.Year, &result.AlphRegionCount); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *ReservesOilNgsRepository) GetReservesOilNgsTotalProduction(ctx context.Context, oilType string) ([]rawModels.ReservesOilNgsTotalProductionSummary, error) {
	query := `
        SELECT EXTRACT(YEAR FROM "Год")::INT AS year, SUM("Добычи, потери(А+В+С1)") 
        FROM raw_data.ngs_reserves_oil 
        WHERE "Тип" = $1 AND "Год" BETWEEN '1990-01-01' AND '2022-12-31'
        GROUP BY year
        ORDER BY year
    `

	rows, err := r.Db.QueryContext(ctx, query, oilType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []rawModels.ReservesOilNgsTotalProductionSummary
	for rows.Next() {
		var summary rawModels.ReservesOilNgsTotalProductionSummary
		if err := rows.Scan(&summary.Year, &summary.Sum); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		results = append(results, summary)
	}

	return results, nil
}

func (r *ReservesOilNgsRepository) GetNumberOfDepositsByRegion(ctx context.Context, year int) ([]rawModels.DepositsByRegionSummary, error) {
	query := `
        SELECT "Регион", COUNT(DISTINCT "Месторождения в алфавитном справ.") 
        FROM raw_data.ngs_reserves_oil 
        WHERE EXTRACT(YEAR FROM "Год") = $1
        GROUP BY "Регион"
        ORDER BY "Регион";
    `

	rows, err := r.Db.QueryContext(ctx, query, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []rawModels.DepositsByRegionSummary
	for rows.Next() {
		var regionCount rawModels.DepositsByRegionSummary
		if err := rows.Scan(&regionCount.Region, &regionCount.FieldCount); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		results = append(results, regionCount)
	}

	return results, nil
}

func (r *ReservesOilNgsRepository) GetTopCompaniesByReserves(ctx context.Context, oilType string, year int) ([]models.NgsReservesOilTopCompanies, error) {
	query := `
		SELECT 
			"Недропользователь" AS company_name,
			SUM("Балансовые запасы на конец(А+В+С1)" + "Балансовые запасы на конец(С2)") AS total_reserves
		FROM 
			public.reserves_oil_ngs
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

	var results []models.NgsReservesOilTopCompanies
	for rows.Next() {
		var summary models.NgsReservesOilTopCompanies
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
