package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID          primitive.ObjectID   `bson:"_id"`         // Team id
	Name        string               `bson:"name"`        // Team name
	Description string               `bson:"description"` // Team description
	CreatedBy   primitive.ObjectID   `bson:"created_by"`  // Creator of the team
	Projects    []primitive.ObjectID `bson:"project_ids"` // List of team ids and roles in the teams
	JoinCode    string               `bson:"join_code"`   // Code used to add users to a team
}
