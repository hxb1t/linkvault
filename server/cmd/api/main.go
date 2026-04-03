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

	apiMux := http.NewServeMux()
	mux.Handle(env.ContextPath+"/", http.StripPrefix(env.ContextPath, apiMux))

	// File server for static web applications
	if strings.EqualFold(env.Env, "dev") {
		webPage := http.FileServer(http.Dir("../client"))
		mux.Handle("/", webPage)
	}

	// Init SQLite Database connection
	db := database.ConnectDatabase(env.DatabasePath, env.MaxDbOpenConnectionPool, env.MaxDbIdleConnectionPool)
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
	auth.NewAuthRoute(apiMux, db, redis, env)

	slog.Info("linkvault service is ready to serve", "port", env.Port, "context path", env.ContextPath)
	http.ListenAndServe(":"+env.Port, mux)
}
