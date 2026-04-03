package auth

import (
	"encoding/json"
	"net/http"

	"github.com/hxb1t/linkvault/internal/domain"
)

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
	var request LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		domain.BadRequest(w, r, domain.ErrInvalidRequest)
		return
	}

	response, err := h.Usecase.Login(r.Context(), request)
	if err != nil {
		domain.Error(w, r, http.StatusBadRequest, err)
		return
	}

	domain.Success(w, http.StatusOK, response)
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) RefreshSession(w http.ResponseWriter, r *http.Request) {

}
