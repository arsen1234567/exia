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

	mux.Get("/api/gas/review/financial_indicators/production", http.HandlerFunc(app.gas_review_handler.GetGasProductionSummary))
	mux.Get("/api/gas/review/financial_indicators/taxes", http.HandlerFunc(app.gas_review_handler.GetKgdTaxesSummary))
	mux.Get("/api/gas/review/financial_indicators/recoverable_gas_reserves", http.HandlerFunc(app.gas_review_handler.GetRecoverableGasReservesSummary))
	mux.Get("/api/gas/review/financial_indicators/net_profit", http.HandlerFunc(app.gas_review_handler.GetNetProfitSummary))

	mux.Get("/api/gas/review/perspective/reserve_ratio", http.HandlerFunc(app.gas_review_handler.GetReserveRatio))
	mux.Get("/api/gas/review/perspective/taxes", http.HandlerFunc(app.gas_review_handler.GetAmountOfPredictedTaxes))
	mux.Get("/api/gas/review/perspective/NPV+TV", http.HandlerFunc(app.gas_review_handler.GetNPVplusTV))
	mux.Get("/api/gas/review/perspective/EBITDAmargin", http.HandlerFunc(app.gas_review_handler.GetEBITDAmargin))
	mux.Get("/api/gas/review/perspective/gas_balance", http.HandlerFunc(app.gas_review_handler.GetGasBalance))

	mux.Get("/api/gas/perfomance/service_revenue", http.HandlerFunc(app.gas_perfomance_results_handler.GetRevenueByServiceAndCompanyAndYear))
	mux.Get("/api/gas/perfomance/geography_revenue", http.HandlerFunc(app.gas_perfomance_results_handler.GetRevenueByGeographyAndCompanyAndYear))
	mux.Get("/api/gas/perfomance/cost_items", http.HandlerFunc(app.gas_perfomance_results_handler.GetCostItemsByCompanyAndYear))
	mux.Get("/api/gas/perfomance/revenue", http.HandlerFunc(app.gas_perfomance_results_handler.GetRevenueByCompanyAndYear))
	mux.Get("/api/gas/perfomance/cost", http.HandlerFunc(app.gas_perfomance_results_handler.GetCostOfGoodsWorksServicesSold))
	mux.Get("/api/gas/perfomance/gross_profit", http.HandlerFunc(app.gas_perfomance_results_handler.GetGrossProfit))
	mux.Get("/api/gas/perfomance/CIT", http.HandlerFunc(app.gas_perfomance_results_handler.GetCIT))

	return mux
}
