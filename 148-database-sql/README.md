# 148-database-sql — database/sql and SQLite

## Challenge

Using `database/sql` with `modernc.org/sqlite` (a pure-Go SQLite driver):

```go
// OpenDB opens a SQLite database at path and creates the users table if it doesn't exist.
func OpenDB(path string) (*sql.DB, error)

// InsertUser inserts a user and returns the new row's ID.
func InsertUser(db *sql.DB, name, email string) (int64, error)

// GetUser returns the user with the given id.
// Returns sql.ErrNoRows if not found.
func GetUser(db *sql.DB, id int64) (User, error)

// ListUsers returns all users ordered by id.
func ListUsers(db *sql.DB) ([]User, error)

// DeleteUser deletes the user with the given id.
func DeleteUser(db *sql.DB, id int64) error
```

Where `User` is `struct { ID int64; Name, Email string }`.

Read `skills.md` before you start. Add `modernc.org/sqlite` with `go get modernc.org/sqlite`.
