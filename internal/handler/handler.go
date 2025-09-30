package handler

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	response := map[string]interface{}{
		"status":  "ok",
		"message": "Service is healthy",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
