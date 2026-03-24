package auth

import (
	"context"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hxb1t/linkvault/internal"
)

type AuthUsecase struct {
	AuthRepository AuthRepository
}

func NewAuthUsecase(ar AuthRepository) *AuthUsecase {
	return &AuthUsecase{
		AuthRepository: ar,
	}
}

func (au *AuthUsecase) CreateUser(request SignupRequest, ctx context.Context) error {
	hashedPassword, err := internal.HashPassword(request.Password)
	if err != nil {
		slog.Error("error while hashing the new password")
		return err
	}

	newUser := UserEntity{
		Username: request.Username,
		Password: hashedPassword,
	}

	err = au.AuthRepository.InsertUser(newUser, ctx)
	if err != nil {
		slog.Error("failed when trying to insert new user", "username", request.Username)
		return err
	}

	return nil
}

func (au *AuthUsecase) Login(ctx context.Context) {

}

func createToken(userId int, username string, expTime int, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"user_id":  userId,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}
