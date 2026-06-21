# Skills for 148-database-sql

## What You'll Learn

**Previous:** [147-http-basics](../147-http-basics/skills.md) | **Next:** [149-environment-and-config](../149-environment-and-config/skills.md)

**Challenge:** Connect to SQLite, run queries, handle rows, use transactions.

## The `database/sql` Package

`database/sql` is the standard library database abstraction. You pair it with a driver.

### Opening a Connection

```go
import (
    "database/sql"
    _ "modernc.org/sqlite"  // side-effect import: registers the "sqlite" driver
)

db, err := sql.Open("sqlite", "./myapp.db")
if err != nil {
    return err
}
defer db.Close()

// Verify connection
if err := db.Ping(); err != nil {
    return err
}
```

`sql.Open` doesn't connect — it just validates the arguments. `Ping` forces a real connection.

### Executing Statements (INSERT, UPDATE, DELETE)

```go
result, err := db.Exec(
    "INSERT INTO users (name, email) VALUES (?, ?)",
    "Alice", "alice@example.com",
)
id, _ := result.LastInsertId()
rows, _ := result.RowsAffected()
```

Use `?` placeholders (SQLite). PostgreSQL uses `$1`, `$2`.

### Querying a Single Row

```go
var u User
err := db.QueryRow(
    "SELECT id, name, email FROM users WHERE id = ?", id,
).Scan(&u.ID, &u.Name, &u.Email)

if err == sql.ErrNoRows {
    return User{}, fmt.Errorf("user %d: not found", id)
}
```

`QueryRow` never returns an error directly — the error surfaces on `Scan`.

### Querying Multiple Rows

```go
rows, err := db.Query("SELECT id, name, email FROM users ORDER BY id")
if err != nil {
    return nil, err
}
defer rows.Close()  // ALWAYS close rows

var users []User
for rows.Next() {
    var u User
    if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
        return nil, err
    }
    users = append(users, u)
}
return users, rows.Err()  // check iteration error
```

**Always call `rows.Close()` and check `rows.Err()`.**

### Prepared Statements

```go
stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
if err != nil {
    return err
}
defer stmt.Close()

for _, u := range newUsers {
    _, err = stmt.Exec(u.Name, u.Email)
}
```

Prepare once, execute many times — more efficient for repeated queries.

### Transactions

```go
tx, err := db.Begin()
if err != nil {
    return err
}
defer tx.Rollback()  // no-op if already committed

_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromID)
if err != nil {
    return err  // Rollback called by defer
}
_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toID)
if err != nil {
    return err
}
return tx.Commit()
```

**Pattern:** `defer tx.Rollback()` + explicit `tx.Commit()` at the end.

### Context-Aware Queries

Pass a context to respect timeouts and cancellation:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

rows, err := db.QueryContext(ctx, "SELECT ...")
_, err = db.ExecContext(ctx, "INSERT ...", args...)
```

Always use `Context` variants in production code.

### Connection Pool Settings

```go
db.SetMaxOpenConns(25)          // max open connections
db.SetMaxIdleConns(25)          // max idle connections in pool
db.SetConnMaxLifetime(5 * time.Minute)  // max connection lifetime
```

### Solving the Challenge

```go
package piscine

import (
    "database/sql"
    _ "modernc.org/sqlite"
)

type User struct {
    ID    int64
    Name  string
    Email string
}

func OpenDB(path string) (*sql.DB, error) {
    db, err := sql.Open("sqlite", path)
    if err != nil {
        return nil, err
    }
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id    INTEGER PRIMARY KEY AUTOINCREMENT,
        name  TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE
    )`)
    return db, err
}

func InsertUser(db *sql.DB, name, email string) (int64, error) {
    res, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", name, email)
    if err != nil {
        return 0, err
    }
    return res.LastInsertId()
}

func GetUser(db *sql.DB, id int64) (User, error) {
    var u User
    err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).
        Scan(&u.ID, &u.Name, &u.Email)
    return u, err
}

func ListUsers(db *sql.DB) ([]User, error) {
    rows, err := db.Query("SELECT id, name, email FROM users ORDER BY id")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var users []User
    for rows.Next() {
        var u User
        if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    return users, rows.Err()
}

func DeleteUser(db *sql.DB, id int64) error {
    _, err := db.Exec("DELETE FROM users WHERE id = ?", id)
    return err
}
```

**Next:** [149-environment-and-config](../149-environment-and-config/README.md)
