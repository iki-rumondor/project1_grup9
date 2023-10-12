package application

import (
	"errors"
	"testing"
	"time"

	"github.com/iki-rumondor/project1_grup9/internal/domain"
	"github.com/iki-rumondor/project1_grup9/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTasksSuccess(t *testing.T) {
	var mockTask = mocks.NewTaskRepository(t)
	var services = NewTaskService(mockTask)

	mockTask.On("FindAll").Return([]domain.Task{
		{
			Description: "task description 1",
			IsCompleted: false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Description: "task description 2",
			IsCompleted: false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil)

	tasks, err := services.GetAllTasks()

	assert.NotNil(t, tasks, "tasks must be not nil")
	assert.Nil(t, err, "error must be nil")
}

func TestGetAllTasksFailed(t *testing.T) {
	var mockTask = mocks.NewTaskRepository(t)
	var services = NewTaskService(mockTask)

	mockTask.On("FindAll").Return(nil, errors.New("failed to get all tasks"))

	tasks, err := services.GetAllTasks()

	assert.Nil(t, tasks, "tasks must be nil")
	assert.NotNil(t, err, "error must be there")
}

func TestGetTaskByIDSuccess(t *testing.T) {
	var mockTask = mocks.NewTaskRepository(t)
	var services = NewTaskService(mockTask)
	var taskID uint = 1

	mockTask.On("FindByID", taskID).Return(&domain.Task{
		Description: "task description 1",
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil)

	tasks, err := services.GetByTaskID(taskID)

	assert.NotNil(t, tasks, "tasks must be there")
	assert.Nil(t, err, "error must be nil")
}

func TestGetTaskByIDFailed(t *testing.T) {
	var mockTask = mocks.NewTaskRepository(t)
	var services = NewTaskService(mockTask)
	var taskID uint = 1

	mockTask.On("FindByID", taskID).Return(nil, errors.New("failed to get all tasks"))

	tasks, err := services.GetByTaskID(taskID)

	assert.Nil(t, tasks, "tasks must be nil")
	assert.NotNil(t, err, "error must be there")
}
