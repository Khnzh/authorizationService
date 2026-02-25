package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	res, err := json.Marshal(payload)
	if err != nil {
		RespondWithError(w, 500, fmt.Sprintf("%v", err))
	}
	_, err = w.Write(res)
	if err != nil {
		fmt.Println("Error writing to response", err)
	}
}

func RespondWithError(w http.ResponseWriter, code int, payload string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	res, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error occured while marshaling err:", err)
	}
	_, err = w.Write(res)
	if err != nil {
		fmt.Println("Error writing to response", err)
	}
}
