package queue

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type connectionStub struct {
	sqsiface.SQSAPI
}

func (s connectionStub) ReceiveMessage(*sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	return nil, fmt.Errorf("Foo")
}

func TestFoo(t *testing.T) {

	connection := &connectionStub{}
	channel := make(chan *sqs.Message)
	Poll(channel, connection)
}
