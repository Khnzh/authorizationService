package utils

import (
	"os"

	"example.com/authorizationService/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func GenerateToken(user models.User) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	key := os.Getenv("KEY")
	// t := jwt.New(jwt.SigningMethodHS256)
	tc := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	s, err := tc.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return s, nil
}
