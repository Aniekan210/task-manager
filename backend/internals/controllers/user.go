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

func ParseClaims(claims any) (string, string) {
	jwtClaims, _ := claims.(jwt.MapClaims)

	mail := jwtClaims["email"]
	email, _ := mail.(string)

	name := jwtClaims["userename"]
	username, _ := name.(string)

	return email, username
}

func AddUser(username string, password string, email string) error {

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create user object
	user := models.User{
		ID:             primitive.NewObjectID(),
		Username:       username,
		Email:          email,
		HashedPassword: hashedPassword,
		Teams:          []models.TeamInfo{},
		IsVerified:     false,
	}

	// Add user to database
	collection := Client.Database(DBName).Collection("users")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func FindUserByEmail(email string) (*models.User, error) {

	filter := bson.M{"email": email}

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

func FindUserByID(id string) (*models.User, error) {

	// get team id
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	filter := bson.M{"_id": ID}

	// Get the collection
	collection := Client.Database(DBName).Collection("users")

	// Find a single user by the filter
	var user models.User
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			// Return a custom error when the user is not found
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func AddToUserTeamInfo(email string, teamID primitive.ObjectID, role string) error {
	user, err := FindUserByEmail(email)
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

func VerifyUser(id string) error {
	user, err := FindUserByID(id)
	if err != nil {
		return err
	}

	// Check if user is already verified
	if user.IsVerified {
		return errors.New("user is already verified")
	}

	// Get the collection
	collection := Client.Database(DBName).Collection("users")
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "_id", Value: user.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "isVerified", Value: true}}}}

	_, err = collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}
