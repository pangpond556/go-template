package utils

import (
	"math/rand"
	"os"
	// "github.com/joho/godotenv"
)

// MongoConnect - mongodb connection string
var MongoConnect string

// Config - config env
func Config() error {
	var err error
	// for develop using file .env
	// err = godotenv.Load()
	// if err != nil {
	// 	return err
	// }

	MongoConnect = os.Getenv("MONGO_CONNECT")
	if MongoConnect == "" {
		MongoConnect = "mongodb://localhost:27020/test"
	}
	return err
}

// RandomString - random character by n number
func RandomString(n int) string {
	var characters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]byte, n)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return string(b)
}
