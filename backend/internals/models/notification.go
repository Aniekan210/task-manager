package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID        primitive.ObjectID `bson:"_id"`
	Body      string             `bson:"notif_body"`
	CreatedBy string             `bson:"created_by"`
	TeamID    primitive.ObjectID `bson:"team_id"`
}
