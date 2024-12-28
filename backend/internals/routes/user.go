package routes

import (
	"errors"
	"net/http"
	"strings"

	controls "github.com/Aniekan210/taskManager/backend/internals/controllers"
	"github.com/Aniekan210/taskManager/backend/internals/middleware"
	"github.com/Aniekan210/taskManager/backend/internals/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUserRoutes(router *gin.Engine) {
	user := router.Group("/user")
	user.Use(middleware.Authentication())
	{
		user.GET("/", getUser)
		user.POST("/create-team", createTeam)
		user.POST("/join-team", joinTeam)

		user.POST("/create-project", createProject)
	}
}

func getUser(ctx *gin.Context) {

	type response struct {
		Username string
		Teams    []models.TeamInfo
	}

	var res response

	// Get username from claims
	claims, _ := ctx.Get("claims")
	username := controls.GetUsernameFromClaims(claims)

	// get user by username
	user, err := controls.FindUserByUsername(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// fill response body
	res.Teams = user.Teams
	res.Username = user.Username

	ctx.JSON(http.StatusOK, gin.H{
		"user": res,
	})
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
	username := controls.GetUsernameFromClaims(claims)

	teamID, err := controls.CreateTeam(username, req.TeamName, req.TeamDescription)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controls.AddUserToTeam(username, teamID)
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

func joinTeam(ctx *gin.Context) {

	type request struct {
		JoinCode string `json:"join_code" binding:"required"`
	}

	var req request

	//bind the json
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// get username from claims
	claims, _ := ctx.Get("claims")
	username := controls.GetUsernameFromClaims(claims)

	//Get team
	team, err := controls.FindTeamByJoinCode(req.JoinCode)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Add team to user
	err = controls.AddToUserTeamInfo(username, team.ID, "editor")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// add user to team
	err = controls.AddUserToTeam(username, team.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "joined team succesfully",
	})
}

func createProject(ctx *gin.Context) {

	type request struct {
		TeamID             string `json:"team_id" binding:"required"`
		ProjectName        string `json:"project_name" binding:"required"`
		ProjectDescription string `json:"project_description" binding:"required"`
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
	req.ProjectName = strings.TrimSpace(req.ProjectName)
	req.ProjectDescription = strings.TrimSpace(req.ProjectDescription)

	//validate team name and description
	if (len([]rune(req.ProjectDescription)) < 10) || (len([]rune(req.ProjectDescription)) > 200) {
		err = errors.New("project description must be between 10 and 200 characters")
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Get username from claims
	claims, _ := ctx.Get("claims")
	username := controls.GetUsernameFromClaims(claims)

	// create project
	projectID, err := controls.CreateProject(username, req.ProjectName, req.ProjectDescription)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// get team id
	teamID, err := primitive.ObjectIDFromHex(req.TeamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// add project to team
	err = controls.AddProjectToTeam(teamID, projectID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "project created successfully",
	})
}
