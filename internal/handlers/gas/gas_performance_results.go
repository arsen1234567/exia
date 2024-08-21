package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	dmartServices "tender/internal/services/dmart"
)

type GasperformanceResultsHandler struct {
	DfoQazaqgasService   *dmartServices.DfoQazaqgasService
	DfoGgReportesService *dmartServices.DfoGgReportesService
}

func (h *GasperformanceResultsHandler) GetRevenueByServiceAndCompanyAndYear(w http.ResponseWriter, r *http.Request) {
	// Extract and validate the year parameter from the query string
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	// Extract the company parameter from the query string
	company := r.URL.Query().Get("company")
	if company == "" {
		http.Error(w, "Company parameter is required", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the service revenue by company and year
	summary, err := h.DfoQazaqgasService.GetRevenueByServicesAndCompanyAndYear(ctx, company, year)
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

func (h *GasperformanceResultsHandler) GetRevenueByGeographyAndCompanyAndYear(w http.ResponseWriter, r *http.Request) {
	// Extract and validate the year parameter from the query string
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	// Extract the company parameter from the query string
	company := r.URL.Query().Get("company")
	if company == "" {
		http.Error(w, "Company parameter is required", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the revenue by geography for the company and year
	summary, err := h.DfoQazaqgasService.GetRevenueByGeographyAndCompanyAndYear(ctx, company, year)
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

func (h *GasperformanceResultsHandler) GetCostItemsByCompanyAndYear(w http.ResponseWriter, r *http.Request) {
	// Extract and validate the year parameter from the query string
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	// Extract the company parameter from the query string
	company := r.URL.Query().Get("company")
	if company == "" {
		http.Error(w, "Company parameter is required", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the cost items by company and year
	summary, err := h.DfoQazaqgasService.GetCostItemsByCompanyAndYear(ctx, company, year)
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

func (h *GasperformanceResultsHandler) GetRevenueByCompanyAndYear(w http.ResponseWriter, r *http.Request) {
	// Extract and validate the year parameter from the query string
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	// Extract the company parameter from the query string
	company := r.URL.Query().Get("company")
	if company == "" {
		http.Error(w, "Company parameter is required", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the revenue by company and year
	totalRevenue, err := h.DfoGgReportesService.GetRevenueByCompanyAndYear(ctx, company, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the raw value to the response
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.FormatFloat(totalRevenue, 'f', -1, 64)))
}

func (h *GasperformanceResultsHandler) GetCostOfGoodsWorksServicesSold(w http.ResponseWriter, r *http.Request) {
	// Extract and validate the year parameter from the query string
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	// Extract the company parameter from the query string
	company := r.URL.Query().Get("company")
	if company == "" {
		http.Error(w, "Company parameter is required", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the cost of goods, works, and services sold by company and year
	totalCost, err := h.DfoGgReportesService.GetCostOfGoodsWorksServicesSold(ctx, company, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the raw value to the response
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.FormatFloat(totalCost, 'f', -1, 64)))
}

func (h *GasperformanceResultsHandler) GetGrossProfit(w http.ResponseWriter, r *http.Request) {
	// Extract and validate the year parameter from the query string
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	// Extract the company parameter from the query string
	company := r.URL.Query().Get("company")
	if company == "" {
		http.Error(w, "Company parameter is required", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the gross profit by company and year
	grossProfit, err := h.DfoGgReportesService.GetGrossProfit(ctx, company, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the raw value to the response
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.FormatFloat(grossProfit, 'f', -1, 64)))
}

func (h *GasperformanceResultsHandler) GetCIT(w http.ResponseWriter, r *http.Request) {
	// Extract and validate the year parameter from the query string
	yearStr := r.URL.Query().Get("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	// Extract the company parameter from the query string
	company := r.URL.Query().Get("company")
	if company == "" {
		http.Error(w, "Company parameter is required", http.StatusBadRequest)
		return
	}

	// Create a context for the request
	ctx := context.Background()

	// Call the service function to get the CIT by company and year
	totalCIT, err := h.DfoGgReportesService.GetCIT(ctx, company, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the raw value to the response
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.FormatFloat(totalCIT, 'f', -1, 64)))
}
