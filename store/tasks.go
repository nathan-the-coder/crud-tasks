package store

import (
	"context"
	"database/sql"

	"github.com/nathan-the-coder/crud-tasks/db"
)

func NewTaskStore(q *db.Queries) *TaskStore {
	return &TaskStore{queries: q}
}

func (ts *TaskStore) Create(ctx context.Context, title string, description string) (int64, error) {

	arg := db.CreateTaskParams{
		Title: title,
		Description: sql.NullString{String: description, Valid: false},
	}

	result, err := ts.queries.CreateTask(ctx, arg)
	if err != nil {
		return 0, err
	}

	insertedTaskID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertedTaskID, nil
}

func (ts *TaskStore) Get(ctx context.Context, id int64) (*db.Task, error) {

	result, err := ts.queries.GetTask(ctx, id)
	if err != nil {
		return nil, err;	
	}

	return &result, nil
}

func (ts *TaskStore) List(ctx context.Context) ([]db.Task, error) {
	result, err := ts.queries.ListTasks(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ts *TaskStore) Mark(ctx context.Context, id int64, status string) error {

	err := ts.queries.UpdateTaskStatus(ctx, db.UpdateTaskStatusParams{
		ID: id,
		Status: db.TasksStatus(status),
	})

	if err != nil {
		return err
	}

	return nil
}

func (ts *TaskStore) Delete(ctx context.Context, id int64) error {
	err := ts.queries.DeleteTask(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
