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
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the KgdTaxes summary
	summary, err := h.KgdTaxesService.GetKgdTaxesSummary(ctx, year)
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
	// Extract and validate the start year parameter from the query string
	startYearStr := r.URL.Query().Get("start_year")
	startYear, err := strconv.Atoi(startYearStr)
	if err != nil || startYear <= 0 {
		http.Error(w, "Invalid start_year parameter", http.StatusBadRequest)
		return
	}

	// Extract and validate the end year parameter from the query string
	endYearStr := r.URL.Query().Get("end_year")
	endYear, err := strconv.Atoi(endYearStr)
	if err != nil || endYear <= 0 {
		http.Error(w, "Invalid end_year parameter", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the recoverable gas reserves summary
	summary, err := h.NgsReservesGasService.GetRecoverableGasReservesSummary(ctx, startYear, endYear)
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

	// Call the service function to get the amount of predicted taxes
	amountOfPredictedTaxes, err := h.GasStepsService.GetAmountOfPredictedTaxes(ctx)
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
	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the NPV plus Terminal Value
	totalNPVplusTV, err := h.GasTotalsService.GetNPVplusTV(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and encode the response as JSON
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
