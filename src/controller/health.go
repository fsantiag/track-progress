package controller

import (
	"fmt"
	"net/http"
)

// Health returns if the service is up or down
func Health(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	fmt.Fprintf(response, "UP!")
}
