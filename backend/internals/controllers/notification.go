package controls

import (
	"fmt"

	"github.com/Aniekan210/taskManager/backend/internals/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNotif(username string, teamID primitive.ObjectID, notifBody string) error {

	newNotif := models.Notification{
		ID:        primitive.NewObjectID(),
		Body:      notifBody,
		CreatedBy: username,
		TeamID:    teamID,
	}

	fmt.Println(newNotif)
	return nil
}
