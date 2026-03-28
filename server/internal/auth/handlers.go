package auth

import "net/http"

type Handler struct {
	Usecase    Usecase
	Repository Repository
}

func NewAuthHandler(usecase Usecase, repository Repository) *Handler {
	return &Handler{
		Usecase:    usecase,
		Repository: repository,
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {

}
