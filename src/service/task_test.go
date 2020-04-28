package service

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fsantiag/track-progress/src/model"
	"github.com/fsantiag/track-progress/src/repository"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type sessionStub struct {
	repository.SessionInterface
}

type connectionRepository struct {
	mock.Mock
	repository.Repository
}

func (c *connectionRepository) Save(session repository.SessionInterface, task model.Task) (err error) {
	args := c.Called(session, task)
	return args.Error(0)
}

func TestProcessTask(t *testing.T) {
	logger, _ := test.NewNullLogger()
	connection := connectionRepository{}
	session := sessionStub{}
	channel := make(chan *sqs.Message, 1)
	task := model.Task{}
	body := `{"title":"any title","description":"any description","status":"any status"}`

	message := sqs.Message{Body: &body}
	channel <- &message
	json.Unmarshal([]byte(body), &task)

	connection.On("Save", session, task).Return(nil)

	process(session, &connection, channel, logger)

	connection.AssertNumberOfCalls(t, "Save", 1)
	connection.AssertExpectations(t)
}

func TestProcessTaskWithErrorToSaveTask(t *testing.T) {
	logger, hook := test.NewNullLogger()
	connection := connectionRepository{}
	session := sessionStub{}
	channel := make(chan *sqs.Message, 1)
	newError := errors.New("error to persist")

	body := `{"title":"any title","description":"any description","status":"any status"}`
	message := sqs.Message{Body: &body}
	channel <- &message

	task := model.Task{}
	json.Unmarshal([]byte(body), &task)
	connection.On("Save", session, task).Return(newError)

	process(session, &connection, channel, logger)

	assert.Equal(t, 2, len(hook.Entries))
	assert.Equal(t, "Fail to persist task: error to persist", hook.LastEntry().Message)
	connection.AssertNumberOfCalls(t, "Save", 1)
	connection.AssertExpectations(t)
}

func TestUnmarshalTaskError(t *testing.T) {
	logger, hook := test.NewNullLogger()
	connection := connectionRepository{}
	session := sessionStub{}
	channel := make(chan *sqs.Message, 1)

	body := "wrong body"
	message := sqs.Message{Body: &body}
	channel <- &message

	process(session, &connection, channel, logger)

	assert.Equal(t, 2, len(hook.Entries))
	assert.Equal(t, "Could not unmarshal task: invalid character 'w' looking for beginning of value", hook.LastEntry().Message)
}
