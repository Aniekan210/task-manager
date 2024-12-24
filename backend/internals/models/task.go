package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID   `bson:"_id"`
	CreatedBy   primitive.ObjectID   `bson:"user_id"`
	AssignedTo  []primitive.ObjectID `bson:"assigned_to"` // Multiple users can be assigned to a task
	Title       string               `bson:"title"`
	Description string               `bson:"description"`
	DueDate     time.Time            `bson:"due_date"`
	Status      string               `bson:"status"` // e.g., "pending", "in-progress", "completed"
}
