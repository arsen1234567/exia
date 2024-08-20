package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()

	mux.Get("/api/oil/review/production", http.HandlerFunc(app.oil_review_handler.GetInvestmentOilProductionSummary))
	mux.Get("/api/oil/review/reserves", http.HandlerFunc(app.oil_review_handler.GetInvestmentReservesSummary))
	mux.Get("/api/oil/review/taxes", http.HandlerFunc(app.oil_review_handler.GetKgdTaxesProd))

	mux.Get("/api/gas/review/production", http.HandlerFunc(app.gas_review_handler.GetGasProductionSummary))
	mux.Get("/api/gas/review/taxes", http.HandlerFunc(app.gas_review_handler.GetKgdTaxesSummary))
	mux.Get("/api/gas/review/recoverable_gas_reserves", http.HandlerFunc(app.gas_review_handler.GetRecoverableGasReservesSummary))
	mux.Get("/api/gas/review/net_profit", http.HandlerFunc(app.gas_review_handler.GetNetProfitSummary))
	mux.Get("/api/gas/review/forecast/reserve_ratio", http.HandlerFunc(app.gas_review_handler.GetReserveRatio))
	mux.Get("/api/gas/review/forecast/taxes", http.HandlerFunc(app.gas_review_handler.GetAmountOfPredictedTaxes))
	mux.Get("/api/gas/review/forecast/NPV+TV", http.HandlerFunc(app.gas_review_handler.GetNPVplusTV))
	mux.Get("/api/gas/review/forecast/EBITDAmargin", http.HandlerFunc(app.gas_review_handler.GetEBITDAmargin))
	mux.Get("/api/gas/review/forecast/gas_balance", http.HandlerFunc(app.gas_review_handler.GetGasBalance))

	return mux
}
