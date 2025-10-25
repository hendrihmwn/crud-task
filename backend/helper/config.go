package helper

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	MongoDBUrl string
	JWTSecret  string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	return Config{
		MongoDBUrl: os.Getenv("MONGODB_URL"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}
