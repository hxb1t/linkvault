package auth

import (
	"database/sql"
	"net/http"

	"github.com/hxb1t/linkvault/internal"
)

const LOGIN_API = "POST /api/login"
const SIGNUP_API = "POST /api/signup"

func NewAuthRoute(mux *http.ServeMux, db *sql.DB) {
	// Depedencies
	authRepository := NewAuthRepository(db)
	authService := NewAuthUsecase(*authRepository)
	authHandler := NewAuthHandler(*authService, *authRepository)

	// Routes
	mux.Handle(LOGIN_API, internal.MiddlewareNoAuth(http.HandlerFunc(authHandler.Login)))
	mux.Handle(SIGNUP_API, internal.MiddlewareNoAuth(http.HandlerFunc(authHandler.SignUp)))
}
