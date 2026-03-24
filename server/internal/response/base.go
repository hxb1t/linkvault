package response

import (
	"encoding/json"
	"net/http"
)

const HEADER_CONTENT_TYPE = "Content-Type"
const HEADER_APPLICATION_JSON = "application/json"
const UNAUTHORIZED = "unauthorized"
const INTERNAL_SERVER_ERROR = "internal server error"
const SUCCESS = "success"

type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, code int, message string, data any) {
	w.Header().Set(HEADER_CONTENT_TYPE, HEADER_APPLICATION_JSON)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ApiResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Success(w http.ResponseWriter, code int, data any) {
	JSON(w, code, SUCCESS, data)
}
