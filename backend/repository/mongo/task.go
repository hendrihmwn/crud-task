package mongo

import (
	"backend/model"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository struct {
	coll *mongo.Collection
}

func NewTaskRepository(client *mongo.Client, dbName, collName string) *TaskRepository {
	coll := client.Database(dbName).Collection(collName)
	// ensure simple indexes (e.g., on created_at and title) -- ignore errors here
	_ = ensureIndexes(context.Background(), coll)
	return &TaskRepository{coll: coll}
}

func ensureIndexes(ctx context.Context, coll *mongo.Collection) error {
	indexes := coll.Indexes()
	models := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "created_at", Value: -1}},
		},
		{
			Keys:    bson.D{{Key: "title", Value: 1}},
			Options: options.Index().SetBackground(true),
		},
	}
	_, err := indexes.CreateMany(ctx, models)
	return err
}

func (r *TaskRepository) Create(ctx context.Context, req *model.Task) (res *model.Task, err error) {
	now := time.Now().UTC()
	if req == nil {
		return nil, errors.New("task is nil")
	}
	req.CreatedAt = now
	req.UpdatedAt = now
	if req.ID.IsZero() {
		req.ID = primitive.NewObjectID()
	}
	_, err = r.coll.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id string) (req *model.Task, err error) {
	var t model.Task
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &t, err
	}
	if err := r.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&t); err != nil {
		return &t, err
	}
	return &t, nil
}

func (r *TaskRepository) List(ctx context.Context, filter bson.M, page, limit int64, sortField string, sortOrder int, searchText string) ([]*model.Task, int64, error) {
	if filter == nil {
		filter = bson.M{}
	}

	// apply simple text search (case-insensitive) across common fields
	if searchText != "" {
		search := bson.M{"$or": []bson.M{
			{"title": bson.M{"$regex": searchText, "$options": "i"}},
		}}
		if len(filter) == 0 {
			filter = search
		} else {
			filter = bson.M{"$and": []bson.M{filter, search}}
		}
	}

	// total matching documents (ignores pagination)
	total, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var findOpts options.FindOptions

	// apply pagination if limit > 0
	if limit > 0 {
		if page <= 0 {
			page = 1
		}
		skip := (page - 1) * limit
		findOpts.Limit = &limit
		findOpts.Skip = &skip
	}

	// apply sorting if provided (default order = 1 if invalid)
	if sortField != "" {
		if sortOrder != 1 && sortOrder != -1 {
			sortOrder = 1
		}
		findOpts.Sort = bson.D{{Key: sortField, Value: sortOrder}}
	}

	cur, err := r.coll.Find(ctx, filter, &findOpts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var result []*model.Task
	for cur.Next(ctx) {
		var t model.Task
		if err := cur.Decode(&t); err != nil {
			return nil, 0, err
		}
		result = append(result, &t)
	}
	if err := cur.Err(); err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (r *TaskRepository) Update(ctx context.Context, id string, data bson.M) (res *model.Task, err error) {
	var updated model.Task
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// prevent updating the document ID
	delete(data, "_id")

	data["updated_at"] = time.Now().UTC()
	updateDoc := bson.D{{Key: "$set", Value: data}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	if err := r.coll.FindOneAndUpdate(ctx, bson.M{"_id": oid}, updateDoc, opts).Decode(&updated); err != nil {
		return &updated, err
	}
	return &updated, nil
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res, err := r.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (r *TaskRepository) Count(ctx context.Context, filter bson.M) (int64, error) {
	if filter == nil {
		filter = bson.M{}
	}
	return r.coll.CountDocuments(ctx, filter)
}
