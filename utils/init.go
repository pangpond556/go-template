package utils

import (
	"os"

	"github.com/joho/godotenv"
)

// MongoURI - mongodb connection string
var MongoURI string

// Init - init app
func Init() error {
	var err error
	godotenv.Load()

	MongoURI = os.Getenv("MONGO_URI")
	if MongoURI == "" {
		MongoURI = "mongodb://localhost:27020/test"
	}
	return err
}
