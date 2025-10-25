package usecase

import (
	"backend/helper"
	"backend/model"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthUseCase struct {
	config helper.Config
}

func NewAuthUseCase(config helper.Config) AuthUseCase {
	return AuthUseCase{
		config: config,
	}
}

func (a AuthUseCase) Login(ctx context.Context, param model.LoginParam) (res model.AuthResponse, err error) {
	// Mock credentials
	const mockUser = "admin"
	const mockPass = "password"
	jwtSecret := a.config.JWTSecret

	if param.Username != mockUser || param.Password != mockPass {
		err = errors.New("invalid username or password")
		return
	}

	// Create JWT token (HMAC SHA256)
	expiryTime := time.Now().Add(24 * time.Hour).Unix()
	claims := jwt.MapClaims{
		"sub": param.Username,
		"iat": time.Now().Unix(),
		"exp": expiryTime,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := t.SignedString([]byte(jwtSecret))
	if err != nil {
		return
	}

	res = model.AuthResponse{
		Token:          signed,
		ExpirationTime: time.Now().Add(24 * time.Hour).Unix(),
	}
	return
}
