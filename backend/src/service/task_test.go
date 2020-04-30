package service

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fsantiag/track-progress/backend/src/model"
	"github.com/fsantiag/track-progress/backend/src/repository"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
	repository.Repository
}

func (c *mockRepository) Save(task model.Task) (err error) {
	args := c.Called(task)
	return args.Error(0)
}

func TestProcessTask(t *testing.T) {
	logger, _ := test.NewNullLogger()
	repository := mockRepository{}
	channel := make(chan *sqs.Message, 1)
	body := `{"title":"any title","description":"any description","status":"any status"}`
	message := sqs.Message{Body: &body}
	channel <- &message
	task := model.Task{}
	json.Unmarshal([]byte(body), &task)

	repository.On("Save", task).Return(nil)

	service := NewTaskService(logger, &repository)
	service.process(channel)

	repository.AssertNumberOfCalls(t, "Save", 1)
	repository.AssertExpectations(t)
}

func TestProcessTaskWithErrorToSaveTask(t *testing.T) {
	logger, hook := test.NewNullLogger()
	repository := mockRepository{}
	channel := make(chan *sqs.Message, 1)
	newError := errors.New("error to persist")

	body := `{"title":"any title","description":"any description","status":"any status"}`
	message := sqs.Message{Body: &body}
	channel <- &message

	task := model.Task{}
	json.Unmarshal([]byte(body), &task)
	repository.On("Save", task).Return(newError)

	service := NewTaskService(logger, &repository)
	service.process(channel)

	assert.Equal(t, 2, len(hook.Entries))
	assert.Equal(t, "Fail to persist task: error to persist", hook.LastEntry().Message)
	repository.AssertNumberOfCalls(t, "Save", 1)
	repository.AssertExpectations(t)
}

func TestUnmarshalTaskError(t *testing.T) {
	logger, hook := test.NewNullLogger()
	repository := mockRepository{}
	channel := make(chan *sqs.Message, 1)

	body := "wrong body"
	message := sqs.Message{Body: &body}
	channel <- &message

	service := NewTaskService(logger, &repository)
	service.process(channel)

	assert.Equal(t, 2, len(hook.Entries))
	assert.Equal(t, "Could not unmarshal task: invalid character 'w' looking for beginning of value", hook.LastEntry().Message)
}
