package task

import (
	"encoding/json"
	"net/http"

	"github.com/fsantiag/track-progress/back-for-front/src/configuration"
	model "github.com/fsantiag/track-progress/back-for-front/src/model/task"
	"github.com/fsantiag/track-progress/back-for-front/src/queue"
	"github.com/fsantiag/track-progress/back-for-front/src/queue/task"
	"github.com/fsantiag/track-progress/back-for-front/src/util"
)

var (
	queueURL = util.Getenv("SQS_HOST", "http://localhost:4576") + util.QueueName
)

//SaveTask is a method to receive endpoint data and send to save task
func SaveTask(writer http.ResponseWriter, request *http.Request) {
	session := configuration.NewSession()
	taskQueue := task.NewTaskQueue(session, &configuration.Logger)
	saveTask(writer, request, taskQueue)
}

func saveTask(writer http.ResponseWriter, request *http.Request, taskQueue queue.TaskQueue) {
	var taskToSend model.Task
	_ = json.NewDecoder(request.Body).Decode(&taskToSend)

	taskQueue.SendTask(taskToSend, queueURL)

	writer.WriteHeader(http.StatusCreated)
}
