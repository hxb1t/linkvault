package auth

import (
	"database/sql"
	"net/http"
)

const LOGIN_API = "POST /api/login"
const SIGNUP_API = "POST /api/signup"

func NewAuthRoute(http *http.ServeMux, db *sql.DB) {
	// Depedencies
	authRepository := NewAuthRepository(db)
	authService := NewAuthUsecase(*authRepository)
	authHandler := NewAuthHandler(*authService, *authRepository)

	// Routes
	http.HandleFunc(LOGIN_API, authHandler.Login)
	http.HandleFunc(SIGNUP_API, authHandler.SignUp)
}
