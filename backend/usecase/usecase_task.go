package usecase

import (
	"context"
	"errors"
	"github.com/hendrihmwn/crud-task-backend/model"
	"github.com/hendrihmwn/crud-task-backend/usecase/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskUseCase struct {
	TaskMongoRepository interfaces.TaskMongoRepository
}

func NewTaskUseCase(taskMongoRepository interfaces.TaskMongoRepository) TaskUseCase {
	return TaskUseCase{
		TaskMongoRepository: taskMongoRepository,
	}
}

func (t TaskUseCase) ListTask(ctx context.Context, param model.TaskListParam) (res []model.TaskResponse, size int, err error) {
	filter := bson.M{}
	if param.Status != "" {
		filter = bson.M{
			"status": param.Status,
		}
	}

	list, count, err := t.TaskMongoRepository.List(
		ctx,
		filter,
		int64(param.Page),
		int64(param.Limit),
		param.SortBy,
		param.Order,
		param.Search)
	if err != nil {
		return []model.TaskResponse{}, 0, err
	}
	size = int(count)
	for _, v := range list {
		res = append(res, model.TaskResponse{
			ID:          v.ID.Hex(),
			Title:       v.Title,
			Description: v.Description,
			Status:      v.Status,
			CreatedAt:   v.CreatedAt,
		})
	}
	return
}

func (t TaskUseCase) GetTask(ctx context.Context, id string) (res *model.TaskResponse, err error) {
	data, err := t.TaskMongoRepository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return &model.TaskResponse{}, err
	}

	res = &model.TaskResponse{
		ID:          data.ID.Hex(),
		Title:       data.Title,
		Description: data.Description,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt,
	}
	return
}

func (t TaskUseCase) CreateTask(ctx context.Context, body model.TaskBodyParam) (res *model.TaskResponse, err error) {

	created, err := t.TaskMongoRepository.Create(ctx, &model.Task{
		Title:       body.Title,
		Description: body.Description,
		Status:      body.Status,
	})
	if err != nil {
		return nil, err
	}

	res = &model.TaskResponse{
		ID:          created.ID.Hex(),
		Title:       created.Title,
		Description: created.Description,
		Status:      created.Status,
		CreatedAt:   created.CreatedAt,
	}
	return
}

func (t TaskUseCase) UpdateTask(ctx context.Context, id string, body model.TaskBodyParam) (res *model.TaskResponse, err error) {
	set := bson.M{}

	if body.Title != "" {
		set["title"] = body.Title
	}
	if body.Description != "" {
		set["description"] = body.Description
	}
	if body.Status != "" {
		set["status"] = body.Status
	}

	if len(set) == 0 {
		return &model.TaskResponse{}, errors.New("no update data provided")
	}

	data, err := t.TaskMongoRepository.Update(ctx, id, set)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return &model.TaskResponse{}, err
	}

	res = &model.TaskResponse{
		ID:          data.ID.Hex(),
		Title:       data.Title,
		Description: data.Description,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt,
	}
	return
}

func (t TaskUseCase) DeleteTask(ctx context.Context, id string) (err error) {
	err = t.TaskMongoRepository.Delete(ctx, id)
	return
}
