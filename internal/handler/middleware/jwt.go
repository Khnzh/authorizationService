package cust_middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"example.com/authorizationService/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type authHandler func(http.ResponseWriter, *http.Request, models.User)

func AuthCheck(f authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		testVal := r.Header.Get("Authorization")
		tokenVal := strings.Split(testVal, " ")[1]
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error parsing env")
		}
		key := os.Getenv("KEY")
		user := models.User{}
		token, err := jwt.ParseWithClaims(tokenVal, &user, func(t *jwt.Token) (interface{}, error) {
			if key == "" {
				return nil, errors.New("key not passed")
			}
			return []byte(key), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		fmt.Println(token.Valid)
		fmt.Println(user)
		// tokenParser := jwt.Parser{}
		// _, payload, err := tokenParser.ParseUnverified(testVal, user)
		// token, err := jwt.Parse(testVal, func(t *jwt.Token) (interface{}, error) {
		// 	if key == "" {
		// 		return nil, errors.New("key not passed")
		// 	}
		// 	return []byte(key), nil
		// }, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		// fmt.Println(err)
		// fmt.Println(token.Header, "middleware/jwt.go")
		// cookie, err := r.Cookie("ref")
		// fmt.Println(cookie.Value)
		// http.SetCookie(w, &http.Cookie{
		// 	Name:     "ref",
		// 	Value:    tokenVal,
		// 	Path:     "/",
		// 	HttpOnly: true,
		// 	Secure:   true,
		// 	SameSite: http.SameSiteStrictMode,
		// 	Expires:  time.Now().Add(time.Hour * 24 * 7),
		// })
		f(w, r, user)
	}
}
