package main

import (
	"log"

	db "github.com/Aniekan210/taskManager/backend/internals/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	r := gin.Default()

	// Initialize Database Contoller
	DB := &db.Database{
		Client: &mongo.Client{},
		DBname: "aniekan",
	}

	// Initialize Database
	err := DB.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		if err := DB.Disconnect(); err != nil {
			log.Fatal("Error disconnecting from the database: ", err)
		}
	}()
	r.Run()
}
