package main

import (
	"log"
	"net/http"

	"github.com/fsantiag/track-progress/src/configuration"
	"github.com/fsantiag/track-progress/src/server"
)

func main() {
	configuration.Migrate()
	s := server.InitRouter()
	log.Println("Server started...")
	log.Fatal(http.ListenAndServe(":8080", s))
}
