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

func NewTaskStore() *TaskStore {
	return &TaskStore{tasks: make(map[string]Task)}
}

func (ts *TaskStore) Create(title string, description string) (string, error) {
	id := utils.IDGen()

	// Check for duplicates
	for _, v := range ts.tasks {
		if v.Title == title {
			return "", fmt.Errorf("Existing record of task with a title of '%s' found.", title)
		}
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

func (ts *TaskStore) Mark(id string, status string) (string, error) {

	time := time.Now().Format("2006-01-02 15:04:05")

	task := ts.tasks[id]
	fmt.Println(task.Id, id)
	if task.Id == "" {
		return "", fmt.Errorf("Task (%s) doesn't exists.", id)
	}

	task.Status = status
	switch status {
case "done":
		task.CompletedAt = &time
	case "in-progress", "todo":
		task.UpdatedAt = &time
	default:
		return "", fmt.Errorf("Task Status '%s' doesn't exist. Please try 'todo', 'in-progress' or 'done'.", status)
	}

	ts.tasks[id] = task
	return id, nil
}

func (ts *TaskStore) Delete(id string) {
	maps.DeleteFunc(ts.tasks, func(k string, v Task) bool {
		if k == id {
			return true 
		}
		return false
	})
}
