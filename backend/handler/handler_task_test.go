package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hendrihmwn/crud-task-backend/handler/interfaces/mocks"
	"github.com/hendrihmwn/crud-task-backend/helper"
	"github.com/hendrihmwn/crud-task-backend/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type TaskHandlerTestSuite struct {
	suite.Suite
	Module          *MainInstance
	Config          helper.Config
	TaskUseCaseMock *mocks.TaskUseCase
}

func (suite *TaskHandlerTestSuite) SetupTest() {
	suite.TaskUseCaseMock = mocks.NewTaskUseCase(suite.T())
	suite.Module = &MainInstance{
		config:      helper.Config{},
		taskUseCase: suite.TaskUseCaseMock,
	}
}

func TestTaskHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskHandlerTestSuite))
}

func (suite *TaskHandlerTestSuite) TestListTaskHandler() {
	app := gin.New()
	app.GET("/test", MockToken(), suite.Module.listTask)

	tests := []struct {
		name     string
		args     model.TaskListParam
		mock     func()
		wantCode int
	}{
		{
			name:     "error - bad request",
			args:     model.TaskListParam{},
			mock:     func() {},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "error - get list",
			args: model.TaskListParam{
				Limit: 10,
				Page:  1,
			},
			mock: func() {
				suite.TaskUseCaseMock.EXPECT().ListTask(mock.Anything, mock.Anything).
					Return([]model.TaskResponse{}, 0, errors.New("some error")).Once()
			},
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			args: model.TaskListParam{
				Limit: 10,
				Page:  1,
			},
			mock: func() {
				suite.TaskUseCaseMock.EXPECT().ListTask(mock.Anything, mock.Anything).
					Return([]model.TaskResponse{{
						ID:          "XXX",
						Title:       "title",
						Description: "description",
						Status:      "backlog",
						CreatedAt:   time.Now(),
					}}, 1, nil).Once()
			},
			wantCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			tt.mock()
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/test?limit=%d&page=%d", tt.args.Limit, tt.args.Page), nil)
			req.Header.Set("Accept", "application/json")
			app.ServeHTTP(w, req)
			suite.Equal(tt.wantCode, w.Code)
		})
	}
}

func (suite *TaskHandlerTestSuite) TestGetTaskHandler() {
	app := gin.New()
	app.GET("/test/:id", MockToken(), suite.Module.getTask)

	tests := []struct {
		name     string
		args     string
		mock     func()
		wantCode int
	}{
		{
			name: "error - get",
			args: "68fc6a818c54acf4a737d7ab",
			mock: func() {
				suite.TaskUseCaseMock.EXPECT().GetTask(mock.Anything, mock.Anything).
					Return(&model.TaskResponse{}, errors.New("some error")).Once()
			},
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "error - data not found",
			args: "68fc6a818c54acf4a737d7ab",
			mock: func() {
				suite.TaskUseCaseMock.EXPECT().GetTask(mock.Anything, mock.Anything).
					Return(nil, nil).Once()
			},
			wantCode: http.StatusNotFound,
		},
		{
			name: "success",
			args: "68fc6a818c54acf4a737d7ab",
			mock: func() {
				suite.TaskUseCaseMock.EXPECT().GetTask(mock.Anything, mock.Anything).
					Return(&model.TaskResponse{
						ID:          "XXX",
						Title:       "title",
						Description: "description",
						Status:      "backlog",
						CreatedAt:   time.Now(),
					}, nil).Once()
			},
			wantCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			tt.mock()
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/test/%s", tt.args), nil)
			req.Header.Set("Accept", "application/json")
			app.ServeHTTP(w, req)
			suite.Equal(tt.wantCode, w.Code)
		})
	}
}

func (suite *TaskHandlerTestSuite) TestCreateTaskHandler() {
	app := gin.New()
	app.POST("/test", MockToken(), suite.Module.createTask)

	tests := []struct {
		name     string
		args     model.TaskBodyParam
		mock     func()
		wantCode int
	}{
		{
			name: "error - bad request",
			args: model.TaskBodyParam{},
			mock: func() {
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "error - create",
			args: model.TaskBodyParam{
				Title:       "title",
				Description: "description",
				Status:      "backlog",
			},
			mock: func() {
				suite.TaskUseCaseMock.EXPECT().CreateTask(mock.Anything, mock.Anything).
					Return(nil, errors.New("some error")).Once()
			},
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			args: model.TaskBodyParam{
				Title:       "title",
				Description: "description",
				Status:      "backlog",
			},
			mock: func() {
				suite.TaskUseCaseMock.EXPECT().CreateTask(mock.Anything, mock.Anything).
					Return(&model.TaskResponse{
						ID:          "XXX",
						Title:       "title",
						Description: "description",
						Status:      "backlog",
						CreatedAt:   time.Now(),
					}, nil).Once()
			},
			wantCode: http.StatusCreated,
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

func (suite *TaskHandlerTestSuite) TestUpdateTaskHandler() {
	app := gin.New()
	app.PUT("/test/:id", MockToken(), suite.Module.updateTask)

	tests := []struct {
		name     string
		args     model.TaskBodyParam
		mock     func()
		wantCode int
	}{
		{
			name: "error - bad request",
			args: model.TaskBodyParam{},
			mock: func() {
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "error - not found",
			args: model.TaskBodyParam{
				Title:       "title",
				Description: "description",
				Status:      "backlog",
			},
			mock: func() {
				suite.TaskUseCaseMock.EXPECT().UpdateTask(mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil).Once()
			},
			wantCode: http.StatusNotFound,
		},
		{
			name: "error - update",
			args: model.TaskBodyParam{
				Title:       "title",
				Description: "description",
				Status:      "backlog",
			},
			mock: func() {
				suite.TaskUseCaseMock.EXPECT().UpdateTask(mock.Anything, mock.Anything, mock.Anything).
					Return(&model.TaskResponse{}, errors.New("some error")).Once()
			},
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			args: model.TaskBodyParam{
				Title:       "title",
				Description: "description",
				Status:      "backlog",
			},
			mock: func() {
				suite.TaskUseCaseMock.EXPECT().UpdateTask(mock.Anything, mock.Anything, mock.Anything).
					Return(&model.TaskResponse{
						ID:          "XXX",
						Title:       "title",
						Description: "description",
						Status:      "backlog",
						CreatedAt:   time.Now(),
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
			req := httptest.NewRequest("PUT", "/test/68f9f4f82464ad9c35d5b699", bytes.NewBuffer(jsonBody))
			req.Header.Set("Accept", "application/json")
			req.Header.Set("Content-Type", "application/json")
			app.ServeHTTP(w, req)
			suite.Equal(tt.wantCode, w.Code)
		})
	}
}

func (suite *TaskHandlerTestSuite) TestDeleteTaskHandler() {
	app := gin.New()
	app.DELETE("/test/:id", MockToken(), suite.Module.deleteTask)

	tests := []struct {
		name     string
		args     string
		mock     func()
		wantCode int
	}{
		{
			name: "error - delete",
			args: "68fc6a818c54acf4a737d7ab",
			mock: func() {
				suite.TaskUseCaseMock.EXPECT().DeleteTask(mock.Anything, mock.Anything).
					Return(errors.New("some error")).Once()
			},
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			args: "68fc6a818c54acf4a737d7ab",
			mock: func() {
				suite.TaskUseCaseMock.EXPECT().DeleteTask(mock.Anything, mock.Anything).
					Return(nil).Once()
			},
			wantCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			tt.mock()
			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", fmt.Sprintf("/test/%s", tt.args), nil)
			req.Header.Set("Accept", "application/json")
			app.ServeHTTP(w, req)
			suite.Equal(tt.wantCode, w.Code)
		})
	}
}
