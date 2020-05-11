package service

import (
	"github.com/fsantiag/track-progress/back-for-front/src/model/task"
)

type (
	//TaskService represents a interface to be implemented
	TaskService interface {
		SendTask(task task.Task)
	}
)