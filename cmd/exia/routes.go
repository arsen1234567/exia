package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()

	mux.Get("/api/oil/review/production", http.HandlerFunc(app.investment_oil_production_handler.GetInvestmentOilProductionSummary))
	mux.Get("/api/oil/review/reserves", http.HandlerFunc(app.investment_reserves_handler.GetInvestmentReservesSummary))
	mux.Get("/api/oil/review/taxes", http.HandlerFunc(app.kgd_taxes_prod_handler.GetKgdTaxesProd))

	// mux.Get("/api/gas/review/production", http.HandlerFunc(app.gas_review_handler.GetGasProductionSummary))
	// mux.Get("/api/gas/review/taxes", http.HandlerFunc(app.gas_review_handler.GetKgdTaxesSummary))

	return mux
}
