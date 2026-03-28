package socialmedia

import "net/http"

type SocialMediaHandler struct {
}

func NewSocialMediaHandler() *SocialMediaHandler {
	return &SocialMediaHandler{}
}

func (s *SocialMediaHandler) GetLinks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Golang!"))
}

func (s *SocialMediaHandler) CreateSocialMedia(w http.ResponseWriter, r *http.Request) {}

func (s *SocialMediaHandler) UpdateSocialMedia(w http.ResponseWriter, r *http.Request) {}

func (s *SocialMediaHandler) DeleteSocialMedia(w http.ResponseWriter, r *http.Request) {}
