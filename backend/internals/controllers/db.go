package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
	DBname string
}

func (d *Database) Init() error {
	mongoURI := "mongodb+srv://aniekan:DUwKNrxGtBJFoxz7@aniekan.qieqv.mongodb.net/?retryWrites=true&w=majority&appName=aniekan"
	opts := options.Client().ApplyURI(mongoURI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	d.Client = client
	return nil
}

func (d *Database) Disconnect() error {
	if d.Client != nil {
		err := d.Client.Disconnect(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}
