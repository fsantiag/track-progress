package task

import (
	"fmt"

	"github.com/fsantiag/track-progress/back-for-front/src/model"
	"github.com/fsantiag/track-progress/back-for-front/src/service"
)

type taskServiceImpl struct {}

//NewTaskService is a constructor to create `instances` from TaskService with your dependencies
func NewTaskService() service.TaskService {
	return taskServiceImpl{}
}

//SendTask is a method that send a message to save
func (ts taskServiceImpl) SendTask(task model.Task) {
	fmt.Println(task)
}