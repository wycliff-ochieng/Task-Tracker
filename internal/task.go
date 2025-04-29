package internal

import (
	"time"
)

type status string

const (
	STATUS_TODO        status = "todo"
	STATUS_IN_PROGRESS status = "in-progree"
	STATUS_DONE        status = "done"
)

type Task struct {
	ID          int32     `json:"id"`
	Description string    `json:"description"`
	Status      status    `json:"status"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
}

func NewTask(id int32, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
