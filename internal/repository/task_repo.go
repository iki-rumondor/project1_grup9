package repository

import "github.com/iki-rumondor/project1_grup9/internal/domain"

type TaskRepository interface {
	GetAll() ([]domain.Task, error)
	GetByID(id uint) (*domain.Task, error)
	Delete(id uint) (*domain.Task, error)
}
