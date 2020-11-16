package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitDatabase - init database
func InitDatabase() (*mongo.Client, context.Context, context.CancelFunc, error) {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	options := options.Client().ApplyURI(MongoConnect)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, nil, cancel, err
	}
	return client, ctx, cancel, nil
}
