package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	dmartServices "tender/internal/services/dmart"
)

type OilperformanceResultsHandler struct {
	InvestmentsDashService *dmartServices.InvestmentsDashService
}

func (h *OilperformanceResultsHandler) GetInvestmentsDash(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	reportType := r.URL.Query().Get("reportType")
	reportYear := r.URL.Query().Get("year")

	// Parse the report year
	year, err := strconv.Atoi(reportYear)
	if err != nil {
		http.Error(w, "Invalid year", http.StatusBadRequest)
		return
	}

	// Validate and convert the reportType from boolean representation
	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	// Fetch the data from the service
	data, err := h.InvestmentsDashService.GetInvestmentsDash(ctx, company, reportTypeString, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashOilProduction(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	productionunit := r.URL.Query().Get("unit")
	reportType := r.URL.Query().Get("reportType")
	reportyear_str := r.URL.Query().Get("year")

	reportyear, err := strconv.Atoi(reportyear_str)
	if err != nil {
		http.Error(w, "Invalid year parameter.", http.StatusBadRequest)
		return
	}

	if productionunit != "barrels" && productionunit != "tons" && productionunit != "bpd" {
		http.Error(w, "Invalid unit. Only 'barrels', 'pbd' and 'tons' are allowed.", http.StatusBadRequest)
		return
	}

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	summary, err := h.InvestmentsDashService.GetInvestmentsDashOilProduction(ctx, company, productionunit, reportTypeString, reportyear)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashSpecificRevenue(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	currencyunit := r.URL.Query().Get("currency")
	productionunit := r.URL.Query().Get("unit")
	reportType := r.URL.Query().Get("reportType")
	reportyear_str := r.URL.Query().Get("year")

	reportyear, err := strconv.Atoi(reportyear_str)
	if err != nil {
		http.Error(w, "Invalid year parameter.", http.StatusBadRequest)
		return
	}

	if currencyunit != "KZT" && currencyunit != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	summary, err := h.InvestmentsDashService.GetInvestmentsDashSpecificRevenue(ctx, currencyunit, company, productionunit, reportTypeString, reportyear)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashROA(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	reportType := r.URL.Query().Get("reportType")
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

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	summary, err := h.InvestmentsDashService.GetInvestmentsDashROA(ctx, company, reportTypeString, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashNetProfitMargin(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	reportType := r.URL.Query().Get("reportType")
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

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	summary, err := h.InvestmentsDashService.GetInvestmentsDashNetProfitMargin(ctx, company, reportTypeString, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashSpecificNetProfit(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	currencyunit := r.URL.Query().Get("currency")
	productionunit := r.URL.Query().Get("unit")
	reportType := r.URL.Query().Get("reportType")
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

	if currencyunit != "KZT" && currencyunit != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	summary, err := h.InvestmentsDashService.GetInvestmentsDashSpecificNetProfit(ctx, currencyunit, company, productionunit, reportTypeString, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashRevenue(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	currencyunit := r.URL.Query().Get("currency")
	reportType := r.URL.Query().Get("reportType")
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

	if currencyunit != "KZT" && currencyunit != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	revenueSummary, err := h.InvestmentsDashService.GetInvestmentsDashRevenue(ctx, currencyunit, company, reportTypeString, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(revenueSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashOperatingProfit(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	currencyunit := r.URL.Query().Get("currency")

	reportType := r.URL.Query().Get("reportType")
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

	if currencyunit != "KZT" && currencyunit != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	operatingProfitSummary, err := h.InvestmentsDashService.GetInvestmentsDashOperatingProfit(ctx, currencyunit, company, reportTypeString, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(operatingProfitSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashEBITDA(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	currencyunit := r.URL.Query().Get("currency")
	reportType := r.URL.Query().Get("reportType")
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

	if currencyunit != "KZT" && currencyunit != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	totalSumSummary, err := h.InvestmentsDashService.GetInvestmentsDashEBITDA(ctx, currencyunit, company, reportTypeString, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(totalSumSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashNetProfit(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	currencyunit := r.URL.Query().Get("currency")
	reportType := r.URL.Query().Get("reportType")
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

	if currencyunit != "KZT" && currencyunit != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	totalSumSummary, err := h.InvestmentsDashService.GetInvestmentsDashNetProfit(ctx, currencyunit, company, reportTypeString, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(totalSumSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashTotalTaxes(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	currencyunit := r.URL.Query().Get("currency")
	reportType := r.URL.Query().Get("reportType")
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

	if currencyunit != "KZT" && currencyunit != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	totalSumSummary, err := h.InvestmentsDashService.GetInvestmentsDashTotalTaxes(ctx, currencyunit, company, reportTypeString, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(totalSumSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashTaxBurden(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	currencyunit := r.URL.Query().Get("currency")
	reportType := r.URL.Query().Get("reportType")
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

	if currencyunit != "KZT" && currencyunit != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	totalSumSummary, err := h.InvestmentsDashService.GetInvestmentsDashTaxBurden(ctx, currencyunit, company, reportTypeString, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(totalSumSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashSpecificTaxes(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	currencyunit := r.URL.Query().Get("currency")
	productionunit := r.URL.Query().Get("unit")
	reportType := r.URL.Query().Get("reportType")
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

	if currencyunit != "KZT" && currencyunit != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	totalSumSummary, err := h.InvestmentsDashService.GetInvestmentsDashSpecificTaxes(ctx, currencyunit, company, productionunit, reportTypeString, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(totalSumSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashAssets(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	currencyunit := r.URL.Query().Get("currency")
	reportType := r.URL.Query().Get("reportType")
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

	if currencyunit != "KZT" && currencyunit != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	totalSumSummary, err := h.InvestmentsDashService.GetInvestmentsDashAssets(ctx, currencyunit, company, reportTypeString, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(totalSumSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashCapital(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	currencyunit := r.URL.Query().Get("currency")
	reportType := r.URL.Query().Get("reportType")
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

	if currencyunit != "KZT" && currencyunit != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	totalSumSummary, err := h.InvestmentsDashService.GetInvestmentsDashCapital(ctx, currencyunit, company, reportTypeString, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(totalSumSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OilperformanceResultsHandler) GetInvestmentsDashLiabilities(w http.ResponseWriter, r *http.Request) {
	company := r.URL.Query().Get("company")
	currencyunit := r.URL.Query().Get("currency")
	reportType := r.URL.Query().Get("reportType")
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

	if currencyunit != "KZT" && currencyunit != "USD" {
		http.Error(w, "Invalid currency unit. Only 'KZT' and 'USD' are allowed.", http.StatusBadRequest)
		return
	}

	var reportTypeString string
	switch reportType {
	case "1":
		reportTypeString = "Консолидированный"
	case "0":
		reportTypeString = "Не консолидированный"
	default:
		http.Error(w, "Invalid report type. Only '1' (Консолидированный) and '0' (Не консолидированный) are allowed.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	totalSumSummary, err := h.InvestmentsDashService.GetInvestmentsDashLiabilities(ctx, currencyunit, company, reportTypeString, reportyear)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(totalSumSummary); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
