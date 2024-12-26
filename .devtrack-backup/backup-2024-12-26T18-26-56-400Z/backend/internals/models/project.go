package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID   `bson:"_id"`         // Id of the project
	Name        string               `bson:"name"`        // Project name
	Description string               `bson:"description"` // Project description
	Tasks       []primitive.ObjectID `bson:"task_ids"`    // List of task ids in the project
	CreatedBy   primitive.ObjectID   `bson:"created_by"`  // Creator of the project
}
