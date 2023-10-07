package application

import "github.com/iki-rumondor/project1_grup9/internal/repository"

type TaskService struct {
	Repo repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService{
	return &TaskService{
		Repo: repo,
	}
}