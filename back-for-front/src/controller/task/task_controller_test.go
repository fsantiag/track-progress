package task

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fsantiag/track-progress/back-for-front/src/controller/task/internal/mock"
	model "github.com/fsantiag/track-progress/back-for-front/src/model/task"
	"github.com/stretchr/testify/assert"
)

var (
	JSON = `{"id": "e55be5e4-9167-11ea-bb37-0242ac130002", "title":"Any Title", "description":"Any description", "status":"Any status"}`
)

func TestSaveTask_ReturnStatus201(t *testing.T) {
	err, recorder := executeMethodPostToTask()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestSaveTask_ShouldCallServiceToSendMessageToSQS(t *testing.T) {
	task := model.Task{}
	_ = json.Unmarshal([]byte(JSON), &task)
	request, _ := http.NewRequest(http.MethodPost, "/task", strings.NewReader(JSON))
	recorder := httptest.NewRecorder()
	mockedTaskService := mock.TaskServiceMock{}
	mockedTaskService.On("SendTask", task)

	saveTask(recorder, request, &mockedTaskService)

	mockedTaskService.AssertNumberOfCalls(t, "SendTask", 1)
	mockedTaskService.AssertExpectations(t)
}

func executeMethodPostToTask() (error, *httptest.ResponseRecorder) {
	request, err := http.NewRequest(http.MethodPost, "/task", strings.NewReader(JSON))
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SaveTask)
	handler.ServeHTTP(recorder, request)
	return err, recorder
}
