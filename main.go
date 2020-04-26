package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fsantiag/track-progress/src/database"
	"github.com/fsantiag/track-progress/src/queue"
	"github.com/fsantiag/track-progress/src/server"
	"github.com/fsantiag/track-progress/src/service"
)

func main() {
	// Setup the database
	session, _ := database.NewSession()
	defer session.Close()
	database.Migrate(session)

	// Start the goroutine listeners
	channel := make(chan *sqs.Message)
	connection := queue.NewSession()
	queue.CreateQueues(connection)
	go queue.Poll(channel, connection)
	go service.ProcessTaskMessage(channel)

	// Init the server
	s := server.InitRouter()
	log.Println("Server started...")
	log.Fatal(http.ListenAndServe(":8080", s))
}
