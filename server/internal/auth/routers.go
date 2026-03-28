package auth

import (
	"database/sql"
	"net/http"

	"github.com/hxb1t/linkvault/configs"
	"github.com/redis/go-redis/v9"
)

const LOGIN_API = "POST /api/auth/login"
const REFRESH_AUTH_API = "POST /api/auth/refresh"
const LOGOUT_API = "POST /api/auth/logout"
const SIGNUP_API = "POST /api/auth/signup"

func NewAuthRoute(mux *http.ServeMux, db *sql.DB, redis *redis.Client, env configs.Env) {
	// Depedencies
	authRepository := NewAuthRepository(db)
	authService := NewAuthUsecase(*authRepository, *redis, env)
	authHandler := NewAuthHandler(*authService, *authRepository)

	// Routes
	mux.Handle(LOGIN_API, http.HandlerFunc(authHandler.Login))
	mux.Handle(SIGNUP_API, http.HandlerFunc(authHandler.SignUp))
	mux.Handle(LOGOUT_API, http.HandlerFunc(authHandler.Login))
}
