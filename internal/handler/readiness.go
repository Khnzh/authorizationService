package handlers

import (
	"fmt"
	"net/http"

	"example.com/authorizationService/internal/utils"
)

func (api *ServiceApis) PingDatabases(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := api.DB.PingContext(ctx)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Error connecting to db:%v", err))
		return
	}

	rdbStatus := api.Redis.Ping(ctx)
	response := map[string]string{
		"Connection to postgres database": "success",
		"Connection to redis database":    rdbStatus.String(),
	}
	utils.RespondWithJSON(w, 201, response)
}
