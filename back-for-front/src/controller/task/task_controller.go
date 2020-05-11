package task

import (
	"encoding/json"
	"net/http"

	model "github.com/fsantiag/track-progress/back-for-front/src/model/task"
	"github.com/fsantiag/track-progress/back-for-front/src/service"
	"github.com/fsantiag/track-progress/back-for-front/src/service/task"
)

//SaveTask is a method to receive endpoint data and send to save task
func SaveTask(writer http.ResponseWriter, request *http.Request) {
	taskService := task.NewTaskService()
	saveTask(writer, request, taskService)
}

func saveTask(writer http.ResponseWriter, request *http.Request, service service.TaskService) {
	var taskToSend model.Task
	_ = json.NewDecoder(request.Body).Decode(&taskToSend)

	service.SendTask(taskToSend)

	writer.WriteHeader(http.StatusCreated)
}