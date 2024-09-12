package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	publicServies "tender/internal/services/public"
)

type ReservesOilHandler struct {
	OilReservesService publicServies.ReservesOilNgsService
}

func (h *ReservesOilHandler) GetDeposit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := h.OilReservesService.GetDeposit(ctx)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (h *ReservesOilHandler) GetNumberOfCompanies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := h.OilReservesService.GetNumberOfCompanies(ctx)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (h *ReservesOilHandler) GetReservesOilNgsTotalProduction(w http.ResponseWriter, r *http.Request) {
	oilType := r.URL.Query().Get("type")
	if oilType == "" {
		http.Error(w, "Missing 'type' query parameter", http.StatusBadRequest)
		return
	}

	var oilTypeString string
	switch oilType {
	case "geological":
		oilTypeString = "Геологические"
	case "recoverable":
		oilTypeString = "Извлекаемые"
	default:
		http.Error(w, "Invalid report type. Only 'geological' (Геологические) and 'recoverable' (Извлекаемые) are allowed.", http.StatusBadRequest)
		return
	}

	result, err := h.OilReservesService.GetTotalProduction(context.Background(), oilTypeString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *ReservesOilHandler) GetNumberOfDepositsByRegion(w http.ResponseWriter, r *http.Request) {
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

	result, err := h.OilReservesService.GetNumberOfDepositsByRegion(context.Background(), year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *ReservesOilHandler) GetTopCompaniesByReserves(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	oilType := r.URL.Query().Get("oilType")

	if yearStr == "" || oilType == "" {
		http.Error(w, "Missing year or oilType parameter", http.StatusBadRequest)
		return
	}

	var oilTypeString string
	switch oilType {
	case "geological":
		oilTypeString = "Геологические"
	case "recoverable":
		oilTypeString = "Извлекаемые"
	default:
		http.Error(w, "Invalid report type. Only 'geological' (Геологические) and 'recoverable' (Извлекаемые) are allowed.", http.StatusBadRequest)
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	reserves, err := h.OilReservesService.GetTopCompaniesByReserves(r.Context(), year, oilTypeString)
	if err != nil {
		http.Error(w, "Failed to get reserves data", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reserves)
}
