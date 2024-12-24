package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	TeamIDs      []TeamInfo         `bson:"team_ids"`
	Username     string             `bson:"username"`
	PasswordHash string             `bson:"passwordHash"`
}

type TeamInfo struct {
	TeamID primitive.ObjectID `bson:"_id"`
	Role   string             `bson:"role"`
}
