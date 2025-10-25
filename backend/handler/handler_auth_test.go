package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hendrihmwn/crud-task-backend/handler/interfaces/mocks"
	"github.com/hendrihmwn/crud-task-backend/helper"
	"github.com/hendrihmwn/crud-task-backend/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AuthHandlerTestSuite struct {
	suite.Suite
	Module          *MainInstance
	Config          helper.Config
	AuthUseCaseMock *mocks.AuthUseCase
}

func (suite *AuthHandlerTestSuite) SetupTest() {
	suite.AuthUseCaseMock = mocks.NewAuthUseCase(suite.T())
	suite.Module = &MainInstance{
		config:      helper.Config{},
		authUseCase: suite.AuthUseCaseMock,
	}
}

func TestAuthHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthHandlerTestSuite))
}

func (suite *AuthHandlerTestSuite) TestLoginAuthHandler() {
	app := gin.New()
	app.POST("/test", suite.Module.login)

	tests := []struct {
		name     string
		args     model.LoginParam
		mock     func()
		wantCode int
	}{
		{
			name:     "error - bad request",
			args:     model.LoginParam{},
			mock:     func() {},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "error - login",
			args: model.LoginParam{
				Username: "admin",
				Password: "123",
			},
			mock: func() {
				suite.AuthUseCaseMock.EXPECT().Login(mock.Anything, mock.Anything).
					Return(model.AuthResponse{}, errors.New("some error")).Once()
			},
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			args: model.LoginParam{
				Username: "admin",
				Password: "123",
			},
			mock: func() {
				suite.AuthUseCaseMock.EXPECT().Login(mock.Anything, mock.Anything).
					Return(model.AuthResponse{
						Token:          "xxx",
						ExpirationTime: 123,
					}, nil).Once()
			},
			wantCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			tt.mock()
			jsonBody, _ := json.Marshal(tt.args)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/test", bytes.NewBuffer(jsonBody))
			req.Header.Set("Accept", "application/json")
			req.Header.Set("Content-Type", "application/json")
			app.ServeHTTP(w, req)
			suite.Equal(tt.wantCode, w.Code)
		})
	}
}
