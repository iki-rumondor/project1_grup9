package repository

import "github.com/iki-rumondor/project1_grup9/internal/domain"

type TaskRepository interface {
	Upsert(*domain.Task) error
}
