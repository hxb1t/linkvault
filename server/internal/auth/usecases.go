package auth

import (
	"context"
	"log/slog"

	"github.com/hxb1t/linkvault/configs"
	"github.com/hxb1t/linkvault/internal/utils"
	"github.com/redis/go-redis/v9"
)

type AuthUsecase struct {
	AuthRepository AuthRepository
	Redis          redis.Client
	Env            configs.Env
}

func NewAuthUsecase(ar AuthRepository, redis redis.Client, env configs.Env) *AuthUsecase {
	return &AuthUsecase{
		AuthRepository: ar,
		Redis:          redis,
		Env:            env,
	}
}

func (au *AuthUsecase) CreateUser(request SignupRequest, ctx context.Context) error {
	hashedPassword, err := utils.HashPassword(request.Password)
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
