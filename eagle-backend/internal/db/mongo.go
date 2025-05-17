package db

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance     *mongo.Client
	clientInstanceErr  error
	mongoOnce          sync.Once
)

const dbName = "eagle"

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		uri := os.Getenv("MONGODB_URI")
		if uri == "" {
			uri = "mongodb://localhost:27017/eagle"
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			clientInstanceErr = err
			return
		}
		if err = client.Ping(ctx, nil); err != nil {
			clientInstanceErr = err
			return
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceErr
}

func GetUserCollection() *mongo.Collection {
	client, err := GetMongoClient()
	if err != nil {
		log.Fatalf("Mongo connection error: %v", err)
	}
	return client.Database(dbName).Collection("users")
}
