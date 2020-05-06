package service

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fsantiag/track-progress/backend/src/model"
	"github.com/fsantiag/track-progress/backend/src/repository"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
	repository.Repository
}

func (c *mockRepository) Save(task model.Task) (err error) {
	args := c.Called(task)
	return args.Error(0)
}

func TestProcessTask(t *testing.T) {
	logger, _ := test.NewNullLogger()
	mockedRepository := mockRepository{}
	body := `{"title":"any title","description":"any description","status":"any status"}`
	message := sqs.Message{Body: &body}
	task := model.Task{}
	_ = json.Unmarshal([]byte(body), &task)

	mockedRepository.On("Save", task).Return(nil)

	service := NewTaskService(logger, &mockedRepository)
	service.process(&message)

	mockedRepository.AssertNumberOfCalls(t, "Save", 1)
	mockedRepository.AssertExpectations(t)
}

func TestProcessTaskWithErrorToSaveTask(t *testing.T) {
	logger, hook := test.NewNullLogger()
	mockedRepository := mockRepository{}
	expectedError := errors.New("error to persist")

	body := `{"title":"any title","description":"any description","status":"any status"}`
	message := sqs.Message{Body: &body}

	task := model.Task{}
	_ = json.Unmarshal([]byte(body), &task)
	mockedRepository.On("Save", task).Return(expectedError)

	service := NewTaskService(logger, &mockedRepository)
	service.process(&message)

	assert.Equal(t, 2, len(hook.Entries))
	assert.Equal(t, "Fail to persist task: error to persist", hook.LastEntry().Message)
	mockedRepository.AssertNumberOfCalls(t, "Save", 1)
	mockedRepository.AssertExpectations(t)
}

func TestUnmarshalTaskError(t *testing.T) {
	logger, hook := test.NewNullLogger()
	mockedRepository := mockRepository{}

	body := "wrong body"
	message := sqs.Message{Body: &body}

	service := NewTaskService(logger, &mockedRepository)
	service.process(&message)

	assert.Equal(t, 2, len(hook.Entries))
	assert.Equal(t, "Could not unmarshal task: invalid character 'w' looking for beginning of value", hook.LastEntry().Message)
}
