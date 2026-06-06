package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3"
)

// User represents a user in the system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

// Task represents a task in the system
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// APIError represents an error response
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Claims represents JWT claims
type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// App represents the application
type App struct {
	DB     *sql.DB
	Secret string
}

// CreateTaskManagerAPI creates a complete task manager API.
// Implement all features:
// - User registration and login
// - Task CRUD operations
// - JWT authentication
// - Database storage
// - Error handling
// - Logging
// Hint: Use the App struct to hold dependencies
// Hint: Create separate handlers for auth and tasks
// Hint: Use middleware for authentication
// Hint: Add proper validation
func CreateTaskManagerAPI() {
	// TODO: Implement this function
	// 1. Initialize database
	// 2. Create tables
	// 3. Set up routes
	// 4. Add middleware
	// 5. Start server
}

func main() {
	fmt.Println("Starting Task Manager API on :8080...")
	CreateTaskManagerAPI()
}
