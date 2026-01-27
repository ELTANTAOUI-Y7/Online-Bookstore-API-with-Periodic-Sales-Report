package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"online-bookstore-api/models"
	"strconv"
	"strings"
)

// respondWithJSON sends a JSON response
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// respondWithError sends an error response with consistent structure
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	// Log error responses (4xx and 5xx)
	if statusCode >= 400 {
		LogError("HTTP", "Error response", nil)
		LogInfo("HTTP", "Error details", map[string]interface{}{
			"status_code": statusCode,
			"message": message,
		})
	}
	respondWithJSON(w, statusCode, models.ErrorResponse{Error: message})
}

// checkContext checks if context is done and responds appropriately
func checkContext(ctx context.Context, w http.ResponseWriter) bool {
	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.Canceled:
			respondWithError(w, http.StatusRequestTimeout, "Request was canceled")
		case context.DeadlineExceeded:
			respondWithError(w, http.StatusRequestTimeout, "Request timeout exceeded")
		default:
			respondWithError(w, http.StatusRequestTimeout, "Request context error")
		}
		return true
	default:
		return false
	}
}

// extractID extracts an ID from a URL path
func extractID(path, prefix string) (int, error) {
	// Remove prefix and trailing slashes
	idStr := strings.TrimPrefix(path, prefix)
	idStr = strings.TrimSuffix(idStr, "/")
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}
