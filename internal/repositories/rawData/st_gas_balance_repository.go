package repositories

import (
	"context"
	"database/sql"
	models "tender/internal/models/rawData"
)

type StGasBalanceRepository struct {
	Db *sql.DB
}

func (r *StGasBalanceRepository) GetGasBalance(ctx context.Context, year int64) ([]models.GasBalanceSummary, error) {
	query := `
		SELECT
		    year,
		    SUM(CASE WHEN item_group = 'total_gas_production' THEN value ELSE 0 END) AS Gas_Production,
		    SUM(CASE WHEN item_group = 'import' THEN value ELSE 0 END) AS Import,
		    SUM(CASE WHEN item_group = 'gas_sales_export' THEN value ELSE 0 END) AS Export,
		    SUM(CASE WHEN item_group = 'own_needs_losses' THEN value ELSE 0 END) AS Personal_Needs,
		    SUM(CASE WHEN item_group = 'reinjection' THEN value ELSE 0 END) AS Reinjection_into_Reservoir,
		    SUM(CASE WHEN item_group = 'gas_sales_domestic' THEN value ELSE 0 END) AS Gas_Sales
		FROM
		    raw_data.st_gas_balance
		WHERE
		    year BETWEEN 2021 AND $1
		GROUP BY
		    year
		ORDER BY
		    year ASC;
	`

	rows, err := r.Db.QueryContext(ctx, query, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summaries []models.GasBalanceSummary
	for rows.Next() {
		var summary models.GasBalanceSummary
		if err := rows.Scan(
			&summary.Year,
			&summary.TotalGasProduction,
			&summary.Import,
			&summary.Export,
			&summary.PersonalNeeds,
			&summary.ReinjectionIntoReservoir,
			&summary.GasSales,
		); err != nil {
			return nil, err
		}
		summaries = append(summaries, summary)
	}

	return summaries, nil
}
