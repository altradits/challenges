# gin/04-json-crud-api — Full CRUD REST API

## Challenge

Build a complete Task Manager REST API with Gin:

```
POST   /tasks          → create task
GET    /tasks          → list all tasks (optional ?status=pending|done)
GET    /tasks/:id      → get one task
PUT    /tasks/:id      → update task (full replacement)
PATCH  /tasks/:id      → partial update (only provided fields)
DELETE /tasks/:id      → delete task
GET    /tasks/stats    → { "total": N, "pending": N, "done": N }
```

## Task Model

```go
type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title" binding:"required"`
    Description string    `json:"description"`
    Status      string    `json:"status"`   // "pending" or "done"
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## Requirements

- Return proper HTTP status codes (201 for create, 200 for others, 404 for not found, 400 for validation errors)
- `Title` is required — return 400 if missing
- `Status` defaults to `"pending"` on create
- New tasks get auto-incrementing IDs
- Thread-safe (use a mutex)

## Test

```bash
# Create
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Buy groceries","description":"Milk, eggs"}'

# List
curl http://localhost:8080/tasks

# Filter by status
curl "http://localhost:8080/tasks?status=pending"

# Get one
curl http://localhost:8080/tasks/1

# Update status
curl -X PATCH http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"status":"done"}'

# Stats
curl http://localhost:8080/tasks/stats
```

Read `skills.md` before you start.
