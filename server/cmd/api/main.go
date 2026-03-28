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
	env := configs.LoadEnv()

	// File server for static web applications
	if strings.EqualFold(env.Env, "dev") {
		webPage := http.FileServer(http.Dir("../client"))
		mux.Handle("/", webPage)
	}

	// Init SQLite Database connection
	db := database.ConnectDatabase(env.DatabasePath)
	// Close the database connection once the application server is shutting down
	defer db.Close()

	// run table migrations, if failed stop the application
	if err := database.Migrate(db); err != nil {
		slog.Error("migration database failed", "error", err)
		return
	}

	// Init Redis connection
	redis, err := database.ConnectRedis(env.RedisHost, env.RedisPassword, env.RedisDatabase)
	if err != nil {
		panic("failed connect to redis: " + err.Error())
	}

	// Initilize routes
	auth.NewAuthRoute(mux, db, redis, env)

	http.ListenAndServe(":"+env.Port, mux)
}
