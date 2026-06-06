# 06. Database Integration

## What You'll Learn

This exercise teaches you **database integration in Go**. By the end, you'll understand:
- Using `database/sql` with SQLite
- SQL queries (SELECT, INSERT, UPDATE, DELETE)
- Connection pooling
- Struct scanning
- Database migrations

## Theory: Database in Go

### The database/sql Package

Go's standard library provides `database/sql` for database operations:

```go
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)
```

### Opening a Database

```go
db, err := sql.Open("sqlite3", "./app.db")
defer db.Close()
```

### Queries

```go
// Query (read)
rows, err := db.Query("SELECT * FROM users")
// QueryRow (single result)
row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
// Exec (write)
result, err := db.Exec("INSERT INTO users (name) VALUES (?)", name)
```

### Scanning Results

```go
var user User
err := row.Scan(&user.ID, &user.Name)
```

## Step-by-Step Approach

1. **Open database connection**
2. **Create tables** if not exist
3. **Implement CRUD operations**
4. **Handle errors** properly
5. **Close connections** on shutdown

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Not closing rows | Memory leak | Use `defer rows.Close()` |
| Not handling errors | Silent failures | Always check errors |
| Not using prepared statements | SQL injection | Use `?` placeholders |
| Not using connection pool | Performance issues | `sql.Open` handles pooling |

## Practice Tips

- Use `sqlx` or `gorm` for production (but learn `database/sql` first)
- Always use prepared statements
- Handle `sql.ErrNoRows` for not found
- Use transactions for multiple operations

## The Challenge

Create a REST API with SQLite database storage.

### Expected Function

```go
func CreateDatabaseAPI() {
    // Your code here
    // Implement:
    // - Initialize SQLite database
    // - Create users table
    // - CRUD endpoints for users
    // - Use database for storage
}
```

### Test Cases

| Request | Expected Response |
|---------|-----------------|
| POST /api/v1/users | 201 + created user in DB |
| GET /api/v1/users | 200 + all users from DB |
| GET /api/v1/users/1 | 200 + user from DB |
| PUT /api/v1/users/1 | 200 + updated user in DB |
| DELETE /api/v1/users/1 | 204 + user removed from DB |

### Usage Example

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Starting database API on :8080...")
    CreateDatabaseAPI()
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. How do you open a SQLite database in Go?
2. How do you scan a row into a struct?
3. What's the difference between `Query` and `Exec`?

## Next Steps

After completing this, you'll be ready for:
- **07-auth**: Adding authentication
- **08-errors**: Better error handling