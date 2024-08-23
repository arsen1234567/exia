package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	dmartServices "tender/internal/services/dmart"
	prodServices "tender/internal/services/prod"
)

type OilBenchmarkingHandler struct {
	InvestmentsDashService *dmartServices.InvestmentsDashService
	SpecificTaxesService   *prodServices.SpecificTaxesService
	TaxBurdenService       *prodServices.TaxBurdenService
	KgdTaxesProdService    *prodServices.KgdTaxesProdService
}

func (h *OilBenchmarkingHandler) GetInvestmentsDashSpecificNetProfitGraph(w http.ResponseWriter, r *http.Request) {
	currencyunit := r.URL.Query().Get("currency")
	productionunit := r.URL.Query().Get("unit")
	reportyear_str := r.URL.Query().Get("year")
	reporttype := r.URL.Query().Get("reportType")

	if reporttype == "" {
		http.Error(w, "Report type is required.", http.StatusBadRequest)
		return
	}

	if reportyear_str == "" {
		http.Error(w, "Year parameter is required.", http.StatusBadRequest)
		return
	}

	reportyear, err := strconv.Atoi(reportyear_str)
	if err != nil {
		http.Error(w, "Invalid year parameter. Must be an integer.", http.StatusBadRequest)
		return
	}

	if currencyunit != "KZT" && currencyunit != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	totalSumSummary, err := h.InvestmentsDashService.GetInvestmentsDashSpecificNetProfitGraph(ctx, currencyunit, productionunit, reporttype, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(totalSumSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilBenchmarkingHandler) GetInvestmentsDashROAGraph(w http.ResponseWriter, r *http.Request) {
	reporttype := r.URL.Query().Get("reportType")
	reportyear_str := r.URL.Query().Get("year")

	if reportyear_str == "" {
		http.Error(w, "Year parameter is required.", http.StatusBadRequest)
		return
	}

	reportyear, err := strconv.Atoi(reportyear_str)
	if err != nil {
		http.Error(w, "Invalid year parameter. Must be an integer.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	totalSumSummary, err := h.InvestmentsDashService.GetInvestmentsDashROAGraph(ctx, reporttype, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(totalSumSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilBenchmarkingHandler) GetSpecificTaxesGraph(w http.ResponseWriter, r *http.Request) {

	reportyear_str := r.URL.Query().Get("year")
	currencyunit := r.URL.Query().Get("currency")
	reporttype := r.URL.Query().Get("reportType")

	if reportyear_str == "" {
		http.Error(w, "Year parameter is required.", http.StatusBadRequest)
		return
	}

	reportyear, err := strconv.Atoi(reportyear_str)
	if err != nil {
		http.Error(w, "Invalid year parameter. Must be an integer.", http.StatusBadRequest)
		return
	}

	if currencyunit != "KZT" && currencyunit != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	if reporttype != "Консолидированный" && reporttype != "Не консолидированный" {
		http.Error(w, "Invalid report type. Only 'Консолидированный' and 'Не консолидированный' are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	totalSumSummary, err := h.SpecificTaxesService.GetSpecificTaxes(ctx, reportyear, currencyunit, reporttype)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(totalSumSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilBenchmarkingHandler) GetTaxBurdenGraph(w http.ResponseWriter, r *http.Request) {

	year_str := r.URL.Query().Get("year")
	currency := r.URL.Query().Get("currency")

	if year_str == "" {
		http.Error(w, "Year parameter is required.", http.StatusBadRequest)
		return
	}

	year, err := strconv.Atoi(year_str)
	if err != nil {
		http.Error(w, "Invalid year parameter. Must be an integer.", http.StatusBadRequest)
		return
	}

	if currency != "KZT" && currency != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	totalSumSummary, err := h.TaxBurdenService.GetTaxBurden(ctx, year, currency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(totalSumSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilBenchmarkingHandler) GetSummaAllTaxes(w http.ResponseWriter, r *http.Request) {

	year_str := r.URL.Query().Get("year")
	currency := r.URL.Query().Get("currency")
	reporttype := r.URL.Query().Get("reportType")

	if year_str == "" {
		http.Error(w, "Year parameter is required.", http.StatusBadRequest)
		return
	}

	year, err := strconv.Atoi(year_str)
	if err != nil {
		http.Error(w, "Invalid year parameter. Must be an integer.", http.StatusBadRequest)
		return
	}

	if currency != "KZT" && currency != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	if reporttype != "Консолидированный" && reporttype != "Не консолидированный" {
		http.Error(w, "Invalid report type. Only 'Консолидированный' and 'Не консолидированный' are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	totalSumSummary, err := h.KgdTaxesProdService.GetSummaAllTaxes(ctx, year, currency, reporttype)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(totalSumSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
