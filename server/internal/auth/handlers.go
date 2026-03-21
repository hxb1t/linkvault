package auth

import "net/http"

type AuthHandler struct {
	AuthUsecase    AuthUsecase
	AuthRepository AuthRepository
}

func NewAuthHandler(au AuthUsecase, ar AuthRepository) *AuthHandler {
	return &AuthHandler{
		AuthUsecase:    au,
		AuthRepository: ar,
	}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	return
}

func (ah *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	return
}

func MiddlewareAuth(next http.Handler) http.Handler {
	// Todo handle JWT auth
	return nil
}

func MiddlewareNoAuth(next http.Handler) http.Handler {
	return nil
}
