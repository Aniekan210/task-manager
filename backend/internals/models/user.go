package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id"`      // User id
	Teams          []TeamInfo         `bson:"teams"`    // List of team ids and roles in the teams
	Username       string             `bson:"username"` // username of the user
	Email          string             `bson:"email"`
	HashedPassword []byte             `bson:"hashed_password"` // Hashed password of the user
	IsVerified     bool               `bson:"isVerified"`
}

type TeamInfo struct {
	ID   primitive.ObjectID `bson:"team_id"` // Id of team
	Role string             `bson:"role"`    // "creator", "editor", "viewer"
}
