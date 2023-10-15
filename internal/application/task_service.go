package application

import (
	"github.com/iki-rumondor/project1_grup9/internal/adapter/http/request"
	"github.com/iki-rumondor/project1_grup9/internal/adapter/http/response"
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

func (s *TaskService) GetAllTasks() ([]*response.Task, error) {
	tasks, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []*response.Task

	for _, task := range tasks {
		res = append(res, &response.Task{
			Description: task.Description,
			IsCompleted: task.IsCompleted,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return res, nil
}

func (s *TaskService) GetByTaskID(id uint) (*response.Task, error) {
	task, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	res := response.Task{
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}

	return &res, nil
}

func (s *TaskService) DeleteTask(id uint) error {
	task, err := s.Repo.FindByID(id)
	if err != nil {
		return err
	}

	if err := s.Repo.Delete(task); err != nil {
		return err
	}

	return nil
}

func (s *TaskService) CreateTask(body *request.Task) error {
	task := domain.Task{
		Description: body.Description,
		IsCompleted: false,
	}

	if err := s.Repo.Upsert(&task); err != nil {
		return err
	}

	return nil
}

func (s *TaskService) UpdateTask(body *request.TaskWithID) error {

	task, err := s.Repo.FindByID(body.ID)
	if err != nil {
		return err
	}

	task.Description = body.Description

	if err := s.Repo.Upsert(task); err != nil {
		return err
	}

	return nil
}
