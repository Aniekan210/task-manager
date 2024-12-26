package main

import (
	"log"

	controls "github.com/Aniekan210/taskManager/backend/internals/controllers"
	"github.com/Aniekan210/taskManager/backend/internals/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	// Load environment variables from .env file
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Initialize Database
	err = controls.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		err := controls.Disconnect()
		if err != nil {
			log.Fatal("Error disconnecting from the database: ", err)
		}
	}()

	routes.RegisterAuthRoutes(r)
	routes.RegisterUserRoutes(r)

	r.Run()
}
