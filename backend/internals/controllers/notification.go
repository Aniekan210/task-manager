package controls

import (
	"context"

	"github.com/Aniekan210/taskManager/backend/internals/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNotif(teamID primitive.ObjectID, notifBody string) error {

	newNotif := models.Notification{
		ID:     primitive.NewObjectID(),
		Body:   notifBody,
		TeamID: teamID,
	}

	// Add user to database
	collection := Client.Database(DBName).Collection("notifs")
	_, err := collection.InsertOne(context.Background(), newNotif)
	if err != nil {
		return err
	}

	return nil
}
