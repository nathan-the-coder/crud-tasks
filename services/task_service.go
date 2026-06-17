package services

import (
	"context"
	"fmt"
	"strings"
)

func (ts *TaskService) CreateTask(ctx context.Context, title string, description string) error {
	if title == "" {
		return fmt.Errorf("Title of the task must not be empty.")	
	}

	if len(title) > 300 {
		return fmt.Errorf("Title must be below 300 characters.")	
	}

	new_title := strings.ToTitle(title)

	description = strings.TrimSpace(description)
	var new_description *string
	if description == "" {
		new_description = nil
	} else {
		new_description = &description
	}

	exists, _ := ts.store.Exists(ctx, new_title)

	if exists {
		return fmt.Errorf("Task named '%s' already exists.", new_title)
	}

	ts.store.Create(ctx, new_title, *new_description)

	return nil
}
