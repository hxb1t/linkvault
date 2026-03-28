package middlewares

import (
	"log/slog"
	"net/http"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("API Request to", "method", r.Method, "url", r.URL.Path, "payload", r.GetBody)
		next.ServeHTTP(w, r)
	})
}

func MiddlewareNoAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("API Request to", "method", r.Method, "url", r.URL.Path, "payload", r.GetBody)
		next.ServeHTTP(w, r)
	})
}
