package main

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/hxb1t/linkvault/configs"
	"github.com/hxb1t/linkvault/internal/auth"
	"github.com/hxb1t/linkvault/internal/database"
)

func main() {
	mux := http.NewServeMux()
	config := configs.Load()

	// File server for static web applications
	if strings.EqualFold(config.Env, "dev") {
		webPage := http.FileServer(http.Dir("../client"))
		mux.Handle("/", webPage)
	}

	// Init SQLite Database Connection
	db := database.ConnectDatabase(config.DatabasePath)
	// Close the database connection once the application server is shutting down
	defer db.Close()

	// run table migrations, if failed stop the application
	if err := database.Migrate(db); err != nil {
		slog.Error("migration database failed", "error", err)
		return
	}

	// Initilize routes
	auth.NewAuthRoute(mux, db)

	http.ListenAndServe(":"+config.Port, mux)
}
