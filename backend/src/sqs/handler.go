package sqs

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/sirupsen/logrus"
)

// Poll messages from SQS and send them into the passed channel
func Poll(channel chan<- *sqs.Message, queueURL string, connection sqsiface.SQSAPI, logger *logrus.Logger) {
	for {
		processMessage(channel, queueURL, connection, logger)
	}
}

func processMessage(channel chan<- *sqs.Message, queueURL string, connection sqsiface.SQSAPI, logger *logrus.Logger) {
	result, err := connection.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:          aws.String(queueURL),
		VisibilityTimeout: aws.Int64(20),
		WaitTimeSeconds:   aws.Int64(10),
	})

	if err != nil {
		logger.Error("Fail to receive message: ", err)
	} else if len(result.Messages) > 0 {
		logger.Info("Message received: ", result.Messages[0].Body)
		channel <- result.Messages[0]

		_, err := connection.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      aws.String(queueURL),
			ReceiptHandle: result.Messages[0].ReceiptHandle,
		})

		if err != nil {
			logger.Warn(err.Error())
		}
	}
	time.Sleep(time.Second)
}
