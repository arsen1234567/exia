package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	publicServies "tender/internal/services/public"
)

type ReservesGasHandler struct {
	ReservesGasService publicServies.ReservesGasNgsService
}

func (h *ReservesGasHandler) GetDeposit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := h.ReservesGasService.GetDeposit(ctx)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (h *ReservesGasHandler) GetNumberOfCompanies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := h.ReservesGasService.GetNumberOfCompanies(ctx)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (h *ReservesGasHandler) GetTotalReserves(w http.ResponseWriter, r *http.Request) {
	oilType := r.URL.Query().Get("type")
	if oilType == "" {
		http.Error(w, "Missing 'type' query parameter", http.StatusBadRequest)
		return
	}

	results, err := h.ReservesGasService.GetTotalReserves(context.Background(), oilType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ReservesGasHandler) GetProduction(w http.ResponseWriter, r *http.Request) {
	oilType := r.URL.Query().Get("type")
	if oilType == "" {
		http.Error(w, "Missing 'type' query parameter", http.StatusBadRequest)
		return
	}

	results, err := h.ReservesGasService.GetProduction(context.Background(), oilType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ReservesGasHandler) GetNumberOfDepositsByRegion(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	if yearStr == "" {
		http.Error(w, "Missing 'year' query parameter", http.StatusBadRequest)
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Invalid 'year' query parameter", http.StatusBadRequest)
		return
	}

	result, err := h.ReservesGasService.GetNumberOfDepositsByRegion(context.Background(), year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *ReservesGasHandler) GetTopCompaniesByReserves(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	oilType := r.URL.Query().Get("oilType")

	if yearStr == "" || oilType == "" {
		http.Error(w, "Missing year or oilType parameter", http.StatusBadRequest)
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	reserves, err := h.ReservesGasService.GetTopCompaniesByReserves(r.Context(), year, oilType)
	if err != nil {
		http.Error(w, "Failed to get reserves data", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reserves)
}
