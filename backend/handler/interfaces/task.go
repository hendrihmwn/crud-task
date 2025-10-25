package interfaces

import (
	"backend/model"
	"context"
)

type TaskUseCase interface {
	ListTask(ctx context.Context, param model.TaskListParam) (res []model.TaskResponse, size int, err error)
	GetTask(ctx context.Context, id string) (res *model.TaskResponse, err error)
	CreateTask(ctx context.Context, body model.TaskBodyParam) (res *model.TaskResponse, err error)
	UpdateTask(ctx context.Context, id string, body model.TaskBodyParam) (res *model.TaskResponse, err error)
	DeleteTask(ctx context.Context, id string) (err error)
}
