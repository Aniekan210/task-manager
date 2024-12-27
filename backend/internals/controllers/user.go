package controls

import (
	"context"
	"errors"

	"github.com/Aniekan210/taskManager/backend/internals/models"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func GetUsernameFromClaims(claims any) string {
	jwtClaims, _ := claims.(jwt.MapClaims)

	name := jwtClaims["username"]
	username, _ := name.(string)

	return username
}

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

func AddToUserTeamInfo(username string, teamID primitive.ObjectID, role string) error {
	user, err := FindUserByUsername(username)
	if err != nil {
		return err
	}

	// Check if user is in team
	for _, team := range user.Teams {
		if team.ID == teamID {
			return errors.New("user is already in team")
		}
	}

	newTeams := append(user.Teams, models.TeamInfo{
		ID:   teamID,
		Role: role,
	})

	// Get the collection
	collection := Client.Database(DBName).Collection("users")
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "_id", Value: user.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "teams", Value: newTeams}}}}

	_, err = collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}
