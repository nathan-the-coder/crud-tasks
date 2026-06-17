package store

import (
	"time"

	"github.com/nathan-the-coder/crud-tasks/db"
)

type Task struct {
	Id          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type TaskStore struct {
	queries *db.Queries
}
