package main

import (
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fsantiag/track-progress/src/database"
	"github.com/fsantiag/track-progress/src/queue"
	"github.com/fsantiag/track-progress/src/repository"
	"github.com/fsantiag/track-progress/src/server"
	"github.com/fsantiag/track-progress/src/service"
	"github.com/fsantiag/track-progress/src/util"
	"github.com/sirupsen/logrus"
)

var logger logrus.Logger

func init() {
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
}

func main() {
	session := setupDatabase()
	startGoroutineListeners(session)
	initServer()
}

func setupDatabase() repository.SessionInterface {
	session, err := database.NewSession()
	if err != nil {
		logrus.Fatal("Database connection error:", err.Error())
	}
	defer session.Close()
	database.Migrate(session)
	return session
}

func startGoroutineListeners(session repository.SessionInterface) {
	channel := make(chan *sqs.Message, 100)
	connection := queue.NewSession()
	queueURL := queue.CreateQueues(connection)
	go queue.Poll(channel, queueURL, connection, &logger)
	go service.ProcessTaskMessage(session, repository.TaskRepository{}, channel, &logger)
}

func initServer() {
	s := server.InitRouter()
	logger.Println("Server started...")
	logger.Fatal(http.ListenAndServe(server.Address[util.Getenv("PROFILE_ENV", "dev")], s))
}
