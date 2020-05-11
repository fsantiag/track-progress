package task

import (
	"fmt"

	"github.com/fsantiag/track-progress/back-for-front/src/model/task"
	"github.com/fsantiag/track-progress/back-for-front/src/service"
)

type taskServiceImpl struct {}

//NewTaskService is a constructor to create `instances` from TaskService with your dependencies
func NewTaskService() service.TaskService {
	return taskServiceImpl{}
}

//SendTask is a method that send a message to save
func (ts taskServiceImpl) SendTask(task task.Task) {
	fmt.Println(task)
}