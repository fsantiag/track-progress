package controller

import (
	"encoding/json"
	"net/http"
)

type health struct {
	Status string `json:"status"`
}

// Health returns if the service is up or down
func Health(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)

	health := health{"Up!"}

	responseJSON, _ := json.Marshal(health)

	response.Write(responseJSON)
}
