package application

import (
	"github.com/iki-rumondor/project1_grup9/internal/adapter/http/request"
	"github.com/iki-rumondor/project1_grup9/internal/domain"
	"github.com/iki-rumondor/project1_grup9/internal/repository"
)

type TaskService struct {
	Repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService{
	return &TaskService{
		Repo: repo,
	}
}

func (s *TaskService) CreateTask(body *request.CreateTask) error{
	task := domain.Task{
		Description: body.Description,
		IsCompleted: false,
	}

	if err := s.Repo.Upsert(&task); err != nil{
		return err
	}
	
	return nil
}