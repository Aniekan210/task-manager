package controls

import (
	"context"
	"errors"

	"github.com/Aniekan210/taskManager/backend/internals/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateTeam(username string, teamName string, teamDescription string) (primitive.ObjectID, error) {

	id := primitive.NewObjectID()

	newTeam := models.Team{
		ID:          id,
		Name:        teamName,
		Description: teamDescription,
		CreatedBy:   username,
		Projects:    []primitive.ObjectID{},
		JoinCode:    id.Hex()[4:8] + id.Hex()[20:],
	}

	// Add user to database
	collection := Client.Database(DBName).Collection("teams")
	_, err := collection.InsertOne(context.Background(), newTeam)
	if err != nil {
		return primitive.NilObjectID, err
	}

	err = AddToUserTeamInfo(username, newTeam.ID, "creator")
	if err != nil {
		return primitive.NilObjectID, err
	}

	return newTeam.ID, nil
}

func AddUserToTeam(username string, teamID primitive.ObjectID) error {

	// Get the team
	team, err := FindTeamByID(teamID)
	if err != nil {
		return err
	}

	// Check if user is already in team
	for _, user := range team.Users {
		if user == username {
			return errors.New("user is already in team")
		}
	}

	var newUsers []string = append(team.Users, username)

	// Get the collection
	collection := Client.Database(DBName).Collection("teams")
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "_id", Value: teamID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "users", Value: newUsers}}}}

	_, err = collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}

func FindTeamByID(teamID primitive.ObjectID) (*models.Team, error) {

	filter := bson.M{"_id": teamID}

	// Get the collection
	collection := Client.Database(DBName).Collection("teams")

	// Find a single team by the filter
	var team models.Team
	err := collection.FindOne(context.Background(), filter).Decode(&team)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func FindTeamByJoinCode(joinCode string) (*models.Team, error) {
	filter := bson.M{"join_code": joinCode}

	// Get the collection
	collection := Client.Database(DBName).Collection("teams")

	// Find a single team by the filter
	var team models.Team
	err := collection.FindOne(context.Background(), filter).Decode(&team)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			// Return a custom error when the user is not found
			return nil, errors.New("invalid join code")
		}
		return nil, err
	}

	return &team, nil
}

func JoinTeam(username string, joinCode string) error {
	//Get team
	team, err := FindTeamByJoinCode(joinCode)
	if err != nil {
		return err
	}

	// Add team to user
	err = AddToUserTeamInfo(username, team.ID, "editor")
	if err != nil {
		return err
	}

	// add user to team
	err = AddUserToTeam(username, team.ID)
	if err != nil {
		return err
	}

	return nil
}
