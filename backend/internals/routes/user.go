package routes

import (
	"errors"
	"net/http"
	"strings"

	controls "github.com/Aniekan210/taskManager/backend/internals/controllers"
	"github.com/Aniekan210/taskManager/backend/internals/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RegisterUserRoutes(router *gin.Engine) {
	auth := router.Group("/user")
	auth.Use(middleware.Authentication())
	{
		auth.POST("/create-team", createTeam)
	}
}

func createTeam(ctx *gin.Context) {

	type request struct {
		TeamName        string `json:"team_name" binding:"required"`
		TeamDescription string `json:"team_description" binding:"required"`
	}

	var req request

	// Get request body
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Strip team name and description
	req.TeamName = strings.TrimSpace(req.TeamName)
	req.TeamDescription = strings.TrimSpace(req.TeamDescription)

	//validate team name and description
	if (len([]rune(req.TeamDescription)) < 10) || (len([]rune(req.TeamDescription)) > 200) {
		err = errors.New("team description must be between 10 and 200 characters")
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Get username from claims
	claims, _ := ctx.Get("claims")
	jwtClaims, _ := claims.(jwt.MapClaims)

	name := jwtClaims["username"]
	username, _ := name.(string)

	teamID, err := controls.CreateTeam(username, req.TeamName, req.TeamDescription)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controls.AddToUserTeamInfo(username, teamID, "creator")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "team created successfully",
	})
}
