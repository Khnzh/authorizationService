package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/authorizationService/db"
	handlers "example.com/authorizationService/internal/handler"
	"example.com/authorizationService/internal/models"
	"github.com/joho/godotenv"

	// "example.com/authorizationService/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
)

type connection interface {
	Close() error
}

func closeConns(c connection) {
	if error := c.Close(); error != nil {
		panic("Connection is not closed")
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	APIConnections := &models.ServiceApis{}
	APIConnections.DB = db.New(fmt.Sprintf("postgres://%v:%v@%v:%v/authServ?sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT")))

	APIConnections.Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	api := handlers.NewServiceApis(APIConnections)

	defer closeConns(APIConnections.DB)
	defer closeConns(APIConnections.Redis)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", api.HandleCreateUser)
	r.Post("/login", api.HandleFetchUser)
	http.ListenAndServe(":3000", r)
}
