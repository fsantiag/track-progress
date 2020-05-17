package sqs

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

var (
	receiveMessage func(*sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)
)

type connectionStub struct {
	sqsiface.SQSAPI
}

func (s connectionStub) ReceiveMessage(receiveMessageInput *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	return receiveMessage(receiveMessageInput)
}

func (s connectionStub) DeleteMessage(*sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	deleteMessageOutput := sqs.DeleteMessageOutput{}
	return &deleteMessageOutput, nil
}

func TestProcessMessage(t *testing.T) {
	logger, _ := test.NewNullLogger()
	connection := &connectionStub{}
	channel := make(chan *sqs.Message, 1)
	receiveMessage = func(receiveMessageInput *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
		body := "Any"
		message := sqs.Message{Body: &body}
		receiveMessageOutput := sqs.ReceiveMessageOutput{}
		receiveMessageOutput.Messages = []*sqs.Message{&message}
		return &receiveMessageOutput, nil
	}

	processMessage(channel, "queue", connection, logger)

	resultMsg := <-channel
	assert.Equal(t, "Any", *resultMsg.Body)
}

func TestProcessMessageWithErrorToReceiveMessage(t *testing.T) {
	logger, hook := test.NewNullLogger()
	connection := &connectionStub{}
	channel := make(chan *sqs.Message)
	receiveMessage = func(receiveMessageInput *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
		return nil, errors.New("wrong message")
	}

	processMessage(channel, "queue", connection, logger)

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, "Fail to receive message: wrong message", hook.LastEntry().Message)
}
