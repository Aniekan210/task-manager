package routes

import (
	"errors"
	"net/http"
	"strings"

	controls "github.com/Aniekan210/taskManager/backend/internals/controllers"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", registerUser)
		auth.POST("/login", userLogin)
	}
}

func registerUser(ctx *gin.Context) {

	type request struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validating if username and password do not have spaces
	if (strings.Contains(req.Password, " ")) || (strings.Contains(req.Username, " ")) {
		err = errors.New("password and username cannot contain spaces")
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validating length of username and password
	if (len([]rune(req.Password)) < 8) || (len([]rune(req.Username)) < 8) {
		err = errors.New("password and username cannot be less than 8 characters")
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check if user exists already
	user, err := controls.FindUserByUsername(req.Username)
	if (err != nil) && (err.Error() != "user not found") {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if user != nil {
		err = errors.New("user already exists")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Add the user to database
	err = controls.AddUser(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "created succesfully",
	})
}

func userLogin(ctx *gin.Context) {

	type request struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Get request body
	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Strip leading and trailing whitespace from password and username
	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)

	// Find user in database
	user, err := controls.FindUserByUsername(req.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(req.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Password",
		})
		return
	}

	// Generate JWT
	jwt, err := controls.GenerateJWT(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":   jwt,
		"message": "logged in succesfully",
	})
}
