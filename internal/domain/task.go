package domain

import "time"

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `gorm:"not null; varchar(120)"`
	IsCompleted bool   `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
