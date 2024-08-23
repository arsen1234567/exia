package repositories

import (
	"context"
	"database/sql"
	"log"
)

type InvestmentsDashRepository struct {
	Db *sql.DB
}

func (r *InvestmentsDashRepository) GetInvestmentsDash(ctx context.Context, companyName, finReportType string, reportYear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		COALESCE(SUM("Balance(А+В+С1)" + "Balance(С2)"), 0.0) AS total_balance
	FROM 
		dmart.investments_dash
	WHERE 
		"name_short_en" = $1 AND
		"report_type" = $2 AND
		"report_year" <= $3 AND
		"ProductionUnit" = 'barrels' AND
		"currencyunit" = 'USD' 
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, companyName, finReportType, reportYear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make(map[int]float64)
	for rows.Next() {
		var year int
		var totalBalance float64
		if err := rows.Scan(&year, &totalBalance); err != nil {
			return nil, err
		}
		results[year] = totalBalance
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashOilProduction(ctx context.Context, companyname, productionunit, finreporttype string, reportYear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		SUM("Production") AS total_production
	FROM 
		dmart.investments_dash
	WHERE 
		"name_short_en" = $1 AND
		"ProductionUnit" = $2 AND
		"report_type" = $3 AND
		"report_year" <= $4 AND
		"currencyunit" = 'KZT' 
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, companyname, productionunit, finreporttype, reportYear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make(map[int]float64)
	for rows.Next() {
		var year int
		var totalProduction float64
		if err := rows.Scan(&year, &totalProduction); err != nil {
			return nil, err
		}
		results[year] = totalProduction
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashSpecificRevenue(ctx context.Context, currencyunit, companyname, productionunit, finreporttype string, reportyear int) (map[int]float64, error) {
	log.Println("currencyunit ", currencyunit)
	log.Println("companyname", companyname)
	log.Println("productionunit", productionunit)
	log.Println("reportyear", reportyear)
	log.Println("finreporttype", finreporttype)
	query := `
	SELECT 
		"report_year",
		COALESCE(AVG("Revenue" / NULLIF("Production", 0)), 0) AS avg_specific_revenue
	FROM 
		dmart.investments_dash
	WHERE 
		"currencyunit" = $1 AND
		"name_short_en" = $2 AND
		"ProductionUnit" = $3 AND
		"report_type" = $4 AND
		"report_year" <= $5
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, currencyunit, companyname, productionunit, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]float64)
	for rows.Next() {
		var year int
		var avgSpecificRevenue float64
		if err := rows.Scan(&year, &avgSpecificRevenue); err != nil {
			return nil, err
		}
		result[year] = avgSpecificRevenue
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashROA(ctx context.Context, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		COALESCE(AVG("NetProfit" / NULLIF("ShortAssets" + "LongAssets" + "ShortAssetsSale", 0)), 0) AS avg_roa
	FROM 
		dmart.investments_dash
	WHERE 
		"name_short_en" = $1 AND
		"report_type" = $2 AND
		"report_year" BETWEEN 2010 AND $3
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]float64)
	for rows.Next() {
		var year int
		var avgROA float64
		if err := rows.Scan(&year, &avgROA); err != nil {
			return nil, err
		}
		result[year] = avgROA
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashNetProfitMargin(ctx context.Context, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		COALESCE(AVG("NetProfit" / NULLIF("Revenue", 0)), 0) AS avg_net_profit_to_revenue
	FROM 
		dmart.investments_dash
	WHERE 
		"name_short_en" = $1 AND
		"report_type" = $2 AND
		"report_year" <= $3
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]float64)
	for rows.Next() {
		var year int
		var avgNetProfitToRevenue float64
		if err := rows.Scan(&year, &avgNetProfitToRevenue); err != nil {
			return nil, err
		}
		result[year] = avgNetProfitToRevenue
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashSpecificNetProfit(ctx context.Context, currencyunit, companyname, productionunit, finreporttype string, reportyear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		COALESCE(AVG("NetProfit" / NULLIF("Production", 0)), 0) AS avg_net_profit_to_revenue
	FROM 
		dmart.investments_dash
	WHERE 
		"currencyunit" = $1 AND
		"name_short_en" = $2 AND
		"ProductionUnit" = $3 AND
		"report_type" = $4 AND
		"report_year" <= $5
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, currencyunit, companyname, productionunit, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]float64)
	for rows.Next() {
		var year int
		var avgNetProfitToRevenue float64
		if err := rows.Scan(&year, &avgNetProfitToRevenue); err != nil {
			return nil, err
		}
		result[year] = avgNetProfitToRevenue
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashRevenue(ctx context.Context, currencyunit, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		COALESCE(SUM("Revenue"), 0) AS total_revenue
	FROM 
		dmart.investments_dash
	WHERE 
		"currencyunit" = $1 AND
		"name_short_en" = $2 AND
		"report_type" = $3 AND
		"report_year" <= $4
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, currencyunit, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]float64)
	for rows.Next() {
		var year int
		var totalRevenue float64
		if err := rows.Scan(&year, &totalRevenue); err != nil {
			return nil, err
		}
		result[year] = totalRevenue
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashOperatingProfit(ctx context.Context, currencyunit, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		COALESCE(SUM("OperatingProfit"), 0) AS total_operating_profit
	FROM 
		dmart.investments_dash
	WHERE 
		"currencyunit" = $1 AND
		"name_short_en" = $2 AND
		"report_type" = $3 AND
		"ProductionUnit" = 'tons' AND
		"report_year" <= $4
		
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, currencyunit, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]float64)
	for rows.Next() {
		var year int
		var totalOperatingProfit float64
		if err := rows.Scan(&year, &totalOperatingProfit); err != nil {
			return nil, err
		}
		result[year] = totalOperatingProfit
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashEBITDA(ctx context.Context, currencyunit, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		COALESCE(SUM("EBT" + "FinExpenses" + "DepreciationAndAmortization"), 0) AS total_sum
	FROM 
		dmart.investments_dash
	WHERE 
		"currencyunit" = $1 AND
		"name_short_en" = $2 AND
		"ProductionUnit" = 'tons' AND
		"report_type" = $3 AND
		"report_year" <= $4
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, currencyunit, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]float64)
	for rows.Next() {
		var year int
		var totalSum float64
		if err := rows.Scan(&year, &totalSum); err != nil {
			return nil, err
		}
		result[year] = totalSum
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashNetProfit(ctx context.Context, currencyunit, companyname, finreporttype string, reportYear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		SUM("NetProfit") AS total_production
	FROM 
		dmart.investments_dash
	WHERE 
		"currencyunit" = $1 AND
		"name_short_en" = $2 AND
		"ProductionUnit" = 'tons' AND
		"report_type" = $3 AND
		"report_year" <= $4
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, currencyunit, companyname, finreporttype, reportYear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make(map[int]float64)
	for rows.Next() {
		var year int
		var totalProduction float64
		if err := rows.Scan(&year, &totalProduction); err != nil {
			return nil, err
		}
		results[year] = totalProduction
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashTotalTaxes(ctx context.Context, currencyunit, companyname, finreporttype string, reportYear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		COALESCE(SUM("TotalTaxes"), 0) AS total_taxes
	FROM 
		dmart.investments_dash
	WHERE 
		"currencyunit" = $1 AND
		"name_short_en" = $2 AND
		"ProductionUnit" = 'tons' AND
		"report_type" = $3 AND
		"report_year" BETWEEN 2012 AND $4
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, currencyunit, companyname, finreporttype, reportYear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make(map[int]float64)
	for rows.Next() {
		var year int
		var totalTaxes sql.NullFloat64
		if err := rows.Scan(&year, &totalTaxes); err != nil {
			return nil, err
		}
		if totalTaxes.Valid {
			results[year] = totalTaxes.Float64
		} else {
			results[year] = 0 // or handle as needed
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashTaxBurden(ctx context.Context, currencyunit, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		COALESCE(AVG("TotalTaxes" / NULLIF("Revenue", 0)), 0) AS avg_net_profit_to_revenue
	FROM 
		dmart.investments_dash
	WHERE 
		"currencyunit" = $1 AND
		"name_short_en" = $2 AND
		"ProductionUnit" = 'tons' AND
		"report_type" = $3 AND
		"report_year" BETWEEN 2012 and $4
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, currencyunit, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]float64)
	for rows.Next() {
		var year int
		var avgNetProfitToRevenue float64
		if err := rows.Scan(&year, &avgNetProfitToRevenue); err != nil {
			return nil, err
		}
		result[year] = avgNetProfitToRevenue
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashSpecificTaxes(ctx context.Context, currencyunit, companyname, productionunit, finreporttype string, reportyear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		COALESCE(AVG("TotalTaxes" / NULLIF("Production", 0)), 0) AS avg_net_profit_to_revenue
	FROM 
		dmart.investments_dash
	WHERE 
		"currencyunit" = $1 AND
		"name_short_en" = $2 AND
		"ProductionUnit" = $3 AND
		"report_type" = $4 AND
		"report_year" BETWEEN 2012 and $5
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, currencyunit, companyname, productionunit, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]float64)
	for rows.Next() {
		var year int
		var avgNetProfitToRevenue float64
		if err := rows.Scan(&year, &avgNetProfitToRevenue); err != nil {
			return nil, err
		}
		result[year] = avgNetProfitToRevenue
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashAssets(ctx context.Context, currencyunit, companyname, finreporttype string, reportYear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		COALESCE(SUM("ShortAssets" + "LongAssets" + "ShortAssetsSale"), 0) AS total_assets_sum
	FROM 
		dmart.investments_dash
	WHERE 
		"currencyunit" = $1 AND
		"name_short_en" = $2 AND
		"ProductionUnit" = 'tons' AND
		"report_type" = $3 AND
		"report_year" BETWEEN 2010 and $4
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, currencyunit, companyname, finreporttype, reportYear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]float64)
	for rows.Next() {
		var year int
		var totalAssetsSum float64
		if err := rows.Scan(&year, &totalAssetsSum); err != nil {
			return nil, err
		}
		result[year] = totalAssetsSum
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashCapital(ctx context.Context, currencyunit, companyname, finreporttype string, reportYear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		COALESCE(SUM("Capital"), 0) AS total_taxes
	FROM 
		dmart.investments_dash
	WHERE 
		"currencyunit" = $1 AND
		"name_short_en" = $2 AND
		"ProductionUnit" = 'tons' AND
		"report_type" = $3 AND
		"report_year" <= $4
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, currencyunit, companyname, finreporttype, reportYear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make(map[int]float64)
	for rows.Next() {
		var year int
		var totalTaxes sql.NullFloat64
		if err := rows.Scan(&year, &totalTaxes); err != nil {
			return nil, err
		}
		if totalTaxes.Valid {
			results[year] = totalTaxes.Float64
		} else {
			results[year] = 0 // or handle as needed
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashLiabilities(ctx context.Context, currencyunit, companyname, finreporttype string, reportYear int) (map[int]float64, error) {
	query := `
	SELECT 
		"report_year",
		COALESCE(SUM("LongLiabilities" + "ShortLiabilities"), 0) AS total_liabilities_sum
	FROM 
		dmart.investments_dash
	WHERE 
		"currencyunit" = $1 AND
		"name_short_en" = $2 AND
		"ProductionUnit" = 'tons' AND
		"report_type" = $3 AND
		"report_year" <= $4
	GROUP BY 
		"report_year"
	ORDER BY 
		"report_year";
	`

	rows, err := r.Db.QueryContext(ctx, query, currencyunit, companyname, finreporttype, reportYear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]float64)
	for rows.Next() {
		var year int
		var totalLiabilitiesSum float64
		if err := rows.Scan(&year, &totalLiabilitiesSum); err != nil {
			return nil, err
		}
		result[year] = totalLiabilitiesSum
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashSpecificNetProfitGraph(ctx context.Context, currencyunit, productionunit, reporttype string, reportYear int) (map[string]float64, error) {
	query := `
    SELECT 
        "name_short_en",
        COALESCE(AVG(NULLIF("NetProfit", 0) / NULLIF("Production", 0)), 0) AS avg_net_profit_per_production
    FROM 
        dmart.investments_dash
    WHERE 
        "currencyunit" = $1 AND
        "ProductionUnit" = $2 AND
		"report_type" = $3 AND
        "report_year" = $4
        
    GROUP BY 
        "name_short_en";
    `

	rows, err := r.Db.QueryContext(ctx, query, currencyunit, productionunit, reporttype, reportYear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]float64)
	for rows.Next() {
		var nameShortRu string
		var avgNetProfitPerProduction float64
		if err := rows.Scan(&nameShortRu, &avgNetProfitPerProduction); err != nil {
			return nil, err
		}
		result[nameShortRu] = avgNetProfitPerProduction
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *InvestmentsDashRepository) GetInvestmentsDashROAGraph(ctx context.Context, reporttype string, reportYear int) (map[string]float64, error) {
	query := `
    SELECT 
        "name_short_en",
        COALESCE(AVG(NULLIF("NetProfit", 0) / NULLIF(("ShortAssets" + "LongAssets" + "ShortAssetsSale"), 0)), 0) AS avg_roa
    FROM 
        dmart.investments_dash
    WHERE 
        "report_type" = $1 AND
		"report_year" = $2
    GROUP BY 
        "name_short_en";
    `

	rows, err := r.Db.QueryContext(ctx, query, reporttype, reportYear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]float64)
	for rows.Next() {
		var nameShortRu string
		var avgROA float64
		if err := rows.Scan(&nameShortRu, &avgROA); err != nil {
			return nil, err
		}
		result[nameShortRu] = avgROA
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
