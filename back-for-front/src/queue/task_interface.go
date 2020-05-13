package queue

import (
	"github.com/fsantiag/track-progress/back-for-front/src/model/task"
)

type (
	//TaskQueue represents a interface to be implemented
	TaskQueue interface {
		SendTask(task task.Task, queueURL string)
	}
)
