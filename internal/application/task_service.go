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

func (s *TaskService) GetAll() ([]domain.Task, error) {
	return s.Repo.GetAll()
}

func (s *TaskService) GetByID(id uint) (*domain.Task, error) {
	return s.Repo.GetByID(id)
}

func (s *TaskService) Delete(id uint) (*domain.Task, error) {
	task, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if _, err := s.Repo.Delete(id); err != nil {
		return nil, err
	}

	return task, nil
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

func (s *TaskService) UpdateTask(body *request.UpdateTask) error{

	task, err := s.Repo.FindByID(body.ID)
	if err != nil{
		return err
	}

	task.Description = body.Description

	if err := s.Repo.Upsert(task); err != nil{
		return err
	}
	
	return nil
}
