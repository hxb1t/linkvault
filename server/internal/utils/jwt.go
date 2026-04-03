package utils

import (
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hxb1t/linkvault/internal/domain"
)

func CreateJWT(userId int, username string, expDurationInSeconds int, secretKey string) (string, error) {
	expDuration := time.Duration(expDurationInSeconds) * time.Second
	claims := domain.AuthClaims{
		UserId:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	slog.Debug("jwt claims", "claims", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func ValidateJWT(tokenString string, secretKey string) (*domain.AuthClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		slog.Error("failed when parse user token", "error", err)
		return nil, err
	}

	claims, ok := token.Claims.(*domain.AuthClaims)
	if !ok || !token.Valid {
		return nil, domain.ErrUnauthorized
	}

	return claims, err
}
