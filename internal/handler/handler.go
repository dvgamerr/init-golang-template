package handler

import (
	"encoding/json"
	"net/http"
)

// Handler represents HTTP handlers
type Handler struct {
	// Add your dependencies here (service, logger, etc.)
}

// NewHandler creates a new handler instance
func NewHandler() *Handler {
	return &Handler{}
}

// HealthCheck handles health check requests
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
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
