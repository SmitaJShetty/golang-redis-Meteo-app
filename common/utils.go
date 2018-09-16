package common

import (
	"encoding/json"
	"net/http"
)

//AppError application error
type AppError struct {
	Message    string
	StatusCode int
}

//NewAppError returns an apperror
func NewAppError(msg string, statusCode int) *AppError {
	return &AppError{
		Message:    msg,
		StatusCode: statusCode,
	}
}

//SendResult sends result over http response
func SendResult(w http.ResponseWriter, r *http.Request, resultJSON []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}

//SendErrorResponse sends error response
func SendErrorResponse(w http.ResponseWriter, r *http.Request, appErr *AppError) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	json, _ := json.Marshal(appErr)
	w.Write(json)
}
