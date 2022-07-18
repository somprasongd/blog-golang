package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Tasks []*Task

type Task struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title       string    `gorm:"not null"`
	Description string
	Status      TaskStatus `sql:"task_status" gorm:"default:'open'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Task) Open() {
	t.Status = OPEN
}

func (t *Task) InProgress() {
	t.Status = IN_PROGRESS
}

func (t *Task) Done() {
	t.Status = DONE
}
