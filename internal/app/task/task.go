package task

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Date_created int64  `json:"date_created"`
	Done         bool   `json:"done_status"`
}

func NewTask(name string) *Task {
	now := time.Now().Unix()
	uuid := uuid.New().String()

	return &Task{
		Id:           uuid,
		Name:         name,
		Done:         false,
		Date_created: now,
	}
}

// Task struct getters
func (task *Task) GetId() string {
	return task.Id
}

func (task *Task) SetStatus(status bool) {
	task.Done = status
}
