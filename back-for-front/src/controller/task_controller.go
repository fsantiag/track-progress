package controller

import (
	"encoding/json"
	"github.com/fsantiag/track-progress/back-for-front/src/model"
	"github.com/fsantiag/track-progress/back-for-front/src/service"
	"github.com/fsantiag/track-progress/back-for-front/src/service/task"
	"net/http"
)

//SaveTask is a method to receice endpoint data and send to save task
func SaveTask(writer http.ResponseWriter, request *http.Request) {
	taskService := task.NewTaskService()
	saveTask(writer, request, taskService)
}

func saveTask(writer http.ResponseWriter, request *http.Request, service service.TaskService) {
	var task model.Task
	json.NewDecoder(request.Body).Decode(&task)

	service.SendTask(task)

	writer.WriteHeader(http.StatusCreated)
}
