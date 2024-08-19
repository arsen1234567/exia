package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	dmartService "tender/internal/services/dmart"
	kgdService "tender/internal/services/rawData"
)

type OilReviewHandler struct {
	InvestmentOilProductionService *dmartService.InvestmentOilProductionService
	InvestmentReservesService      *dmartService.InvestmentReservesService
	KgdTaxesProdService            *kgdService.KgdTaxesService
}

func (h *OilReviewHandler) GetInvestmentOilProductionSummary(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	unit := r.URL.Query().Get("unit")
	if unit == "" {
		http.Error(w, "Unit parameter is required", http.StatusBadRequest)
		return
	}

	// Optionally: Validate unit value (e.g., ensure it's either 'barrels' or 'tons')
	if unit != "barrels" && unit != "tons" {
		http.Error(w, "Invalid unit parameter", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	summary, err := h.InvestmentOilProductionService.GetInvestmentOilProductionSummary(ctx, year, unit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilReviewHandler) GetKgdTaxesProd(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	summary, err := h.KgdTaxesProdService.GetKgdTaxesProdSummary(ctx, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilReviewHandler) GetInvestmentReservesSummary(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	unit := r.URL.Query().Get("unit")
	if unit == "" {
		http.Error(w, "Unit parameter is required", http.StatusBadRequest)
		return
	}

	// Optionally: Validate unit value (e.g., ensure it's either 'barrels' or 'tons')
	if unit != "barrels" && unit != "tons" {
		http.Error(w, "Invalid unit parameter", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	summary, err := h.InvestmentReservesService.GetInvestmentReservesSummary(ctx, year, unit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
