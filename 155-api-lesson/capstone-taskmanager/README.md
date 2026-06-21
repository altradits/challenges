# Capstone: Task Manager API

## What You'll Build

A complete REST API for task management with authentication, database storage, and full documentation.

## Features

- User registration and authentication (JWT)
- CRUD operations for tasks
- Task filtering and pagination
- Input validation
- Comprehensive error handling
- Full test coverage
- OpenAPI/Swagger documentation

## Project Structure

```
capstone-taskmanager/
├── main.go           # Entry point
├── go.mod            # Dependencies
├── handlers/         # HTTP handlers
│   ├── auth.go       # Authentication handlers
│   └── tasks.go      # Task handlers
├── middleware/       # Middleware
│   ├── auth.go       # JWT validation
│   └── logging.go    # Request logging
├── models/           # Data models
│   ├── user.go       # User model
│   └── task.go       # Task model
├── database/         # Database setup
│   └── db.go         # DB connection and migrations
├── docs/             # Swagger documentation
│   └── swagger.json
└── tests/            # Integration tests
    └── api_test.go
```

## API Endpoints

### Authentication

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /api/v1/register | Register new user |
| POST | /api/v1/login | Login and get JWT |

### Tasks

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/v1/tasks | List all tasks (with pagination) |
| GET | /api/v1/tasks/{id} | Get task by ID |
| POST | /api/v1/tasks | Create new task |
| PUT | /api/v1/tasks/{id} | Update task |
| DELETE | /api/v1/tasks/{id} | Delete task |

## Models

### User

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2024-01-01T00:00:00Z"
}
```

### Task

```json
{
  "id": 1,
  "title": "Buy groceries",
  "description": "Milk, eggs, bread",
  "completed": false,
  "user_id": 1,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

## Requirements

1. **Authentication**
   - Password hashing with bcrypt
   - JWT token generation
   - Protected routes

2. **Database**
   - SQLite for storage
   - Proper schema with migrations
   - Connection pooling

3. **Validation**
   - Email format validation
   - Required field validation
   - Return 400 for invalid input

4. **Error Handling**
   - Custom error types
   - Consistent error responses
   - Proper status codes

5. **Testing**
   - Unit tests for handlers
   - Integration tests
   - Mock database for tests

6. **Documentation**
   - OpenAPI 3.0 spec
   - Swagger UI
   - Example requests/responses

## Getting Started

1. Run `go mod tidy` to install dependencies
2. Run `go run .` to start the server
3. Visit `http://localhost:8080/swagger/index.html` for documentation

## Testing

```bash
# Run all tests
go test -v ./...

# Run with coverage
go test -cover ./...
```

## Skills Applied

- All concepts from exercises 01-10
- Real database design
- Security best practices
- Production deployment patterns