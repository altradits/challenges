package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

// CreateDatabaseAPI creates a REST API with SQLite database.
// Implement the following:
// - Initialize SQLite database (app.db)
// - Create users table if not exists
// - CRUD endpoints for users
// - Use database for persistent storage
// Hint: Use sql.Open to connect to SQLite
// Hint: Use db.Exec for INSERT/UPDATE/DELETE
// Hint: Use db.Query for SELECT all
// Hint: Use db.QueryRow for SELECT by ID
func CreateDatabaseAPI() {
	// TODO: Implement this function
	// 1. Open SQLite database
	// 2. Create users table
	// 3. Implement handlers for CRUD operations
	// 4. Start the server on port 8080
}

func main() {
	fmt.Println("Starting database API on :8080...")
	CreateDatabaseAPI()
}
