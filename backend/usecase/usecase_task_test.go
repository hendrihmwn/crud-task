package usecase_test

import (
	"context"
	"errors"
	"github.com/hendrihmwn/crud-task-backend/model"
	"github.com/hendrihmwn/crud-task-backend/usecase"
	"github.com/hendrihmwn/crud-task-backend/usecase/interfaces/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

type TaskUseCaseTestSuite struct {
	suite.Suite

	TaskMongoRepository *mocks.TaskMongoRepository
	UseCase             usecase.TaskUseCase
}

func TestTaskUseCaseSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseTestSuite))
}

func (s *TaskUseCaseTestSuite) SetupTest() {
	t := s.T()

	s.TaskMongoRepository = mocks.NewTaskMongoRepository(t)
	s.UseCase = usecase.NewTaskUseCase(
		s.TaskMongoRepository,
	)
}

func (s *TaskUseCaseTestSuite) TestListTask() {
	type args struct {
		ctx    context.Context
		params model.TaskListParam
	}
	tests := []struct {
		name       string
		args       args
		mock       func()
		afterTest  func()
		wantErr    bool
		wantErrMsg string
		size       int
	}{
		{
			name: "error - get list",
			args: args{
				ctx: context.TODO(),
				params: model.TaskListParam{
					Limit:  10,
					Page:   1,
					Status: "backlog",
				},
			},
			mock: func() {
				s.TaskMongoRepository.EXPECT().List(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, 0, errors.New("some error")).Once()
			},
			afterTest: func() {

			},
			wantErr:    true,
			wantErrMsg: "some error",
		},
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				params: model.TaskListParam{
					Limit:  10,
					Page:   1,
					Status: "backlog",
				},
			},
			mock: func() {
				s.TaskMongoRepository.EXPECT().List(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return([]*model.Task{{
						ID:          primitive.NewObjectID(),
						Title:       "TASK",
						Description: "description of task",
						Status:      "backlog",
						CreatedAt:   time.Now(),
					}}, 1, nil).Once()
			},
			afterTest: func() {

			},
			wantErr: false,
			size:    1,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock()
			_, size, err := s.UseCase.ListTask(tt.args.ctx, tt.args.params)
			tt.afterTest()
			if tt.wantErr {
				s.Error(err)
				s.EqualError(err, tt.wantErrMsg)
			} else {
				s.Equal(size, tt.size)
				s.NoError(err)
			}
		})
	}
}

func (s *TaskUseCaseTestSuite) TestGetTask() {
	type args struct {
		ctx context.Context
		id  string
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
			name: "error - get no document",
			args: args{
				ctx: context.TODO(),
				id:  "68fc6a818c54acf4a737d7ab",
			},
			mock: func() {
				s.TaskMongoRepository.EXPECT().GetByID(mock.Anything, mock.Anything).
					Return(nil, mongo.ErrNoDocuments).Once()
			},
			afterTest: func() {

			},
			wantErr:    true,
			wantErrMsg: mongo.ErrNoDocuments.Error(),
		},
		{
			name: "error - get some error",
			args: args{
				ctx: context.TODO(),
				id:  "68fc6a818c54acf4a737d7ab",
			},
			mock: func() {
				s.TaskMongoRepository.EXPECT().GetByID(mock.Anything, mock.Anything).
					Return(nil, errors.New("some error")).Once()
			},
			afterTest: func() {

			},
			wantErr:    true,
			wantErrMsg: "some error",
		},
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				id:  "68fc6a818c54acf4a737d7ab",
			},
			mock: func() {
				s.TaskMongoRepository.EXPECT().GetByID(mock.Anything, mock.Anything).
					Return(&model.Task{
						ID:          primitive.NewObjectID(),
						Title:       "TASK",
						Description: "description of task",
						Status:      "backlog",
						CreatedAt:   time.Now(),
					}, nil).Once()
			},
			afterTest: func() {

			},
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock()
			_, err := s.UseCase.GetTask(tt.args.ctx, tt.args.id)
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

func (s *TaskUseCaseTestSuite) TestCreateTask() {
	type args struct {
		ctx    context.Context
		params model.TaskBodyParam
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
			name: "error - create",
			args: args{
				ctx: context.TODO(),
				params: model.TaskBodyParam{
					Title:       "title",
					Description: "description",
					Status:      "backlog",
				},
			},
			mock: func() {
				s.TaskMongoRepository.EXPECT().Create(mock.Anything, mock.Anything).
					Return(nil, errors.New("some error")).Once()
			},
			afterTest: func() {

			},
			wantErr:    true,
			wantErrMsg: "some error",
		},
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				params: model.TaskBodyParam{
					Title:       "title",
					Description: "description",
					Status:      "backlog",
				},
			},
			mock: func() {
				s.TaskMongoRepository.EXPECT().Create(mock.Anything, mock.Anything).
					Return(&model.Task{
						ID:          primitive.NewObjectID(),
						Title:       "TASK",
						Description: "description of task",
						Status:      "backlog",
						CreatedAt:   time.Now(),
					}, nil).Once()
			},
			afterTest: func() {

			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock()
			_, err := s.UseCase.CreateTask(tt.args.ctx, tt.args.params)
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

func (s *TaskUseCaseTestSuite) TestUpdateTask() {
	type args struct {
		ctx    context.Context
		params model.TaskBodyParam
		id     string
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
			name: "error - no update data provide",
			args: args{
				ctx:    context.TODO(),
				params: model.TaskBodyParam{},
				id:     "68fc6a818c54acf4a737d7ab",
			},
			mock: func() {

			},
			afterTest: func() {

			},
			wantErr:    true,
			wantErrMsg: "no update data provided",
		},
		{
			name: "error - no document found",
			args: args{
				ctx: context.TODO(),
				params: model.TaskBodyParam{
					Title:       "title",
					Description: "description",
					Status:      "backlog",
				},
				id: "68fc6a818c54acf4a737d7ab",
			},
			mock: func() {
				s.TaskMongoRepository.EXPECT().Update(mock.Anything, mock.Anything, mock.Anything).
					Return(nil, mongo.ErrNoDocuments).Once()
			},
			afterTest: func() {

			},
			wantErr:    true,
			wantErrMsg: mongo.ErrNoDocuments.Error(),
		},
		{
			name: "error - update",
			args: args{
				ctx: context.TODO(),
				params: model.TaskBodyParam{
					Title:       "title",
					Description: "description",
					Status:      "backlog",
				},
				id: "68fc6a818c54acf4a737d7ab",
			},
			mock: func() {
				s.TaskMongoRepository.EXPECT().Update(mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errors.New("some error")).Once()
			},
			afterTest: func() {

			},
			wantErr:    true,
			wantErrMsg: "some error",
		},
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				params: model.TaskBodyParam{
					Title:       "title",
					Description: "description",
					Status:      "backlog",
				},
				id: "68fc6a818c54acf4a737d7ab",
			},
			mock: func() {
				s.TaskMongoRepository.EXPECT().Update(mock.Anything, mock.Anything, mock.Anything).
					Return(&model.Task{
						ID:          primitive.NewObjectID(),
						Title:       "TASK",
						Description: "description of task",
						Status:      "backlog",
						CreatedAt:   time.Now(),
					}, nil).Once()
			},
			afterTest: func() {

			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock()
			_, err := s.UseCase.UpdateTask(tt.args.ctx, tt.args.id, tt.args.params)
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

func (s *TaskUseCaseTestSuite) TestDeleteTask() {
	type args struct {
		ctx context.Context
		id  string
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
			name: "error - get some error",
			args: args{
				ctx: context.TODO(),
				id:  "68fc6a818c54acf4a737d7ab",
			},
			mock: func() {
				s.TaskMongoRepository.EXPECT().Delete(mock.Anything, mock.Anything).
					Return(errors.New("some error")).Once()
			},
			afterTest: func() {

			},
			wantErr:    true,
			wantErrMsg: "some error",
		},
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				id:  "68fc6a818c54acf4a737d7ab",
			},
			mock: func() {
				s.TaskMongoRepository.EXPECT().Delete(mock.Anything, mock.Anything).
					Return(nil).Once()
			},
			afterTest: func() {

			},
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock()
			err := s.UseCase.DeleteTask(tt.args.ctx, tt.args.id)
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
