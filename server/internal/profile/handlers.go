package profile

import "net/http"

type ProfileHandler struct {
	ProfileUsecase ProfileUsecase
}

func NewProfileHandler(pu ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{
		ProfileUsecase: pu,
	}
}

func (ph *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {

}

func (ph *ProfileHandler) GetProfileById(w http.ResponseWriter, r *http.Request) {

}

func (ph *ProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {

}

func (ph *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {

}

func (ph *ProfileHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {

}
