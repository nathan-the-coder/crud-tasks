package store

import (
	"fmt"
	"maps"
	"time"

	"github.com/nathan-the-coder/crud-tasks/utils"
)

type Task struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	CompletedAt *string `json:"completed_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}

type TaskStore struct {
	tasks map[string]Task
}

func NewTaskStore() TaskStore {
	return TaskStore{tasks: make(map[string]Task)}
}

func (ts *TaskStore) Create(title string, description string) (string, error) {
	id := utils.IDGen()

	// Check for duplicates
	if ts.tasks[id].Title == title {
		return "", fmt.Errorf("Existing record of task '%s' found.", title)
	}

	ts.tasks[id] = Task{
		Id: id,
		Title:       title,
		Description: description,
		Status:      "todo",
	}

	return id, nil
}

func (ts *TaskStore) Get(id string) (*Task, error) {
	task, ok := ts.tasks[id]

	if !ok {
		return nil, fmt.Errorf("No record of id ('%s') found", id)
	}

	return &task, nil
}

func (ts *TaskStore) List() map[string]Task {
	return ts.tasks
}

func (ts *TaskStore) Complete(id string) {
	time := time.Now().Format("2006-01-02 15:04:05")

	task := ts.tasks[id]
	task.Status = "done"
	task.CompletedAt = &time

	ts.tasks[id] = task
}

func (ts *TaskStore) Delete(id string) {
	maps.DeleteFunc(ts.tasks, func(k string, v Task) bool {
		if k == id {
			return true 
		}
		return false
	})
}
