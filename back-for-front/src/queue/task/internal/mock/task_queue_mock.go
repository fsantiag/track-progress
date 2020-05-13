package mock

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

var (
	// SendMessage is a method to be implemented in each test
	SendMessage func(*sqs.SendMessageInput) (*sqs.SendMessageOutput, error)
)

// SQSMock is a struct to mock as interface by sqsiface.SQSAPI
type SQSMock struct {
	sqsiface.SQSAPI
}

// SendMessage is a override method by sqsiface.SQSAPI to send message to SQS
func (m SQSMock) SendMessage(message *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	return SendMessage(message)
}
