package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	publicServies "tender/internal/services/public"
)

type ReservesCondHandler struct {
	ReservesCondService publicServies.ReservesCondNgsService
}

func (h *ReservesCondHandler) GetDeposit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := h.ReservesCondService.GetDeposit(ctx)
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

func (h *ReservesCondHandler) GetNumberOfCompanies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := h.ReservesCondService.GetNumberOfCompanies(ctx)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (h *ReservesCondHandler) GetTotalReserves(w http.ResponseWriter, r *http.Request) {
	oilType := r.URL.Query().Get("type")
	if oilType == "" {
		http.Error(w, "Missing 'type' query parameter", http.StatusBadRequest)
		return
	}

	results, err := h.ReservesCondService.GetTotalReserves(context.Background(), oilType)
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

func (h *ReservesCondHandler) GetProduction(w http.ResponseWriter, r *http.Request) {
	oilType := r.URL.Query().Get("type")
	if oilType == "" {
		http.Error(w, "Missing 'type' query parameter", http.StatusBadRequest)
		return
	}

	results, err := h.ReservesCondService.GetProduction(context.Background(), oilType)
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

func (h *ReservesCondHandler) GetNumberOfDepositsByRegion(w http.ResponseWriter, r *http.Request) {
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

	result, err := h.ReservesCondService.GetNumberOfDepositsByRegion(context.Background(), year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *ReservesCondHandler) GetTopCompaniesByReserves(w http.ResponseWriter, r *http.Request) {
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

	reserves, err := h.ReservesCondService.GetTopCompaniesByReserves(r.Context(), year, oilType)
	if err != nil {
		http.Error(w, "Failed to get reserves data", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reserves)
}
