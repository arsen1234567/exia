package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	dmartServices "tender/internal/services/dmart"
	prodServices "tender/internal/services/prod"
	rawDataServices "tender/internal/services/rawData"
)

type GasReviewHandler struct {
	ProductionGasService  *prodServices.ProductionGasService
	GasStepsService       *prodServices.GasStepsService
	GasTotalsService      *prodServices.GasTotalsService
	KgdTaxesService       *rawDataServices.KgdTaxesService
	NgsReservesGasService *rawDataServices.NgsReservesGasService
	StGasBalanceService   *rawDataServices.StGasBalanceService
	DfoGgReportesService  *dmartServices.DfoGgReportesService
	NaturalGasMainService *dmartServices.NaturalGasMainService
}

func (h *GasReviewHandler) GetGasProductionSummary(w http.ResponseWriter, r *http.Request) {
	// Extract and validate the year parameter from the query string
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	// Extract the unit parameter from the query string
	unit := r.URL.Query().Get("unit")
	if unit != "m" && unit != "ft" {
		http.Error(w, "Invalid unit parameter. Must be 'm' or 'ft'", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the gas production summary
	summary, err := h.ProductionGasService.GetGasProductionSummary(ctx, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the unit is "ft", convert the total production to cubic feet
	if unit == "ft" {
		for i := range summary {
			summary[i].TotalProduction *= 35.3147 // Conversion factor from cubic meters to cubic feet
		}
	}

	// Set the Content-Type header and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *GasReviewHandler) GetKgdTaxesSummary(w http.ResponseWriter, r *http.Request) {
	// Extract and validate the year parameter from the query string
	yearStr := r.URL.Query().Get("year")
	currency := r.URL.Query().Get("currency")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	if currency != "USD" && currency != "KZT" {
		http.Error(w, "Invalid currency, nust be USD or KZT", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the KgdTaxes summary
	summary, err := h.KgdTaxesService.GetKgdTaxesSummary(ctx, year, currency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *GasReviewHandler) GetRecoverableGasReservesSummary(w http.ResponseWriter, r *http.Request) {

	// Extract and validate the end year parameter from the query string
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the recoverable gas reserves summary
	summary, err := h.NgsReservesGasService.GetRecoverableGasReservesSummary(ctx, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *GasReviewHandler) GetNetProfitSummary(w http.ResponseWriter, r *http.Request) {
	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the net profit summary
	summary, err := h.DfoGgReportesService.GetNetProfitSummary(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *GasReviewHandler) GetReserveRatio(w http.ResponseWriter, r *http.Request) {
	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the reserve ratio
	reserveRatio, err := h.NaturalGasMainService.GetReserveRatio(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reserveRatio); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *GasReviewHandler) GetAmountOfPredictedTaxes(w http.ResponseWriter, r *http.Request) {
	// Create a context for the request
	ctx := context.Background()

	currency := r.URL.Query().Get("currency")

	if currency != "USD" && currency != "KZT" {
		http.Error(w, "Invalid currency, nust be USD or KZT", http.StatusBadRequest)
		return
	}
	// Call the service function to get the amount of predicted taxes
	amountOfPredictedTaxes, err := h.GasStepsService.GetAmountOfPredictedTaxes(ctx, currency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(amountOfPredictedTaxes); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *GasReviewHandler) GetNPVplusTV(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	currency := r.URL.Query().Get("currency")

	if currency != "USD" && currency != "KZT" {
		http.Error(w, "Invalid currecy, must be USD or KZT", http.StatusBadRequest)
		return
	}

	totalNPVplusTV, err := h.GasTotalsService.GetNPVplusTV(ctx, currency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(totalNPVplusTV); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *GasReviewHandler) GetEBITDAmargin(w http.ResponseWriter, r *http.Request) {
	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the EBITDA margin
	ebitdaMargin, err := h.GasStepsService.GetEBITDAmargin(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ebitdaMargin); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *GasReviewHandler) GetGasBalance(w http.ResponseWriter, r *http.Request) {
	// Extract and validate the year parameter from the query string
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	// Extract the unit parameter from the query string
	unit := r.URL.Query().Get("unit")
	if unit != "m" && unit != "ft" {
		http.Error(w, "Invalid unit parameter. Must be 'm' or 'ft'", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the gas balance summary
	gasBalance, err := h.StGasBalanceService.GetGasBalance(ctx, int64(year), unit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(gasBalance); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
