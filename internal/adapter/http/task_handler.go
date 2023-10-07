package customHTTP

import "github.com/iki-rumondor/project1_grup9/internal/application"

type TaskHandler struct {
	Service *application.TaskService
}

func NewTaskHandler(service *application.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}