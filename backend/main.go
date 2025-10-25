package main

import (
	"backend/handler"
	"backend/helper"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	config := helper.LoadConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(config.MongoDBUrl)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.New()

	//r.Use(func(c *gin.Context) {
	//	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	//	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	//	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	//	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	//
	//	if c.Request.Method == "OPTIONS" {
	//		c.AbortWithStatus(204)
	//		return
	//	}
	//
	//	c.Next()
	//})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//r.Use(cors.New(cors.Config{
	//	AllowOrigins: []string{"http://localhost:5173",
	//		"http://127.0.0.1:5173"}, // Vue dev server
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))

	handler.InitHandler(r, client, config)

	// Server will listen on 0.0.0.0:8080
	r.Run()
}
