package main

import (
	"log"

	db "github.com/Aniekan210/taskManager/backend/internals/controllers"
	"github.com/Aniekan210/taskManager/backend/internals/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize Database
	err := db.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		if err := db.Disconnect(); err != nil {
			log.Fatal("Error disconnecting from the database: ", err)
		}
	}()

	routes.RegisterUserRoutes(r)

	r.Run()
}
