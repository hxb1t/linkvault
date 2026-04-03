package middlewares

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/hxb1t/linkvault/internal/domain"
	"github.com/hxb1t/linkvault/internal/utils"
)

type contextKey string

const claimsKey contextKey = "auth_claims"

func Trace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("API Request to", "method", r.Method, "url", r.URL.Path, "payload", r.GetBody)
		next.ServeHTTP(w, r)
	})
}

func Auth(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken := r.Header.Get("Authorization")
			claims, err := utils.ValidateJWT(accessToken, secretKey)
			if err != nil {
				slog.Error("failed when validating jwt", "error", err)
				domain.Error(w, r, http.StatusUnauthorized, "Unauhtorized", domain.ErrUnauthorized)
				return
			}

			ctx := SetContextClaims(r.Context(), claims)
			slog.Debug("success validate user access token", "claims", claims)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Timeout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func SetContextClaims(ctx context.Context, claims *domain.AuthClaims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

func GetContextClaims(ctx context.Context) (*domain.AuthClaims, bool) {
	claims, ok := ctx.Value(claimsKey).(*domain.AuthClaims)
	return claims, ok
}
