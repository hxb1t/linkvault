package database

import (
	"database/sql"
	"log/slog"
)

func Connect(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		slog.Error("error when open db connection", "error", err)
	}

	db.Exec("PRAGMA journal_mode=WAL")
	db.Exec("PRAGMA busy_timeout=5000")
	db.Exec("PRAGMA foreign_keys=ON")

	db.SetMaxOpenConns(1)

	if err := db.Ping(); err != nil {
		slog.Error("error when trying to ping into db", "error", err)
	}

	slog.Info("connected to db", "db path", path)
	return db
}
