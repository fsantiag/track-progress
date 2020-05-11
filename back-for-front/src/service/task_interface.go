package service

import "github.com/fsantiag/track-progress/back-for-front/src/model"

//TaskService represents a interface to be implemented
type TaskService interface {
	SendTask(task model.Task)
}