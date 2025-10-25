package usecase_test

import (
	"context"
	"github.com/hendrihmwn/crud-task-backend/helper"
	"github.com/hendrihmwn/crud-task-backend/model"
	"github.com/hendrihmwn/crud-task-backend/usecase"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuthUseCaseTestSuite struct {
	suite.Suite

	Config  helper.Config
	UseCase usecase.AuthUseCase
}

func TestAuthUseCaseSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseTestSuite))
}

func (s *AuthUseCaseTestSuite) SetupTest() {
	s.Config = helper.LoadConfig()
	s.UseCase = usecase.NewAuthUseCase(
		s.Config,
	)
}

func (s *AuthUseCaseTestSuite) TestLogin() {
	type args struct {
		ctx    context.Context
		params model.LoginParam
	}
	tests := []struct {
		name       string
		args       args
		mock       func()
		afterTest  func()
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "error - invalid username password",
			args: args{
				ctx:    context.TODO(),
				params: model.LoginParam{},
			},
			mock: func() {

			},
			afterTest: func() {

			},
			wantErr:    true,
			wantErrMsg: "invalid username or password",
		},
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				params: model.LoginParam{
					Username: "admin",
					Password: "password",
				},
			},
			mock: func() {

			},
			afterTest: func() {

			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock()
			_, err := s.UseCase.Login(tt.args.ctx, tt.args.params)
			tt.afterTest()
			if tt.wantErr {
				s.Error(err)
				s.EqualError(err, tt.wantErrMsg)
			} else {
				s.NoError(err)
			}
		})
	}
}
