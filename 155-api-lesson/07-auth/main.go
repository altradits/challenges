package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3"
)

// User represents a user in the system
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // Never serialize password
	Role     string `json:"role"`
}

// Claims represents JWT claims
type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// CreateAuthAPI creates an API with JWT authentication.
// Implement the following:
// - POST /api/v1/register -> Create user with hashed password
// - POST /api/v1/login -> Return JWT token
// - GET /api/v1/profile -> Protected route (requires valid token)
// - GET /api/v1/admin -> Admin-only route
// Hint: Use bcrypt for password hashing
// Hint: Use jwt.NewWithClaims for token creation
// Hint: Create middleware to validate tokens
// Hint: Use context to pass user info to handlers
func CreateAuthAPI() {
	// TODO: Implement this function
	// 1. Open database and create users table
	// 2. Implement register handler (hash password)
	// 3. Implement login handler (issue JWT)
	// 4. Implement auth middleware
	// 5. Implement role-based access control
	// 6. Start the server on port 8080
}

func main() {
	fmt.Println("Starting auth API on :8080...")
	CreateAuthAPI()
}
