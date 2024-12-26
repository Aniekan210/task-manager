package controls

import (
	"context"

	"github.com/Aniekan210/taskManager/backend/internals/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTeam(username string, teamName string, teamDescription string) (primitive.ObjectID, error) {

	id := primitive.NewObjectID()

	newTeam := models.Team{
		ID:          id,
		Name:        teamName,
		Description: teamDescription,
		CreatedBy:   username,
		Projects:    []primitive.ObjectID{},
		JoinCode:    id.Hex()[:6] + id.Hex()[22:],
	}

	// Add user to database
	collection := Client.Database(DBName).Collection("teams")
	_, err := collection.InsertOne(context.Background(), newTeam)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return newTeam.ID, nil
}
