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
	session, err := database.NewSession()
	if err != nil {
		log.Fatal("Database connection error:", err.Error())
	}
	defer session.Close()
	database.Migrate(session)

	// Start the goroutine listeners
	channel := make(chan *sqs.Message, 100)
	connection := queue.NewSession()
	queueURL := queue.CreateQueues(connection)
	go queue.Poll(channel, queueURL, connection)
	go service.ProcessTaskMessage(channel)

	// Init the server
	s := server.InitRouter()
	log.Println("Server started...")
	log.Fatal(http.ListenAndServe(":8080", s))
}
