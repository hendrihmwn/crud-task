package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TaskListParam struct {
	Limit  uint64 `form:"limit" binding:"required" json:"limit"`
	Page   uint64 `form:"page" binding:"required" json:"page"`
	Search string `form:"search" json:"search"`
	Status string `form:"status" json:"status"`
	SortBy string `form:"sort_by" json:"sort_by"`
	Order  int    `form:"order" json:"order"`
}

type TaskGetParam struct {
	ID string `uri:"id"`
}

type TaskBodyParam struct {
	Title       string `form:"title" binding:"required,max=100" json:"title"`
	Description string `form:"description" binding:"required,max=255" json:"description"`
	Status      string `form:"status" binding:"required" json:"status"`
}

type TaskResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskMeta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Status      string             `bson:"status" json:"status"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
