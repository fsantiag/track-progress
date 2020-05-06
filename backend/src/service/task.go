package service

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fsantiag/track-progress/backend/src/model"
	"github.com/fsantiag/track-progress/backend/src/repository"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
	repo   repository.Repository
)

// TaskService represents a instance to process task message
type TaskService struct{}

// NewTaskService create a new instance of TaskService with your arguments
func NewTaskService(loggerLogrus *logrus.Logger, repository repository.Repository) TaskService {
	logger = loggerLogrus
	repo = repository
	return TaskService{}
}

// ProcessTaskMessage receive a message by channel to save task
func (ts TaskService) ProcessTaskMessage(channel <-chan *sqs.Message) {
	for {
		ts.process(<- channel)
	}
}

func (ts TaskService) process(message *sqs.Message) {
	logger.Info("Got this message", message)
	task := model.Task{}
	err := json.Unmarshal([]byte(*message.Body), &task)
	if err != nil {
		logger.Error("Could not unmarshal task: ", err.Error())
		return
	}

	err = repo.Save(task)
	if err != nil {
		logger.Error("Fail to persist task: ", err.Error())
		return
	}
}
