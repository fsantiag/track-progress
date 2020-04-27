package queue

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	Endpoint = "http://localhost:4576"
)

func NewSession() *sqs.SQS {
	session := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(endpoints.UsWest2RegionID),
		Endpoint: aws.String(Endpoint),
	}))

	return sqs.New(session)
}

func CreateQueues(svc *sqs.SQS) {
	result, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String("queue"),
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("10"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
	})
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success", *result.QueueUrl)
}
