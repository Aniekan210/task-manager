package controls

import (
	"context"

	"github.com/Aniekan210/taskManager/backend/internals/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProject(username string, projectName string, projectDescription string) (primitive.ObjectID, error) {

	newProject := models.Project{
		ID:          primitive.NewObjectID(),
		Name:        projectName,
		Description: projectDescription,
		CreatedBy:   username,
		Tasks:       []primitive.ObjectID{},
	}

	// Add project to database
	collection := Client.Database(DBName).Collection("projects")
	_, err := collection.InsertOne(context.Background(), newProject)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return newProject.ID, nil
}
