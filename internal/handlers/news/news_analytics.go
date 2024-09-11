package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	newsServices "tender/internal/services/news"
)

type NewAnalyticsHandler struct {
	GoogleNewsService *newsServices.GoogleNewsService
}

func (h *NewAnalyticsHandler) GetCompanyCount(w http.ResponseWriter, r *http.Request) {

	period := r.URL.Query().Get("period")
	if period == "" {
		http.Error(w, "Missing 'period' query parameter", http.StatusBadRequest)
		return
	}

	companies, err := h.GoogleNewsService.GetGoogleNews(context.Background(), period)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companies)
}

func (h *NewAnalyticsHandler) GetMediaCount(w http.ResponseWriter, r *http.Request) {

	period := r.URL.Query().Get("period")
	if period == "" {
		http.Error(w, "Missing 'period' query parameter", http.StatusBadRequest)
		return
	}

	companies, err := h.GoogleNewsService.GetMediaCount(context.Background(), period)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companies)
}

func (h *NewAnalyticsHandler) GetMediaNews(w http.ResponseWriter, r *http.Request) {

	period := r.URL.Query().Get("period")
	if period == "" {
		http.Error(w, "Missing 'period' query parameter", http.StatusBadRequest)
		return
	}

	companies, err := h.GoogleNewsService.GetMediaNews(context.Background(), period)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companies)
}

func (h *NewAnalyticsHandler) GetSentimentCounts(w http.ResponseWriter, r *http.Request) {
	period := r.URL.Query().Get("period")
	if period == "" {
		http.Error(w, "Missing 'period' query parameter", http.StatusBadRequest)
		return
	}

	sentimentCounts, err := h.GoogleNewsService.GetTonal(context.Background(), period)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sentimentCounts)
}

func (h *NewAnalyticsHandler) GetSentimentDayCounts(w http.ResponseWriter, r *http.Request) {
	period := r.URL.Query().Get("period")
	if period == "" {
		http.Error(w, "Missing 'period' query parameter", http.StatusBadRequest)
		return
	}

	sentimentCounts, err := h.GoogleNewsService.GetSentimentCountsByDay(context.Background(), period)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sentimentCounts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *NewAnalyticsHandler) GetSentimentMap(w http.ResponseWriter, r *http.Request) {
	period := r.URL.Query().Get("period")
	sentiment := r.URL.Query().Get("sentiment")
	if period == "" {
		http.Error(w, "Missing 'period' query parameter", http.StatusBadRequest)
		return
	}

	sentimentCounts, err := h.GoogleNewsService.GetSentimentMap(context.Background(), period, sentiment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sentimentCounts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *NewAnalyticsHandler) GetTopCompanyCount(w http.ResponseWriter, r *http.Request) {
	period := r.URL.Query().Get("period")

	if period == "" {
		http.Error(w, "Missing 'period' query parameter", http.StatusBadRequest)
		return
	}

	sentimentCounts, err := h.GoogleNewsService.GetTopCompanyCount(context.Background(), period)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sentimentCounts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *NewAnalyticsHandler) GetTopCompanyCountDict(w http.ResponseWriter, r *http.Request) {
	period := r.URL.Query().Get("period")

	if period == "" {
		http.Error(w, "Missing 'period' query parameter", http.StatusBadRequest)
		return
	}

	sentimentCounts, err := h.GoogleNewsService.GetTopCompanyCountDict(context.Background(), period)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sentimentCounts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *NewAnalyticsHandler) GetSourceCount(w http.ResponseWriter, r *http.Request) {
	period := r.URL.Query().Get("period")

	if period == "" {
		http.Error(w, "Missing 'period' query parameter", http.StatusBadRequest)
		return
	}

	sentimentCounts, err := h.GoogleNewsService.GetSourceCount(context.Background(), period)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sentimentCounts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *NewAnalyticsHandler) GetNewsByPeriod(w http.ResponseWriter, r *http.Request) {
	period := r.URL.Query().Get("period")

	news, err := h.GoogleNewsService.GetNewsByPeriod(period)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}
