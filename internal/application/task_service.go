package application

import (
	"github.com/iki-rumondor/project1_grup9/internal/domain"
	"github.com/iki-rumondor/project1_grup9/internal/repository"
)

type TaskService struct {
	Repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
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
