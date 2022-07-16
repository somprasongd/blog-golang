package model

import "database/sql/driver"

type TaskStatus string

const (
	OPEN        TaskStatus = "open"
	IN_PROGRESS TaskStatus = "in_progress"
	DONE        TaskStatus = "done"
)

func (e *TaskStatus) Scan(value interface{}) error {
	*e = TaskStatus(value.(string))
	return nil
}

func (e TaskStatus) Value() (driver.Value, error) {
	return string(e), nil
}

func (e TaskStatus) String() string {
	switch e {
	case IN_PROGRESS:
		return "in_progress"
	case DONE:
		return "done"
	default:
		return "open"
	}
}
