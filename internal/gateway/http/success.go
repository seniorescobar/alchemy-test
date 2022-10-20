package http

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Success bool `json:"success"`
}

func renderSuccess(w http.ResponseWriter, status int) {
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(SuccessResponse{true}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
