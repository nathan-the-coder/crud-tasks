-- name: GetTask :one
SELECT * FROM tasks
WHERE id = ?;

-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY created_at DESC;

-- name: CreateTask :execresult
INSERT INTO tasks (
  title, description
) VALUES (
  ?, ?
);

-- name: UpdateTaskStatus :exec
UPDATE tasks
SET status = ?
WHERE id = ?;

-- name: UpdateTask :exec
UPDATE tasks 
  SET title = ?, description = ?
  WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?
