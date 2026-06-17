package store

import (
	"context"
	"database/sql"

	"github.com/nathan-the-coder/crud-tasks/db"
	"github.com/nathan-the-coder/crud-tasks/models"
)

func NewTaskStore(q *db.Queries) *TaskStore {
	return &TaskStore{queries: q}
}

func toDTO(dbTask db.Task) models.Task {
	task := models.Task{
		Id:          dbTask.ID,
		Title:       dbTask.Title,
		Description: dbTask.Description.String,
		Status:      string(dbTask.Status),
	}

	if dbTask.CompletedAt.Valid {
		task.CompletedAt = &dbTask.CompletedAt.Time
	}
	if dbTask.UpdatedAt.Valid {
		task.UpdatedAt = &dbTask.UpdatedAt.Time
	}

	return task
}

func (ts *TaskStore) Exists(ctx context.Context, title string) (bool, error) {
	tasks, err := ts.queries.ListTasks(ctx)
	if err != nil {
		return false, err
	}

	is_duplicate := false

	for i := range tasks {
		task := tasks[i]
		if task.Title == title {
			is_duplicate = true
		} else {
			is_duplicate = false
		}
	}

	return is_duplicate, nil
}

func (ts *TaskStore) Create(ctx context.Context, title string, description string) (int64, error) {

	arg := db.CreateTaskParams{
		Title:       title,
		Description: sql.NullString{String: description, Valid: true},
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

func (ts *TaskStore) Get(ctx context.Context, id int64) (*models.Task, error) {

	result, err := ts.queries.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	task := toDTO(result)

	return &task, nil
}

func (ts *TaskStore) List(ctx context.Context) ([]models.Task, error) {
	results, err := ts.queries.ListTasks(ctx)
	if err != nil {
		return nil, err
	}

	tasks := make([]models.Task, 0, len(results))

	for i := range results {
		result := toDTO(results[i])
		tasks = append(tasks, result)
	}

	return tasks, nil
}

func (ts *TaskStore) Update(ctx context.Context, id int64, newTitle string, newDescription string) error {
	params := db.UpdateTaskParams{
		ID:          id,
		Title:       newTitle,
		Description: sql.NullString{String: newDescription},
	}

	err := ts.queries.UpdateTask(ctx, params)

	if err != nil {
		return err
	}

	return nil
}

func (ts *TaskStore) Mark(ctx context.Context, id int64, status string) error {

	err := ts.queries.UpdateTaskStatus(ctx, db.UpdateTaskStatusParams{
		ID:     id,
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
