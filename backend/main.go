package main

import (
	"net/http"
	"os"

	awssqs "github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fsantiag/track-progress/backend/src/cassandra"
	"github.com/fsantiag/track-progress/backend/src/repository"
	"github.com/fsantiag/track-progress/backend/src/server"
	"github.com/fsantiag/track-progress/backend/src/service"
	"github.com/fsantiag/track-progress/backend/src/sqs"
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
	session, err := cassandra.NewSession(&logger)
	if err != nil {
		logrus.Fatal("Database connection error: ", err.Error())
	}
	cassandra.Migrate(session)
	return session
}

func startGoroutineListeners(databaseSession repository.SessionInterface) {
	sqsSession := sqs.NewSession()
	queueURL, err := sqs.CreateQueues(sqsSession)
	if err != nil {
		logger.Fatal("Queues creator error: ", err.Error())
	}

	channel := make(chan *awssqs.Message, 100)
	go sqs.Poll(channel, queueURL, sqsSession, &logger)

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
