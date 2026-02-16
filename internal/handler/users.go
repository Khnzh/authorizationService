package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/authorizationService/internal/database"
	"example.com/authorizationService/internal/models"
	"example.com/authorizationService/internal/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ServiceApis struct {
	*models.ServiceApis // Embedding
}

func NewServiceApis(original *models.ServiceApis) *ServiceApis {
	return &ServiceApis{ServiceApis: original}
}

func (api *ServiceApis) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing params:%v", err))
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Error while hashing password:%v", err))
		return
	}

	user, err := database.New(api.DB).CreateUser(r.Context(), database.CreateUserParams{
		ID:           id,
		Name:         params.Name,
		Email:        params.Email,
		PasswordHash: string(hash),
		Role:         params.Role,
	})

	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Error creating user:%v", err))
		return
	}

	utils.RespondWithJSON(w, 201, models.DatabaseUserToStruct(user))
}

func (api *ServiceApis) HandleFetchUser(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing params:%v", err))
		return
	}

	user, err := database.New(api.DB).FetchUser(r.Context(), params.Email)
	if err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, 400, "Invalid credentials")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(params.Password))
	if err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, 400, "Invalid credentials")
		return
	}

	utils.RespondWithJSON(w, 200, "Logged in successfully!")
}
