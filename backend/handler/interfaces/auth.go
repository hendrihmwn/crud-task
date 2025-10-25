package interfaces

import (
	"backend/model"
	"context"
)

type AuthUseCase interface {
	Login(ctx context.Context, param model.LoginParam) (res model.AuthResponse, err error)
}
