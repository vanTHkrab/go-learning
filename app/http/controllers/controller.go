package controllers

import (
	"encoding/json"
	"net/http"
)

// BaseController represents a base controller with common functionalities.
type BaseController struct {
	// Can hold common dependencies like logger, config, etc.
}

// SendJSONResponse sends a JSON response with the given status code and data.
func (bc *BaseController) SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// SendErrorResponse sends an error JSON response.
func (bc *BaseController) SendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	bc.SendJSONResponse(w, statusCode, map[string]string{"error": message})
}
