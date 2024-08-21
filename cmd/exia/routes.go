package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()

	mux.Get("/api/oil/review/financial_indicators/production", http.HandlerFunc(app.oil_review_handler.GetInvestmentOilProductionSummary)) // GetOilProduction
	mux.Get("/api/oil/review/financial_indicators/reserves", http.HandlerFunc(app.oil_review_handler.GetInvestmentReservesSummary))        // GetOilReserves
	mux.Get("/api/oil/review/financial_indicators/taxes", http.HandlerFunc(app.oil_review_handler.GetKgdTaxesProd))                        // GetOilTaxes
	mux.Get("/api/oil/review/financial_indicators/investment_profit", http.HandlerFunc(app.oil_review_handler.GetInvestmentNetProfitSummary))

	mux.Get("/api/oil/review/forecast/reserves", http.HandlerFunc(app.oil_review_handler.GetInvestPotentialMainHandler))
	mux.Get("/api/oil/review/forecast/investment-review-steps", http.HandlerFunc(app.oil_review_handler.GetInvestmentReviewForecastStepsSummary)) // Investment_Review_Forecast_Steps
	mux.Get("/api/oil/review/forecast/investment-review-steps-companies", http.HandlerFunc(app.oil_review_handler.GetCompaniesForecastSteps))
	mux.Get("/api/oil/review/forecast/EBITDAmargin", http.HandlerFunc(app.oil_review_handler.GetEbitdaToGrossRevenueRatio)) // Investment_Review_Forecast_Steps
	mux.Get("/api/oil/review/forecast/taxes", http.HandlerFunc(app.oil_review_handler.GetInvestmentReviewForecastTotal))    //  Investment_Review_Forecast_Total
	mux.Get("/api/oil/review/forecast/specopex", http.HandlerFunc(app.oil_review_handler.GetSpecOpEx))

	mux.Get("/api/oil/performance/investment-dash", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDash))
	mux.Get("/api/oil/performance/oil-production", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashOilProduction))
	mux.Get("/api/oil/performance/specific-revenue", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashSpecificRevenue))
	mux.Get("/api/oil/performance/roa", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashROA))
	mux.Get("/api/oil/performance/net-profit-margin", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashNetProfitMargin))
	mux.Get("/api/oil/performance/specific-net-profit", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashSpecificNetProfit))
	mux.Get("/api/oil/performance/revenue", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashRevenue))
	mux.Get("/api/oil/performance/operating-profit", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashOperatingProfit))
	mux.Get("/api/oil/performance/ebitda", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashEBITDA))
	mux.Get("/api/oil/performance/clear-profit", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashNetProfit))
	mux.Get("/api/oil/performance/total-taxes", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashTotalTaxes))
	mux.Get("/api/oil/performance/tax-burden", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashTaxBurden))
	mux.Get("/api/oil/performance/specific-taxes", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashSpecificTaxes))
	mux.Get("/api/oil/performance/assets", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashAssets))
	mux.Get("/api/oil/performance/capital", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashCapital))
	mux.Get("/api/oil/performance/liabilities", http.HandlerFunc(app.oil_performance_results_handler.GetInvestmentsDashLiabilities))

	mux.Get("/api/oil/benchmarking/net-profit-graph", http.HandlerFunc(app.oil_benchmarking_handler.GetInvestmentsDashSpecificNetProfitGraph))
	mux.Get("/api/oil/benchmarking/roa-graph", http.HandlerFunc(app.oil_benchmarking_handler.GetInvestmentsDashROAGraph))
	mux.Get("/api/oil/benchmarking/specific-taxes-graph", http.HandlerFunc(app.oil_benchmarking_handler.GetSpecificTaxesGraph)) // +usd
	mux.Get("/api/oil/benchmarking/tax-burden-graph", http.HandlerFunc(app.oil_benchmarking_handler.GetTaxBurdenGraph))
	mux.Get("/api/oil/benchmarking/summa-taxes-graph", http.HandlerFunc(app.oil_benchmarking_handler.GetSummaAllTaxes))

	mux.Get("/api/gas/review/financial_indicators/production", http.HandlerFunc(app.gas_review_handler.GetGasProductionSummary))
	mux.Get("/api/gas/review/financial_indicators/taxes", http.HandlerFunc(app.gas_review_handler.GetKgdTaxesSummary))
	mux.Get("/api/gas/review/financial_indicators/recoverable_gas_reserves", http.HandlerFunc(app.gas_review_handler.GetRecoverableGasReservesSummary))
	mux.Get("/api/gas/review/financial_indicators/net_profit", http.HandlerFunc(app.gas_review_handler.GetNetProfitSummary))

	mux.Get("/api/gas/review/perspective/reserve_ratio", http.HandlerFunc(app.gas_review_handler.GetReserveRatio))
	mux.Get("/api/gas/review/perspective/taxes", http.HandlerFunc(app.gas_review_handler.GetAmountOfPredictedTaxes))
	mux.Get("/api/gas/review/perspective/NPV+TV", http.HandlerFunc(app.gas_review_handler.GetNPVplusTV))
	mux.Get("/api/gas/review/perspective/EBITDAmargin", http.HandlerFunc(app.gas_review_handler.GetEBITDAmargin))
	mux.Get("/api/gas/review/perspective/gas_balance", http.HandlerFunc(app.gas_review_handler.GetGasBalance))

	mux.Get("/api/gas/performance/service_revenue", http.HandlerFunc(app.gas_performance_results_handler.GetRevenueByServiceAndCompanyAndYear))
	mux.Get("/api/gas/performance/geography_revenue", http.HandlerFunc(app.gas_performance_results_handler.GetRevenueByGeographyAndCompanyAndYear))
	mux.Get("/api/gas/performance/cost_items", http.HandlerFunc(app.gas_performance_results_handler.GetCostItemsByCompanyAndYear))
	mux.Get("/api/gas/performance/revenue", http.HandlerFunc(app.gas_performance_results_handler.GetRevenueByCompanyAndYear))
	mux.Get("/api/gas/performance/cost", http.HandlerFunc(app.gas_performance_results_handler.GetCostOfGoodsWorksServicesSold))
	mux.Get("/api/gas/performance/gross_profit", http.HandlerFunc(app.gas_performance_results_handler.GetGrossProfit))
	mux.Get("/api/gas/performance/CIT", http.HandlerFunc(app.gas_performance_results_handler.GetCIT))

	return mux
}
