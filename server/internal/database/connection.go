package database

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	_ "github.com/glebarez/go-sqlite"
	"github.com/redis/go-redis/v9"
)

func ConnectDatabase(path string) *sql.DB {
	db, err := sql.Open("sqlite", path)
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

func ConnectRedis(addr, password string, database int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       database,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		slog.Error("redis connection failed", "error", err, "address", addr)
		return nil, err
	}

	slog.Info("conncted to redis", "addr", addr)
	return client, nil
}
