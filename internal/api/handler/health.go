package handler

import (
	"encoding/json"
	"net/http"

	"github.com/francisco3ferraz/go-nuts/internal/model"
)

func Health(w http.ResponseWriter, _ *http.Request) {
	respondJSON(w, http.StatusOK, model.HealthResponse{Status: "ok"})
}

func Ping(w http.ResponseWriter, _ *http.Request) {
	respondJSON(w, http.StatusOK, model.PingResponse{Message: "pong"})
}

func respondJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
