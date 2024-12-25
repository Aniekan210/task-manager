package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DBName string

func Init() error {
	mongoURI := "mongodb+srv://aniekan:DUwKNrxGtBJFoxz7@aniekan.qieqv.mongodb.net/?retryWrites=true&w=majority&appName=aniekan"
	opts := options.Client().ApplyURI(mongoURI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	Client = client
	DBName = "aniekan"
	return nil
}

func Disconnect() error {
	if Client != nil {
		err := Client.Disconnect(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}
