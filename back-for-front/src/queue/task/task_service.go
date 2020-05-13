package task

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/fsantiag/track-progress/back-for-front/src/model/task"
	"github.com/fsantiag/track-progress/back-for-front/src/queue"
	"github.com/sirupsen/logrus"
)

var (
	connection sqsiface.SQSAPI
	logger     *logrus.Logger
)

type taskQueueImpl struct{}

//NewTaskQueue is a constructor to create `instances` from TaskQueue with your dependencies
func NewTaskQueue(client sqsiface.SQSAPI, log *logrus.Logger) queue.TaskQueue {
	connection = client
	logger = log
	return taskQueueImpl{}
}

//SendTask is a method that send a message to save
func (tq taskQueueImpl) SendTask(task task.Task, queueURL string) {
	logger.Info("sending message to SQS: ", task)
	body, _ := json.Marshal(&task)

	_, err := connection.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(body)),
		QueueUrl:    aws.String(queueURL),
	})
	if err != nil {
		logger.Error("error while send message to SQS: ", err)
		return
	}
}
