package repositories

import (
	"context"
	"database/sql"
	models "tender/internal/models/dmart" // Ensure this is correct
)

type DfoGgReportesRepository struct {
	Db *sql.DB
}

func (r *DfoGgReportesRepository) GetNetProfitSummary(ctx context.Context) ([]models.NetProfitSummary, error) {
	query := `
	SELECT 
		report_year AS year,
		SUM(period_end) AS net_profit
	FROM 
		dmart.dfo_qg_reports
	WHERE 
		item IN ('Прибыль за год (строка 200 + строка 201) относимая на:')
	GROUP BY 
		report_year
	ORDER BY 
		report_year ASC;
	`

	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.NetProfitSummary
	for rows.Next() {
		var summary models.NetProfitSummary
		if err := rows.Scan(&summary.Year, &summary.NetProfit); err != nil {
			return nil, err
		}
		results = append(results, summary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *DfoGgReportesRepository) GetRevenueByCompanyAndYear(ctx context.Context, company string, year int) (float64, error) {
	query := `
	SELECT 
		SUM(period_end) AS total_period_end
	FROM 
		dmart.dfo_qg_reports
	WHERE 
		item_en IN ('Revenue from sales of goods, works and services')
		AND company = $1
		AND report_year = $2
	`

	var totalPeriodEnd float64
	err := r.Db.QueryRowContext(ctx, query, company, year).Scan(&totalPeriodEnd)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return totalPeriodEnd, nil
}

func (r *DfoGgReportesRepository) GetCostOfGoodsWorksServicesSold(ctx context.Context, company string, year int) (float64, error) {
	query := `
		SELECT 
			SUM(period_end) AS total_period_end
		FROM 
			dmart.dfo_qg_reports 
		WHERE 
			item_en IN ('Cost of goods, works and services sold')
			AND company = $1
			AND report_year = $2
	`

	var totalCost float64
	err := r.Db.QueryRowContext(ctx, query, company, year).Scan(&totalCost)
	if err != nil {
		return 0, err
	}

	return totalCost, nil
}

func (r *DfoGgReportesRepository) GetGrossProfit(ctx context.Context, company string, year int) (float64, error) {
	query := `
		SELECT 
			SUM(period_end) AS total_period_end
		FROM 
			dmart.dfo_qg_reports 
		WHERE 
			item_en IN ('Gross profit (loss) (line 010 – line 011)')
			AND company = $1
			AND report_year = $2
	`

	var grossProfit float64
	err := r.Db.QueryRowContext(ctx, query, company, year).Scan(&grossProfit)
	if err != nil {
		return 0, err
	}

	return grossProfit, nil
}

func (r *DfoGgReportesRepository) GetCIT(ctx context.Context, company string, year int) (float64, error) {
	query := `
		SELECT 
			SUM(period_start) AS total_cit
		FROM 
			dmart.dfo_qg_reports dqr 
		WHERE 
			item_en IN ('Expenses (-) (income (+)) for income tax')
			AND company = $1
			AND report_year = $2
		GROUP BY 
			item_en;
	`

	var totalCIT float64
	err := r.Db.QueryRowContext(ctx, query, company, year).Scan(&totalCIT)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return totalCIT, nil
}
