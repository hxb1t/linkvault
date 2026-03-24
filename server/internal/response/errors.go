package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request, code int, message string, err error) {
	if err != nil {
		slog.ErrorContext(r.Context(), message,
			"error", err.Error(),
			"method", r.Method,
			"path", r.URL.Path,
			"code", code,
		)
	}

	w.Header().Set(HEADER_CONTENT_TYPE, HEADER_APPLICATION_JSON)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ApiError{
		Code:    code,
		Message: message,
	})
}

func BadRequest(w http.ResponseWriter, r *http.Request, message string, err error) {
	Error(w, r, http.StatusBadRequest, message, err)
}

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	Error(w, r, http.StatusUnauthorized, UNAUTHORIZED, nil)
}

func InternalError(w http.ResponseWriter, r *http.Request, err error) {
	Error(w, r, http.StatusInternalServerError, INTERNAL_SERVER_ERROR, err)
}
