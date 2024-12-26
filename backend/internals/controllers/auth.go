package controls

import (
	"errors"
	"os"
	"time"

	"github.com/Aniekan210/taskManager/backend/internals/models"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(user *models.User) (string, error) {

	// Create a new claims object with usermodels
	claims := jwt.MapClaims{
		"username": user.Username,                         // Add username
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Expiration time (3 days)
		"iat":      time.Now().Unix(),                     // Issued at time
	}

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Ensure the signing method is correct
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract the claims (payload) if the token is valid
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
