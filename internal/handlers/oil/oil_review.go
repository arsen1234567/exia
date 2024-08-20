package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	dmartServices "tender/internal/services/dmart"
	prodServices "tender/internal/services/prod"
)

type OilReviewHandler struct {
	InvestmentOilProductionService       *dmartServices.InvestmentOilProductionService
	InvestmentReservesService            *dmartServices.InvestmentReservesService
	KgdTaxesProdService                  *prodServices.KgdTaxesProdService
	InvestmentNetProfitService           *dmartServices.InvestmentNetProfitService
	InvestPotentialMainService           *dmartServices.InvestPotentialMainService
	InvestmentReviewForecastStepsService *dmartServices.InvestmentReviewForecastStepsService
	InvestmentReviewForecastTotalService *dmartServices.InvestmentReviewForecastTotalService
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

func (h *OilReviewHandler) GetInvestmentNetProfitSummary(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	summary, err := h.InvestmentNetProfitService.GetInvestmentNetProfitSummary(ctx, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilReviewHandler) GetInvestPotentialMainHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get the data from the repository
	totalReserveMultiple, err := h.InvestPotentialMainService.GetInvestPotentialMain(ctx)
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
		log.Println("Error retrieving data:", err)
		return
	}

	// Set the content type to plain text
	w.Header().Set("Content-Type", "text/plain")

	// Convert float64 to string without scientific notation
	output := fmt.Sprintf("%.2f", totalReserveMultiple) // Adjust the precision if needed
	if _, err := w.Write([]byte(output)); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		log.Println("Error writing response:", err)
	}
}

func (h *OilReviewHandler) GetInvestmentReviewForecastStepsSummary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	totalReserveMultiple, err := h.InvestmentReviewForecastStepsService.GetInvestmentReviewForecastSteps(ctx)
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
		log.Println("Error retrieving data:", err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	output := fmt.Sprintf("%.2f", totalReserveMultiple)
	if _, err := w.Write([]byte(output)); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		log.Println("Error writing response:", err)
	}
}

func (h *OilReviewHandler) GetEbitdaToGrossRevenueRatio(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	totalReserveMultiple, err := h.InvestmentReviewForecastStepsService.GetEbitdaToGrossRevenueRatio(ctx)
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
		log.Println("Error retrieving data:", err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	output := fmt.Sprintf("%f", totalReserveMultiple)
	if _, err := w.Write([]byte(output)); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		log.Println("Error writing response:", err)
	}
}

func (h *OilReviewHandler) GetInvestmentReviewForecastTotal(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	totalReserveMultiple, err := h.InvestmentReviewForecastTotalService.GetInvestmentReviewForecastTotal(ctx)
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
		log.Println("Error retrieving data:", err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	output := fmt.Sprintf("%f", totalReserveMultiple)
	if _, err := w.Write([]byte(output)); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		log.Println("Error writing response:", err)
	}
}

func (h *OilReviewHandler) GetSpecOpEx(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	totalReserveMultiple, err := h.InvestPotentialMainService.GetSpecOpEx(ctx)
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
		log.Println("Error retrieving data:", err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	output := fmt.Sprintf("%f", totalReserveMultiple)
	if _, err := w.Write([]byte(output)); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		log.Println("Error writing response:", err)
	}
}

func (h *OilReviewHandler) GetCompaniesForecastSteps(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	oilProductions, err := h.InvestmentReviewForecastStepsService.GetCompaniesForecastStepsSummary(ctx)
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
		log.Println("Error retrieving data:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(oilProductions); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
