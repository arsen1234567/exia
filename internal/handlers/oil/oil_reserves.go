package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	publicServices "tender/internal/services/public"
)

type OilReservesHandler struct {
	SubsoilGeojsonService *publicServices.SubsoilGeojsonService
}

func (h *OilReservesHandler) GetSubsoilGeojson(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	geojsonData, err := h.SubsoilGeojsonService.GetSubsoilGeojson(ctx)
	if err != nil {
		http.Error(w, "Error retrieving geojson data", http.StatusInternalServerError)
		log.Println("Error retrieving geojson data:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(geojsonData); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
	}
}
