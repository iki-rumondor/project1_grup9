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

func (r TaskRepoImplementation) GetAll() ([]domain.Task, error) {
    var tasks []domain.Task
    result := r.db.Find(&tasks)
    if result.Error != nil {
        return nil, result.Error
    }
    return tasks, nil
}

func (r TaskRepoImplementation) GetByID(id uint) (*domain.Task, error) {
	var task domain.Task
	result := r.db.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (r TaskRepoImplementation) Delete(id uint) (*domain.Task, error) {
	var task domain.Task
	result := r.db.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	
	deleteResult := r.db.Delete(&task, id)
    if deleteResult.Error != nil {
        return nil, deleteResult.Error
    }

	return &task, nil

}