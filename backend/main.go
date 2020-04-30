package main

import (
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fsantiag/track-progress/backend/src/database"
	"github.com/fsantiag/track-progress/backend/src/queue"
	"github.com/fsantiag/track-progress/backend/src/repository"
	"github.com/fsantiag/track-progress/backend/src/server"
	"github.com/fsantiag/track-progress/backend/src/service"
	"github.com/fsantiag/track-progress/backend/src/util"
	"github.com/sirupsen/logrus"
)

var logger logrus.Logger

func init() {
	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
}

func main() {
	databaseSession := setupDatabase()
	defer databaseSession.Close()
	startGoroutineListeners(databaseSession)
	initServer()
}

func setupDatabase() repository.SessionInterface {
	session, err := database.NewSession(&logger)
	if err != nil {
		logrus.Fatal("Database connection error: ", err.Error())
	}
	database.Migrate(session)
	return session
}

func startGoroutineListeners(databaseSession repository.SessionInterface) {
	channel := make(chan *sqs.Message, 100)
	sqsSession := queue.NewSession()
	queueURL, err := queue.CreateQueues(sqsSession)
	if err != nil {
		logger.Fatal("Queues creator error: ", err.Error())
	}
	go queue.Poll(channel, queueURL, sqsSession, &logger)

	repository := repository.NewTaskRepository(&logger, databaseSession)
	service := service.NewTaskService(&logger, repository)
	go service.ProcessTaskMessage(channel)
}

func initServer() {
	s := server.InitRouter()
	logger.Info("Server started...")
	env := util.Getenv("PROFILE_ENV", "dev")
	logger.Debug("Environment: ", env)
	logger.Fatal(http.ListenAndServe(server.Address[env], s))
}
