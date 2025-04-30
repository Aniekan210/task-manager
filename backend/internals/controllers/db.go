package controls

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DBName string

func Init() error {
	mongoURI := os.Getenv("MONGODB_URI")
	opts := options.Client().ApplyURI(mongoURI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	Client = client
	DBName = os.Getenv("DB_NAME")
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
