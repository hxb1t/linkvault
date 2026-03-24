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
