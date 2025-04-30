package routes

import (
	"errors"
	"fmt"
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
	email, _ := controls.ParseClaims(claims)

	// get user by username
	user, err := controls.FindUserByEmail(email)
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
	email, username := controls.ParseClaims(claims)

	teamID, err := controls.CreateTeam(email, username, req.TeamName, req.TeamDescription)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controls.AddUserToTeam(email, teamID)
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

	// get email and username from claims
	claims, _ := ctx.Get("claims")
	email, username := controls.ParseClaims(claims)

	//Get team
	team, err := controls.FindTeamByJoinCode(req.JoinCode)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Add team to user
	err = controls.AddToUserTeamInfo(email, team.ID, "editor")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// add user to team
	err = controls.AddUserToTeam(email, team.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// send notif to team
	body := fmt.Sprintf("%s has joined Team: %s", username, team.Name)
	err = controls.CreateNotif(team.ID, body)
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

	//Strip project name and description
	req.ProjectName = strings.TrimSpace(req.ProjectName)
	req.ProjectDescription = strings.TrimSpace(req.ProjectDescription)

	//validate project name and description
	if (len([]rune(req.ProjectDescription)) < 10) || (len([]rune(req.ProjectDescription)) > 200) {
		err = errors.New("project description must be between 10 and 200 characters")
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Get email from claims
	claims, _ := ctx.Get("claims")
	_, username := controls.ParseClaims(claims)

	// get team id
	teamID, err := primitive.ObjectIDFromHex(req.TeamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Team id",
		})
		return
	}

	// create project
	projectID, err := controls.CreateProject(username, req.ProjectName, req.ProjectDescription)
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

	// send notif to team
	body := fmt.Sprintf("%s has created a new Project: %s", username, req.ProjectName)
	err = controls.CreateNotif(teamID, body)
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
