package handlers

import "net/http"

func GetLinks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Golang!"))
}
