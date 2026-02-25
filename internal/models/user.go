package models

import (
	"database/sql"
	"errors"
	"time"

	"example.com/authorizationService/internal/database"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type AccToken struct {
	ID    uuid.UUID `json:"id" bd:"email"`
	Name  string    `json:"name" bd:"name"`
	Email string    `json:"email" bd:"email"`
	Role  string    `json:"role" bd:"role"`
	jwt.RegisteredClaims
}

type User struct {
	ID        uuid.UUID    `json:"id" bd:"email"`
	Name      string       `json:"name" bd:"name"`
	Email     string       `json:"email" bd:"email"`
	Role      string       `json:"role" bd:"role"`
	IsActive  sql.NullBool `json:"is_active" bd:"is_active"`
	CreatedAt time.Time    `json:"created_at" bd:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" bd:"updated_at"`
}

func (u User) Valid() error {
	if u.Email != "" {
		return nil
	}
	return errors.New("User not valid")
}

func DatabaseUserToStruct(u database.User) User {
	return User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		IsActive:  u.IsActive,
		Role:      u.Role,
	}
}
