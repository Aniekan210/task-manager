package db

import (
	"context"
	"errors"

	"github.com/Aniekan210/taskManager/backend/internals/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(username string, password string) error {

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create user object
	user := models.User{
		ID:             primitive.NewObjectID(),
		Username:       username,
		HashedPassword: hashedPassword,
		Teams:          []models.TeamInfo{},
	}

	// Add user to database
	collection := Client.Database(DBName).Collection("users")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

// returns false if user doesnt exist
func FindUserByUsername(username string) (*models.User, error) {

	filter := bson.M{"username": username}

	// Get the collection
	collection := Client.Database(DBName).Collection("users")

	// Find a single user by the filter
	var user models.User
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			// Return a custom error when the user is not found
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
