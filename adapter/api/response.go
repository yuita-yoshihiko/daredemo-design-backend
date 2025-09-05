package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

var (
	ErrFailedToFetch           = map[string]string{"error": "failed to fetch data"}
	ErrFailedToPost            = map[string]string{"error": "failed to post data"}
	ErrInvalidRequest          = map[string]string{"error": "invalid request body"}
	ErrFailedToExternalRequest = map[string]string{"error": "failed to external request"}
	ErrNotFound                = map[string]string{"error": "data not found"}
	ErrRequestTooLarge         = map[string]string{"error": "request body is too large"}
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("Failed to encode response", "error", err.Error())
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
