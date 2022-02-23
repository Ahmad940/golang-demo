package util

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"time"
)

func GenerateToken(id string) (string, error) {
	// Create the Claims
	hour, err := strconv.Atoi(os.Getenv("JWT_DURATION"))
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(hour)).Unix(),
		"id":  id,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	encodedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return encodedToken, nil
}
