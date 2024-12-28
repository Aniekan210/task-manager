package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id"`         //id of task
	CreatedBy   string             `bson:"user_id"`     // Username of user that created task
	AssignedTo  []string           `bson:"assigned_to"` // List of usernames. Users assigned to task
	Title       string             `bson:"title"`       // Title of task
	Description string             `bson:"description"` // Description of task
	DueDate     time.Time          `bson:"due_date"`    // When task is due
	Status      string             `bson:"status"`      // "in-progress", "completed", "overdue"
}
