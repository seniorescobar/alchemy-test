package http

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func renderError(w http.ResponseWriter, status int, msg string) {
	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Error: msg,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
}
