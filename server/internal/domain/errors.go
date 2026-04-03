package domain

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	if err != nil {
		slog.ErrorContext(r.Context(), err.Error(),
			"method", r.Method,
			"path", r.URL.Path,
			"code", code,
		)
	}

	w.Header().Set(HEADER_CONTENT_TYPE, HEADER_APPLICATION_JSON)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ApiError{
		Code:    code,
		Message: err.Error(),
	})
}

func BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	Error(w, r, http.StatusBadRequest, err)
}

func Unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	Error(w, r, http.StatusUnauthorized, err)
}

func InternalError(w http.ResponseWriter, r *http.Request, err error) {
	Error(w, r, http.StatusInternalServerError, err)
}
