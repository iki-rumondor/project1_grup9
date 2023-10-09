package repository

import "github.com/iki-rumondor/project1_grup9/internal/domain"

type TaskRepository interface {
	FindByID(uint) (*domain.Task, error)
	Upsert(*domain.Task) error
}
