package repository

import (
	"github.com/iki-rumondor/project1_grup9/internal/domain"
	"gorm.io/gorm"
)

type TaskRepoImplementation struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TaskRepository {
	return &TaskRepoImplementation{
		db: db,
	}
}

func (r *TaskRepoImplementation) Upsert(task *domain.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepoImplementation) FindByID(id uint) (*domain.Task, error) {
	var task domain.Task

	if err := r.db.First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &task, nil
}
