package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"tender/internal/services"
)

type InvestmentOilProductionHandler struct {
	Service *services.InvestmentOilProductionService
}

func (h *InvestmentOilProductionHandler) GetInvestmentOilProductionSummary(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	summary, err := h.Service.GetInvestmentOilProductionSummary(ctx, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}
