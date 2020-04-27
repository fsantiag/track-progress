package queue

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fsantiag/track-progress/src/util"
)

// NewSession create a new connection to SQS
func NewSession() *sqs.SQS {
	session := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(endpoints.UsWest2RegionID),
		Endpoint: aws.String(util.Getenv("SQS_HOST", "http://localhost:4576")),
	}))

	return sqs.New(session)
}

// CreateQueues create a new queue in SQS
func CreateQueues(svc *sqs.SQS) string {
	result, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(util.Getenv("SQS_QUEUE_NAME", "queue")),
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("10"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
	})
	if err != nil {
		fmt.Println("Error", err)
		return ""
	}

	return *result.QueueUrl
}
