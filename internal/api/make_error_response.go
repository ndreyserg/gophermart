package api

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func (a *api) makeErrorResponse(w http.ResponseWriter, message string, code int) {
	w.Header().Set(contentTypeHeader, contentTypeJSON)
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	_ = enc.Encode(errorResponse{Message: message})
}
