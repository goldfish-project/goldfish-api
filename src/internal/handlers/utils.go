package handlers

import (
	"encoding/json"
	"net/http"
)

// sendJSON sends a JSON stringifies object back to the requesting client
func sendJSON(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, "JSON parse error", http.StatusInternalServerError)
	}
}