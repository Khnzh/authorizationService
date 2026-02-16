package models

import (
	// "github.com/go-playground/validator/v10"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email" bd:"email" validate:"required,email"`
	Password string `json:"password" bd:"-" validate:"required,min=8"`
}

type AuthResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	User         string    `json:"user"`
	ExpiresAt    time.Time `json:"time"`
}
