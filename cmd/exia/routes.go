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

	mux.Get("/api/gas/review/production", http.HandlerFunc(app.gas_review_handler.GetGasProductionSummary))
	mux.Get("/api/gas/review/taxes", http.HandlerFunc(app.gas_review_handler.GetKgdTaxesSummary))
	mux.Get("/api/gas/review/recoverable_gas_reserves", http.HandlerFunc(app.gas_review_handler.GetRecoverableGasReservesSummary))

	return mux
}
