package application

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/iki-rumondor/project1_grup9/internal/adapter/http/request"
	"github.com/iki-rumondor/project1_grup9/internal/domain"
	"github.com/iki-rumondor/project1_grup9/internal/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	taskID uint        = 1
	task   domain.Task = domain.Task{
		Description: "task description 1",
		IsCompleted: false,
	}
	reqCreateTask request.CreateTask = request.CreateTask{
		Description: "task description 1",
	}
	reqUpdateTask request.UpdateTask = request.UpdateTask{
		ID: 1,
		Description: "task description 1",
	}
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

	mockTask.On("FindByID", taskID).Return(&task, nil)

	tasks, err := services.GetByTaskID(taskID)

	assert.NotNil(t, tasks, "tasks must be there")
	assert.Nil(t, err, "error must be nil")
}

func TestGetTaskByIDFailed(t *testing.T) {
	var mockTask = mocks.NewTaskRepository(t)
	var services = NewTaskService(mockTask)

	mockTask.On("FindByID", taskID).Return(nil, errors.New("failed to get all tasks"))

	tasks, err := services.GetByTaskID(taskID)

	assert.Nil(t, tasks, "tasks must be nil")
	assert.NotNil(t, err, "error must be there")
}

func TestDeleteTaskSuccess(t *testing.T) {
	var mockTask = mocks.NewTaskRepository(t)
	var services = NewTaskService(mockTask)

	mockTask.On("FindByID", taskID).Return(&task, nil)
	mockTask.On("Delete", &task).Return(nil)

	err := services.DeleteTask(taskID)

	assert.Nil(t, err, "error must be nil")
}

func TestDeleteTaskFailedGetData(t *testing.T) {
	var mockTask = mocks.NewTaskRepository(t)
	var services = NewTaskService(mockTask)

	mockTask.On("FindByID", taskID).Return(nil, fmt.Errorf("task with id %d is not found", taskID))
	err := services.DeleteTask(taskID)

	assert.NotNil(t, err, "error must be there")
}

func TestDeleteTaskFailed(t *testing.T) {
	var mockTask = mocks.NewTaskRepository(t)
	var services = NewTaskService(mockTask)

	mockTask.On("FindByID", taskID).Return(&task, nil)
	mockTask.On("Delete", &task).Return(errors.New("failed to delete task"))

	err := services.DeleteTask(taskID)

	assert.NotNil(t, err, "error must be there")
}

func TestCreateTaskSuccess(t *testing.T) {
	var mockTask = mocks.NewTaskRepository(t)
	var services = NewTaskService(mockTask)

	mockTask.On("Upsert", &task).Return(nil)

	err := services.CreateTask(&reqCreateTask)

	assert.Nil(t, err, "error must be nil")
}

func TestCreateTaskFailed(t *testing.T) {
	var mockTask = mocks.NewTaskRepository(t)
	var services = NewTaskService(mockTask)

	mockTask.On("Upsert", &task).Return(errors.New("failed to create a task"))

	err := services.CreateTask(&reqCreateTask)

	assert.NotNil(t, err, "error must be there")
}

func TestUpdateTaskSuccess(t *testing.T) {
	var mockTask = mocks.NewTaskRepository(t)
	var services = NewTaskService(mockTask)

	mockTask.On("FindByID", taskID).Return(&task, nil)
	mockTask.On("Upsert", &task).Return(nil)

	err := services.UpdateTask(&reqUpdateTask)

	assert.Nil(t, err, "error must be nil")
}

func TestUpdateTaskFailedGetData(t *testing.T) {
	var mockTask = mocks.NewTaskRepository(t)
	var services = NewTaskService(mockTask)

	mockTask.On("FindByID", taskID).Return(nil, fmt.Errorf("task with id %d is not found", taskID))
	err := services.UpdateTask(&reqUpdateTask)

	assert.NotNil(t, err, "error must be there")
}

func TestUpdateTaskFailed(t *testing.T) {
	var mockTask = mocks.NewTaskRepository(t)
	var services = NewTaskService(mockTask)

	mockTask.On("FindByID", taskID).Return(&task, nil)
	mockTask.On("Upsert", &task).Return(errors.New("failed to update task"))

	err := services.UpdateTask(&reqUpdateTask)

	assert.NotNil(t, err, "error must be there")
}
