package repository

import (
	"github.com/iki-rumondor/project1_grup9/internal/domain"
)

type TaskRepository interface {
	FindAll() ([]domain.Task, error)
	FindByID(uint) (*domain.Task, error)
	Upsert(*domain.Task) error
	Delete(*domain.Task) error
}
