package task

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fsantiag/track-progress/back-for-front/src/model/task"
	"github.com/fsantiag/track-progress/back-for-front/src/service/task/internal/mock"
	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

const (
	FirstMessageInLogger  = 0
	SecondMessageInLogger = 1
)

var (
	id, _     = gocql.ParseUUID("57e79ec6-9525-11ea-bb37-0242ac130002")
	taskModel = task.Task{ID: id, Title: "Any Title", Description: "Any description", Status: "Any status"}
)

func TestSendTask_CallSQSMethod(t *testing.T) {
	logger, hook := test.NewNullLogger()
	mockedConnection := mock.SQSMock{}
	mock.SendMessage = func(*sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
		return nil, nil
	}

	taskService := NewTaskQueue(mockedConnection, logger)
	taskService.SendTask(taskModel, "queue")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, "sending message to SQS: {57e79ec6-9525-11ea-bb37-0242ac130002 Any Title Any description Any status}", hook.LastEntry().Message)
}

func TestSendTask_ErrorToSendMessageToSQS(t *testing.T) {
	logger, hook := test.NewNullLogger()
	mockedConnection := mock.SQSMock{}
	mock.SendMessage = func(*sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
		return nil, errors.New("messageError")
	}

	taskService := NewTaskQueue(mockedConnection, logger)
	taskService.SendTask(taskModel, "queue")

	assert.Equal(t, 2, len(hook.Entries))
	assert.Equal(t,
		"sending message to SQS: {57e79ec6-9525-11ea-bb37-0242ac130002 Any Title Any description Any status}",
		hook.AllEntries()[FirstMessageInLogger].Message)
	assert.Equal(t, "error while send message to SQS: messageError", hook.AllEntries()[SecondMessageInLogger].Message)
}
