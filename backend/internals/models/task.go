package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID   `bson:"_id"`         //id of task
	CreatedBy   primitive.ObjectID   `bson:"user_id"`     // User id of user that created task
	AssignedTo  []primitive.ObjectID `bson:"assigned_to"` // List of user ids. Users assigned to task
	Title       string               `bson:"title"`       // Title of task
	Description string               `bson:"description"` // Description of task
	DueDate     time.Time            `bson:"due_date"`    // When task is due
	Status      string               `bson:"status"`      // "in-progress", "completed", "overdue"
}
