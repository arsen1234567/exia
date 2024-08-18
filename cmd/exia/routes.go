package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()

	mux.Get("/api/oil/review/get-oil-production", http.HandlerFunc(app.investment_oil_production_handler.GetInvestmentOilProductionSummary))
	mux.Get("/api/oil/review/get-reserves", http.HandlerFunc(app.investment_oil_production_handler.GetInvestmentOilProductionSummary))

	return mux
}
