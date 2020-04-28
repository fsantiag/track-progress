package service

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fsantiag/track-progress/src/model"
	"github.com/fsantiag/track-progress/src/repository"
	"github.com/sirupsen/logrus"
)

// ProcessTaskMessage receive a message by channel to save task
func ProcessTaskMessage(session repository.SessionInterface, repository repository.Repository, channel <-chan *sqs.Message, logger *logrus.Logger) {
	for {
		process(session, repository, channel, logger)
	}
}

func process(session repository.SessionInterface, repository repository.Repository, channel <-chan *sqs.Message, logger *logrus.Logger) {
	message := <-channel

	logger.Info("Got this message", message)
	task := model.Task{}
	err := json.Unmarshal([]byte(*message.Body), &task)
	if err != nil {
		logger.Error("Could not unmarshal task: ", err.Error())
		return
	}

	err = repository.Save(session, task)
	if err != nil {
		logger.Error("Fail to persist task: ", err.Error())
		return
	}
}
