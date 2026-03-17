package http

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	ErrorCode    string `json:"code"`
	ErrorMessage string `json:"message"`
	ErrorData    any    `json:"data"`
}

func SendResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func SendErrorResponse(w http.ResponseWriter, status int, code, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errResp := ErrorResponse{
		ErrorCode:    code,
		ErrorMessage: message,
		ErrorData:    data,
	}

	json.NewEncoder(w).Encode(errResp)
}

func NotFound(w http.ResponseWriter) {
	SendErrorResponse(w, http.StatusNotFound, "notFound", "Not found", nil)
}

func FailedCreate(w http.ResponseWriter) {
	SendErrorResponse(w, http.StatusBadRequest, "failedCreate", "Failed create", nil)
}

func FailedUpdate(w http.ResponseWriter) {
	SendErrorResponse(w, http.StatusBadRequest, "failedUpdate", "Failed update", nil)
}

func FailedDelete(w http.ResponseWriter) {
	SendErrorResponse(w, http.StatusBadRequest, "failedDelete", "Failed delete", nil)
}
