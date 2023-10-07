package repository

import "gorm.io/gorm"

type TaskRepoImplementation struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TaskRepository {
	return TaskRepoImplementation{
		db: db,
	}
}

