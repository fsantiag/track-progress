package queue

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

//constants
const (
	QueueURL = "http://localhost:4576/123456789012/queue"
)

// Poll messages from SQS and send them into the passed channel
func Poll(channel chan<- *sqs.Message, connection *sqs.SQS) {
	for {
		result, err := connection.ReceiveMessage(&sqs.ReceiveMessageInput{
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            aws.String(QueueURL),
			MaxNumberOfMessages: aws.Int64(1),
			VisibilityTimeout:   aws.Int64(20),
			WaitTimeSeconds:     aws.Int64(10),
		})
		if err != nil {
			fmt.Println("Error", err)
			return
		}

		if len(result.Messages) > 0 {
			channel <- result.Messages[0]

			_, err := connection.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(QueueURL),
				ReceiptHandle: result.Messages[0].ReceiptHandle,
			})

			if err != nil {
				fmt.Println(err.Error())
			}
		}
		time.Sleep(1 * time.Second)
	}
}
