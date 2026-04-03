package auth

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/hxb1t/linkvault/configs"
	"github.com/hxb1t/linkvault/internal/domain"
	"github.com/hxb1t/linkvault/internal/utils"
	"github.com/redis/go-redis/v9"
)

type Usecase struct {
	Repository Repository
	Redis      redis.Client
	Env        configs.Env
}

func NewAuthUsecase(ar Repository, redis redis.Client, env configs.Env) *Usecase {
	return &Usecase{
		Repository: ar,
		Redis:      redis,
		Env:        env,
	}
}

func (au *Usecase) CreateUser(ctx context.Context, request SignupRequest) error {
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		slog.Error("error while hashing the new password")
		return err
	}

	newUser := UserEntity{
		Username: request.Username,
		Password: hashedPassword,
	}

	err = au.Repository.InsertUser(newUser, ctx)
	if err != nil {
		slog.Error("failed when trying to insert new user", "username", request.Username)
		return err
	}

	return nil
}

func (au *Usecase) Login(ctx context.Context, request LoginRequest) (LoginResponse, error) {
	userSessionRedisKey := domain.UserSessionKey(request.Username)
	existingUserSessionCache, err := au.Redis.Get(ctx, userSessionRedisKey).Result()
	response := LoginResponse{}

	if err == nil && len(existingUserSessionCache) > 0 {
		slog.Info("session is still alive", "username", request.Username)
		if err = json.Unmarshal([]byte(existingUserSessionCache), &response); err != nil {
			slog.Error("failed when read user session cache value", "value", existingUserSessionCache)
			return LoginResponse{}, err
		}
		slog.Debug("success read user session", "username", request.Username, "value", existingUserSessionCache)
		return response, nil
	}

	existingUser, err := au.Repository.GetUserByUsername(ctx, request.Username)
	if err != nil {
		slog.Error("failed when get existing user by username", "username", request.Username)
		return LoginResponse{}, err
	}

	if err := utils.CheckPassword(existingUser.Password, request.Password); err != nil {
		slog.Error("invalid password", "username", request.Username)
		return LoginResponse{}, domain.ErrInvalidPassword
	}

	accessToken, err := utils.CreateJWT(existingUser.Id, existingUser.Username, au.Env.JWTExpTime, au.Env.JWTSecret)
	if err != nil {
		slog.Error("failed when generate jwt token", "error", err)
		return LoginResponse{}, domain.ErrInternalServer
	}

	userSessionTtl := time.Duration(au.Env.UserSessionTTL) * time.Second
	response = LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: "",
		ExpTime:      au.Env.JWTExpTime,
	}

	userSessionCache, err := json.Marshal(response)
	if err != nil {
		slog.Error("failed to marshalling user session cache", "error", err)
		return LoginResponse{}, domain.ErrInternalServer
	}

	err = au.Redis.Set(ctx, userSessionRedisKey, string(userSessionCache), userSessionTtl).Err()
	if err != nil {
		slog.Warn("failed to store user session cache", "key", userSessionRedisKey)
	}

	return response, nil
}
