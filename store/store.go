package store

import "github.com/nathan-the-coder/crud-tasks/db"

type Task struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	CompletedAt *string `json:"completed_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}

type TaskStore struct {
	queries *db.Queries
}
