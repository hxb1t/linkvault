package auth

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/hxb1t/linkvault/internal/domain"
)

type Repository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (ar *Repository) InsertUser(newUser UserEntity, ctx context.Context) error {
	slog.Info("insert new user into database")
	slog.Debug("insert new user into database with details", "newUser", newUser)
	_, err := ar.DB.ExecContext(ctx, `INSERT INTO USER (USERNAME, PASSWORD) VALUES (?, ?)`,
		newUser.Username, newUser.Password,
	)

	if err != nil {
		slog.Error("failed while insert new user to database", "error", err)
		return err
	}

	return nil
}

func (ar *Repository) GetUserById(userId int, ctx context.Context) (UserEntity, error) {
	slog.Info("user user table from database by user id")
	slog.Debug("select user database with", "user id", userId)
	var user UserEntity
	err := ar.DB.QueryRowContext(ctx, `SELECT ID, USERNAME, PASSWORD FROM USER WHERE ID = ?`,
		userId).Scan(&user.Id, &user.Username, &user.Password)

	if err == sql.ErrNoRows {
		return UserEntity{}, domain.ErrUserNotFound
	}

	if err != nil {
		slog.Error("error when execute query", "error", err)
		return UserEntity{}, domain.ErrExecuteQuery
	}

	return user, err
}

func (ar *Repository) GetUserByUsername(ctx context.Context, username string) (UserEntity, error) {
	slog.Info("select user database by username")
	slog.Debug("select user database by", "username", username)
	var user UserEntity
	err := ar.DB.QueryRowContext(ctx, `SELECT ID, USERNAME, PASSWORD FROM USER WHERE USERNAME = ?`,
		username).Scan(&user.Id, &user.Username, &user.Password)

	return user, err
}
