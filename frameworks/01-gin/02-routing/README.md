# gin/02-routing — Routing, Groups, and Query Params

## Challenge

Extend your Gin server with:

```
GET  /users               → list of all users (from in-memory slice)
GET  /users/:id           → single user by id, 404 if not found
POST /users               → create a user (JSON body: {"name": "Alice", "email": "alice@example.com"})
PUT  /users/:id           → update user
DELETE /users/:id         → delete user

GET /search?q=alice&limit=10  → search users by name query param
```

Use an in-memory `[]User` slice (no database yet). Thread-safety is not required here.

## User Type

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

## Test

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com"}'

curl http://localhost:8080/users
curl http://localhost:8080/users/1
curl http://localhost:8080/search?q=alice
```

Read `skills.md` before you start.
