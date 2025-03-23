package api

import (
	"encoding/json"
	"net/http"

	"gitgithub.com/VarunPandya07/credit-card-validation/internal/luhn"
)

type CardRequest struct {
	Number string `json:"number"`
}

type CardResponse struct {
	Valid bool   `json:"valid"`
	Error string `json:"error,omitempty"`
}

func CardValidationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}
	valid := luhn.IsVelid(req.Number)
	res := CardResponse{Valid: valid}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
