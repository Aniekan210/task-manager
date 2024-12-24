package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID    primitive.ObjectID   `bson:"_id"`
	Tasks []primitive.ObjectID `bson:"tasks"` // References to task IDs within this team
}
