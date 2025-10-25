package interfaces

import (
	"context"
	"github.com/hendrihmwn/crud-task-backend/model"
)

//go:generate mockery --name=AuthUseCase --keeptree --output=mocks --case=underscore --with-expecter=true
type AuthUseCase interface {
	Login(ctx context.Context, param model.LoginParam) (res model.AuthResponse, err error)
}
