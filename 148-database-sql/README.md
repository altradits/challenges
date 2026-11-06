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

---

## Time limit

**25 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 148
```

or run the tests directly:

```bash
go test ./148-database-sql/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 148` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
