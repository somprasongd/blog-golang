package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Task struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title       string    `gorm:"not null"`
	Description string
	Status      TaskStatus `sql:"task_status" gorm:"default:'open'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tasks []*Task
