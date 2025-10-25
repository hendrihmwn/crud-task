package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hendrihmwn/crud-task-backend/handler/interfaces"
	"github.com/hendrihmwn/crud-task-backend/helper"
	mongo2 "github.com/hendrihmwn/crud-task-backend/repository/mongo"
	"github.com/hendrihmwn/crud-task-backend/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

var InstanceHandler MainInstance

type MainInstance struct {
	clientMongo *mongo.Client
	taskUseCase interfaces.TaskUseCase
	authUseCase interfaces.AuthUseCase
	config      helper.Config
}

func InitHandler(router *gin.Engine, client *mongo.Client, config helper.Config) {
	taskMongoRepository := mongo2.NewTaskRepository(client, config.DBName, config.CollectionName)
	taskUseCase := usecase.NewTaskUseCase(taskMongoRepository)
	authUseCase := usecase.NewAuthUseCase(config)

	InstanceHandler = MainInstance{
		clientMongo: client,
		taskUseCase: taskUseCase,
		authUseCase: authUseCase,
		config:      config,
	}
	registerTaskHandler(router)
	registerAuthHandler(router)
}
