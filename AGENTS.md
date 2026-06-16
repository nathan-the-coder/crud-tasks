# crud-tasks

Go 1.26 CRUD API for tasks using MySQL and sqlc.

## Quick start

```bash
air                  # hot-reload dev server on :8080
sqlc generate        # regen db/ code after editing db/query.sql or db/schema.sql
go build -o ./tmp/main .   # manual build
```

## Architecture

Four layers, no framework:

```
handlers/  (HTTP layer, Go 1.22+ enhanced ServeMux routing with method prefixes)
  -> store/   (business logic, constructor injection)
    -> db/     (sqlc-generated type-safe queries, checked in)
      -> MySQL
```

Entrypoint: `main.go`. DSN is hardcoded (`root:lein2324@/task_manager?parseTime=true`).

## Routes (all in main.go)

| Method | Path | Handler |
|--------|------|---------|
| GET | /health | HealthHandler |
| GET | /tasks | List |
| POST | /tasks | Create |
| GET | /tasks/{id} | Get |
| PUT | /tasks/{status}/{id} | Update |

No DELETE route registered. `Delete` handler is a no-op stub.

## Known bugs

- `store/tasks.go:18`: Description always sets `Valid: false` — descriptions never persist.
- `handlers/tasks.go:25`: List builds `map[string]any` keyed by strconv-formatted ID (uncommitted).
- No tests, no CI, no linter config exist.
