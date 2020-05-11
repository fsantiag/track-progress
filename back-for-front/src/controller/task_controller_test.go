package controller

import (
	"encoding/json"
	"github.com/fsantiag/track-progress/back-for-front/src/controller/internal/mock"
	"github.com/fsantiag/track-progress/back-for-front/src/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	JSON = `{"id": "e55be5e4-9167-11ea-bb37-0242ac130002", "title":"Any Title", "description":"Any description", "status":"Any status"}`
	body = strings.NewReader(JSON)
)

func TestSaveTask_ReturnStatus201(t *testing.T) {
	err, recorder := executeMethodPostToTask()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestSaveTask_ShouldCallServiceToSendMessageToSQS(t *testing.T) {
	task := model.Task{}
	json.Unmarshal([]byte(JSON), &task)

	request, _ := http.NewRequest(http.MethodPost, "/task", body)
	recorder := httptest.NewRecorder()
	taskService := mock.TaskServiceMock{}

	taskService.On("SendTask", task)

	saveTask(recorder, request, &taskService)

	taskService.AssertNumberOfCalls(t, "SendTask", 1)
}

func executeMethodPostToTask() (error, *httptest.ResponseRecorder) {
	request, err := http.NewRequest(http.MethodPost, "/task", body)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SaveTask)
	handler.ServeHTTP(recorder, request)
	return err, recorder
}
