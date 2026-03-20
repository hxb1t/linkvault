package main

import (
	"net/http"

	"github.com/hxb1t/linkvault/configs"
	"github.com/hxb1t/linkvault/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	config := configs.Load()
	if config.Env == "dev" {
		webPage := http.FileServer(http.Dir("../client"))
		mux.Handle("/", webPage)
	}

	// API
	mux.HandleFunc("GET /api/links", handlers.GetLinks)

	http.ListenAndServe(":"+config.Port, mux)
}
