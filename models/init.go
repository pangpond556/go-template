package models

import (
	"appname/utils"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client - mongodb client
var Client *mongo.Client

// Result - response format
var Result = map[string]interface{}{
	"message": nil,
	"data":    nil,
}

// InitDatabase - init database
func InitDatabase() error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options := options.Client().ApplyURI(utils.MongoConnect)
	Client, err = mongo.Connect(ctx, options)
	if err != nil {
		return err
	}
	fmt.Println("Database connected.")
	return nil
}
