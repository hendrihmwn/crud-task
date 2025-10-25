package handler

import (
	"backend/handler/interfaces"
	"backend/helper"
	mongo2 "backend/repository/mongo"
	"backend/usecase"
	"github.com/gin-gonic/gin"
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
	taskMongoRepository := mongo2.NewTaskRepository(client, "database", "tasks")
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
