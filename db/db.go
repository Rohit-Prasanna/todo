package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectDB establishes a connection to the MongoDB database
func ConnectDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://Rohit-todo:1YWXMEgBEy8VKaTb@cluster0.l3ra7.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	return err
}

// DisconnectDB closes the MongoDB connection
func DisconnectDB() {
	if Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := Client.Disconnect(ctx)
		if err != nil {
			return
		}
	}
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("todoapp").Collection(collectionName)
}
