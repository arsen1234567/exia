package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()

	// Define routes using pat
	mux.Get("/get-info", http.HandlerFunc(app.investment_oil_production_handler.GetInvestmentOilProductionSummary))

	mux.Get("/api/gas/review/production", http.HandlerFunc(app.gas_review_handler.GetGasProductionSummary))
	// Add more routes here as needed

	return mux
}
