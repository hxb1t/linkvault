package profile

import (
	"database/sql"
	"net/http"
)

const GET_PROFILE_API = "GET /api/profile"
const GET_PROFILE_BY_ID_API = "GET /api/profile/{id}"
const CREATE_NEW_PROFILE_API = "POST /api/profile"
const UPDATE_PROFILE_API = "PATCH /api/profile/{id}"
const DELETE_PROFILE_API = "DELETE /api/profile/{id}"

func NewProfileRoute(http *http.ServeMux, db *sql.DB) {
	// Depedencies
	profileRepository := NewProfileRepository(db)
	profileUsecase := NewProfileUsecase(*profileRepository)
	profileHandler := NewProfileHandler(*profileUsecase)

	// Routes
	http.HandleFunc(GET_PROFILE_API, profileHandler.GetProfile)
	http.HandleFunc(GET_PROFILE_BY_ID_API, profileHandler.GetProfileById)
	http.HandleFunc(CREATE_NEW_PROFILE_API, profileHandler.CreateProfile)
	http.HandleFunc(UPDATE_PROFILE_API, profileHandler.UpdateProfile)
	http.HandleFunc(DELETE_PROFILE_API, profileHandler.DeleteProfile)
}
