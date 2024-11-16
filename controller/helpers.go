package controller

import (
	"aosmanova/doodocs/models"
	"encoding/json"

	"net/http"
)

func EncodeOK(w http.ResponseWriter, data any) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func serverError(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Add("Content Type", "application/json")
	json.NewEncoder(w).Encode(err)
}
func httpError(w http.ResponseWriter, message string, StatusCode int) {
	w.WriteHeader(StatusCode)
	errResponse := models.ErrorResponse{
		Message: message,
	}
	w.Header().Add("Content Type", "application/json")
	json.NewEncoder(w).Encode(errResponse)
}
