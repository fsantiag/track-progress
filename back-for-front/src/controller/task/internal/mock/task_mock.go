package mock

import (
	"github.com/fsantiag/track-progress/back-for-front/src/model/task"
	"github.com/fsantiag/track-progress/back-for-front/src/service"
	"github.com/stretchr/testify/mock"
)

//TaskServiceMock is a mock to override method from TaskService
type TaskServiceMock struct {
	mock.Mock
	service.TaskService
}

//SendTask is a override method from TaskService to use as mock
func (ts *TaskServiceMock) SendTask(task task.Task) {
	ts.Called(task)
}
