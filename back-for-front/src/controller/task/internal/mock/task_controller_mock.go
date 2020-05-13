package mock

import (
	"github.com/fsantiag/track-progress/back-for-front/src/model/task"
	"github.com/fsantiag/track-progress/back-for-front/src/queue"
	"github.com/stretchr/testify/mock"
)

//TaskQueueMock is a mock to override method from TaskQueue
type TaskQueueMock struct {
	mock.Mock
	queue.TaskQueue
}

//SendTask is a override method from TaskService to use as mock
func (tq *TaskQueueMock) SendTask(task task.Task, queueURL string) {
	tq.Called(task, queueURL)
}
