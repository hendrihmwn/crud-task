package interfaces

import (
	"context"
	"github.com/hendrihmwn/crud-task-backend/model"
	"go.mongodb.org/mongo-driver/bson"
)

//go:generate mockery --name=TaskMongoRepository --keeptree --output=mocks --case=underscore --with-expecter=true
type TaskMongoRepository interface {
	List(ctx context.Context, filter bson.M, page, limit int64, sortField string, sortOrder int, searchText string) ([]*model.Task, int64, error)
	GetByID(ctx context.Context, id string) (req *model.Task, err error)
	Create(ctx context.Context, req *model.Task) (res *model.Task, err error)
	Update(ctx context.Context, id string, update bson.M) (res *model.Task, err error)
	Delete(ctx context.Context, id string) error
}
