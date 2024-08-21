package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()

	mux.Get("/api/oil/review/production", http.HandlerFunc(app.oil_review_handler.GetInvestmentOilProductionSummary)) // GetOilProduction
	mux.Get("/api/oil/review/reserves", http.HandlerFunc(app.oil_review_handler.GetInvestmentReservesSummary))        // GetOilReserves
	mux.Get("/api/oil/review/taxes", http.HandlerFunc(app.oil_review_handler.GetKgdTaxesProd))                        // GetOilTaxes
	mux.Get("/api/oil/review/investment_profit", http.HandlerFunc(app.oil_review_handler.GetInvestmentNetProfitSummary))

	mux.Get("/api/oil/forecast/reserves", http.HandlerFunc(app.oil_review_handler.GetInvestPotentialMainHandler))
	mux.Get("/api/oil/forecast/investment-review-steps", http.HandlerFunc(app.oil_review_handler.GetInvestmentReviewForecastStepsSummary)) // Investment_Review_Forecast_Steps
	mux.Get("/api/oil/forecast/investment-review-steps-companies", http.HandlerFunc(app.oil_review_handler.GetCompaniesForecastSteps))
	mux.Get("/api/oil/forecast/investment-review-steps-ebitda", http.HandlerFunc(app.oil_review_handler.GetEbitdaToGrossRevenueRatio)) // Investment_Review_Forecast_Steps
	mux.Get("/api/oil/forecast/investment-review-total", http.HandlerFunc(app.oil_review_handler.GetInvestmentReviewForecastTotal))    //  Investment_Review_Forecast_Total
	mux.Get("/api/oil/forecast/specopex", http.HandlerFunc(app.oil_review_handler.GetSpecOpEx))

	mux.Get("/api/oil/perfomance/investment-dash", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDash))
	mux.Get("/api/oil/perfomance/oil-production", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashOilProduction))
	mux.Get("/api/oil/perfomance/specific-revenue", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashSpecificRevenue))
	mux.Get("/api/oil/perfomance/roa", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashROA))
	mux.Get("/api/oil/perfomance/net-profit-margin", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashNetProfitMargin))
	mux.Get("/api/oil/perfomance/specific-net-profit", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashSpecificNetProfit))
	mux.Get("/api/oil/perfomance/revenue", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashRevenue))
	mux.Get("/api/oil/perfomance/operating-profit", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashOperatingProfit))
	mux.Get("/api/oil/perfomance/ebitda", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashEBITDA))
	mux.Get("/api/oil/perfomance/clear-profit", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashNetProfit))
	mux.Get("/api/oil/perfomance/total-taxes", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashTotalTaxes))
	mux.Get("/api/oil/perfomance/tax-burden", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashTaxBurden))
	mux.Get("/api/oil/perfomance/specific-taxes", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashSpecificTaxes))
	mux.Get("/api/oil/perfomance/assets", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashAssets))
	mux.Get("/api/oil/perfomance/capital", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashCapital))
	mux.Get("/api/oil/perfomance/liabilities", http.HandlerFunc(app.oil_perfomance_results_handler.GetInvestmentsDashLiabilities))

	mux.Get("/api/oil/benchmarking/net-profit-graph", http.HandlerFunc(app.oil_benchmarking_handler.GetInvestmentsDashSpecificNetProfitGraph))
	mux.Get("/api/oil/benchmarking/roa-graph", http.HandlerFunc(app.oil_benchmarking_handler.GetInvestmentsDashROAGraph))
	mux.Get("/api/oil/benchmarking/specific-taxes-graph", http.HandlerFunc(app.oil_benchmarking_handler.GetSpecificTaxesGraph)) // +usd
	mux.Get("/api/oil/benchmarking/tax-burden-graph", http.HandlerFunc(app.oil_benchmarking_handler.GetTaxBurdenGraph))
	mux.Get("/api/oil/benchmarking/summa-taxes-graph", http.HandlerFunc(app.oil_benchmarking_handler.GetSummaAllTaxes))

	mux.Get("/api/gas/review/production", http.HandlerFunc(app.gas_review_handler.GetGasProductionSummary))
	mux.Get("/api/gas/review/taxes", http.HandlerFunc(app.gas_review_handler.GetKgdTaxesSummary))
	mux.Get("/api/gas/review/recoverable_gas_reserves", http.HandlerFunc(app.gas_review_handler.GetRecoverableGasReservesSummary))
	mux.Get("/api/gas/review/net_profit", http.HandlerFunc(app.gas_review_handler.GetNetProfitSummary))

	mux.Get("/api/gas/review/forecast/reserve_ratio", http.HandlerFunc(app.gas_review_handler.GetReserveRatio))
	mux.Get("/api/gas/review/forecast/taxes", http.HandlerFunc(app.gas_review_handler.GetAmountOfPredictedTaxes))
	mux.Get("/api/gas/review/forecast/NPV+TV", http.HandlerFunc(app.gas_review_handler.GetNPVplusTV))
	mux.Get("/api/gas/review/forecast/EBITDAmargin", http.HandlerFunc(app.gas_review_handler.GetEBITDAmargin))
	mux.Get("/api/gas/review/forecast/gas_balance", http.HandlerFunc(app.gas_review_handler.GetGasBalance))

	mux.Get("/api/gas/perfomance/service_revenue", http.HandlerFunc(app.gas_perfomance_results_handler.GetRevenueByServiceAndCompanyAndYear))
	mux.Get("/api/gas/perfomance/geography_revenue", http.HandlerFunc(app.gas_perfomance_results_handler.GetRevenueByGeographyAndCompanyAndYear))
	mux.Get("/api/gas/perfomance/cost_items", http.HandlerFunc(app.gas_perfomance_results_handler.GetCostItemsByCompanyAndYear))
	mux.Get("/api/gas/perfomance/revenue", http.HandlerFunc(app.gas_perfomance_results_handler.GetRevenueByCompanyAndYear))
	mux.Get("/api/gas/perfomance/cost", http.HandlerFunc(app.gas_perfomance_results_handler.GetCostOfGoodsWorksServicesSold))
	mux.Get("/api/gas/perfomance/gross_profit", http.HandlerFunc(app.gas_perfomance_results_handler.GetGrossProfit))
	mux.Get("/api/gas/perfomance/CIT", http.HandlerFunc(app.gas_perfomance_results_handler.GetCIT))

	return mux
}
