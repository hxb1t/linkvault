package socialmedia

import (
	"net/http"
)

type SocialMediaRouter struct {
}

// Endpoints
const GET_SOCIAL_MEDIA_API = "GET /api/socialmedia"
const CREATE_NEW_SOCIAL_MEDIA_API = "POST /api/socialmedia"
const UPDATE_SOCIAL_MEDIA_API = "PATCH /api/socialmedia/{id}"
const DELETE_SOCIAL_MEDIA = "DELETE /api/socialmedia/{id}"

func SocialMediaRoute(http *http.ServeMux) {
	smh := NewSocialMediaHandler()

	http.HandleFunc(GET_SOCIAL_MEDIA_API, smh.GetLinks)
	http.HandleFunc(CREATE_NEW_SOCIAL_MEDIA_API, smh.CreateSocialMedia)
	http.HandleFunc(UPDATE_SOCIAL_MEDIA_API, smh.UpdateSocialMedia)
	http.HandleFunc(DELETE_SOCIAL_MEDIA, smh.DeleteSocialMedia)
}
