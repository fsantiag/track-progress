package main

import (
	"log"
	"net/http"

	"github.com/fsantiag/track-progress/src/server"
)

func main() {
	server := server.InitRoutes()
	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", server))
}
