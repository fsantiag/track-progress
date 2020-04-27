package queue

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

//constants
const (
	QueueURL = "http://localhost:4576/123456789012/queue"
)

// Poll messages from SQS and send them into the passed channel
func Poll(channel chan<- *sqs.Message, connection sqsiface.SQSAPI) {
	for {
		result, err := connection.ReceiveMessage(&sqs.ReceiveMessageInput{
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:          aws.String(QueueURL),
			VisibilityTimeout: aws.Int64(20),
			WaitTimeSeconds:   aws.Int64(10),
		})
		if err != nil {
			fmt.Println("Error", err)
			return
		}

		if len(result.Messages) > 0 {
			fmt.Println("Message received")
			channel <- result.Messages[0]

			_, err := connection.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(QueueURL),
				ReceiptHandle: result.Messages[0].ReceiptHandle,
			})

			if err != nil {
				fmt.Println(err.Error())
			}
		}
		time.Sleep(time.Second)
	}
}
