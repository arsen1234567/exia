package repositories

import (
	"context"
	"database/sql"
	models "tender/internal/models/dmart"
)

type DfoQazaqgasRepository struct {
	Db *sql.DB
}

// GetServiceRevenueByCompanyAndYear fetches the service revenues by company and year, including the total value
func (r *DfoQazaqgasRepository) GetRevenueByServicesAndCompanyAndYear(ctx context.Context, company string, year int) (models.RevenueByServicesSummary, error) {
	query := `
		SELECT 
			item_eng, 
			SUM(period_end) AS total_period_end
		FROM 
			dmart.dfo_qazaqgas
		WHERE 
			type_eng IN ('Type of Goods and Services')
			AND company = $1
			AND report_year = $2
		GROUP BY 
			item_eng;
	`

	rows, err := r.Db.QueryContext(ctx, query, company, year)
	if err != nil {
		return models.RevenueByServicesSummary{}, err
	}
	defer rows.Close()

	var summary models.RevenueByServicesSummary
	for rows.Next() {
		var itemEng string
		var totalPeriodEnd float64
		if err := rows.Scan(&itemEng, &totalPeriodEnd); err != nil {
			return models.RevenueByServicesSummary{}, err
		}

		// Map each item_eng to the appropriate field in RevenueByServicesSummary
		switch itemEng {
		case "Revenue from gas transportation services":
			summary.RevenueFromGasTransportationServices = totalPeriodEnd
		case "Revenue from gas pipeline maintenance":
			summary.RevenueFromGasPipelineMaintenance = totalPeriodEnd
		case "Other":
			summary.Other = totalPeriodEnd
		case "Gas sales revenue":
			summary.GasSalesRevenue = totalPeriodEnd
		case "Managerial fee":
			summary.ManagerialFee = totalPeriodEnd
		case "Total":
			summary.TotalRevenue = totalPeriodEnd
		}
	}

	if err = rows.Err(); err != nil {
		return models.RevenueByServicesSummary{}, err
	}

	return summary, nil
}

// GetRevenueByGeographyAndCompanyAndYear fetches the revenues by geography and company, including the total value
func (r *DfoQazaqgasRepository) GetRevenueByGeographyAndCompanyAndYear(ctx context.Context, company string, year int) (models.RevenueByGeographySummary, error) {
	query := `
		SELECT 
			item_eng, 
			SUM(period_end) AS total_period_end
		FROM 
			dmart.dfo_qazaqgas
		WHERE 
			type_eng IN ('Geographic markets')
			AND company = $1
			AND report_year = $2
		GROUP BY 
			item_eng;
	`

	rows, err := r.Db.QueryContext(ctx, query, company, year)
	if err != nil {
		return models.RevenueByGeographySummary{}, err
	}
	defer rows.Close()

	var summary models.RevenueByGeographySummary
	for rows.Next() {
		var itemEng string
		var totalPeriodEnd float64
		if err := rows.Scan(&itemEng, &totalPeriodEnd); err != nil {
			return models.RevenueByGeographySummary{}, err
		}

		// Map each item_eng to the appropriate field in RevenueByGeographySummary
		switch itemEng {
		case "China":
			summary.China = totalPeriodEnd
		case "Kazakhstan":
			summary.Kazakhstan = totalPeriodEnd
		case "CIS":
			summary.CIS = totalPeriodEnd
		case "Total":
			summary.TotalRevenue = totalPeriodEnd
		}
	}

	if err = rows.Err(); err != nil {
		return models.RevenueByGeographySummary{}, err
	}

	return summary, nil
}

// GetCostItemsByCompanyAndYear fetches the cost items by company and year, including the total value
func (r *DfoQazaqgasRepository) GetCostItemsByCompanyAndYear(ctx context.Context, company string, year int) (models.CostItemsSummary, error) {
	query := `
		SELECT 
			item_eng, 
			SUM(period_end) AS total_period_end
		FROM 
			dmart.dfo_qazaqgas
		WHERE 
			type_eng IN ('Expense Items')
			AND company = $1
			AND report_year = $2
		GROUP BY 
			item_eng
		ORDER BY 
			total_period_end DESC;
	`

	rows, err := r.Db.QueryContext(ctx, query, company, year)
	if err != nil {
		return models.CostItemsSummary{}, err
	}
	defer rows.Close()

	var summary models.CostItemsSummary
	for rows.Next() {
		var itemEng string
		var totalPeriodEnd float64
		if err := rows.Scan(&itemEng, &totalPeriodEnd); err != nil {
			return models.CostItemsSummary{}, err
		}

		// Map each item_eng to the appropriate field in CostItemsSummary
		switch itemEng {
		case "Cost of sold gas":
			summary.CostOfSoldGas = totalPeriodEnd
		case "Transport costs":
			summary.TransportCosts = totalPeriodEnd
		case "Salary and related deductions":
			summary.SalaryAndRelatedDeductionsCosts = totalPeriodEnd
		case "other expenses":
			summary.OtherExpenses = totalPeriodEnd
		case "Depreciation and Amortization":
			summary.DepreciationAndAmortization = totalPeriodEnd
		case "Total":
			summary.TotalCosts = totalPeriodEnd
		}
	}

	if err = rows.Err(); err != nil {
		return models.CostItemsSummary{}, err
	}

	return summary, nil
}
