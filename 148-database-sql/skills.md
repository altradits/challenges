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

### SQL JOINs — Querying Across Tables

Real schemas have related tables. Use JOINs to combine them:

```sql
-- INNER JOIN: only rows that match in both tables
SELECT u.id, u.name, p.title, p.amount
FROM   users u
INNER JOIN payments p ON p.user_id = u.id
WHERE  u.id = ?

-- LEFT JOIN: all users, even those with no payments
SELECT u.name, COUNT(p.id) AS payment_count
FROM   users u
LEFT JOIN payments p ON p.user_id = u.id
GROUP BY u.id
ORDER BY payment_count DESC
```

In Go, scan multiple tables' columns into a combined struct:

```go
type UserPayment struct {
    UserName      string
    PaymentTitle  string
    Amount        float64
}

rows, err := db.Query(`
    SELECT u.name, p.title, p.amount
    FROM users u
    INNER JOIN payments p ON p.user_id = u.id
    WHERE u.id = ?`, userID)
if err != nil {
    return nil, err
}
defer rows.Close()

var results []UserPayment
for rows.Next() {
    var up UserPayment
    rows.Scan(&up.UserName, &up.PaymentTitle, &up.Amount)
    results = append(results, up)
}
return results, rows.Err()
```

### Database Migrations — Evolving Your Schema

Migrations are numbered SQL files that evolve your schema over time without losing data.

**File naming convention:**
```
migrations/
  001_create_users.up.sql
  001_create_users.down.sql
  002_add_payments.up.sql
  002_add_payments.down.sql
```

**`001_create_users.up.sql`:**
```sql
CREATE TABLE users (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    name       TEXT    NOT NULL,
    email      TEXT    NOT NULL UNIQUE,
    created_at TEXT    NOT NULL DEFAULT (datetime('now'))
);
```

**`002_add_payments.up.sql`:**
```sql
CREATE TABLE payments (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id    INTEGER NOT NULL REFERENCES users(id),
    title      TEXT    NOT NULL,
    amount     REAL    NOT NULL,
    created_at TEXT    NOT NULL DEFAULT (datetime('now'))
);
```

**Applying migrations in Go** (without an external library):

```go
func applyMigrations(db *sql.DB) error {
    _, err := db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (
        version INTEGER PRIMARY KEY,
        applied_at TEXT NOT NULL
    )`)
    if err != nil {
        return err
    }

    migrations := []struct {
        version int
        sql     string
    }{
        {1, `CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL, email TEXT NOT NULL UNIQUE
        )`},
        {2, `CREATE TABLE IF NOT EXISTS payments (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL REFERENCES users(id),
            title TEXT NOT NULL, amount REAL NOT NULL
        )`},
    }

    for _, m := range migrations {
        var exists int
        db.QueryRow("SELECT COUNT(*) FROM schema_migrations WHERE version=?", m.version).Scan(&exists)
        if exists > 0 {
            continue  // already applied
        }
        if _, err := db.Exec(m.sql); err != nil {
            return fmt.Errorf("migration %d: %w", m.version, err)
        }
        db.Exec("INSERT INTO schema_migrations (version, applied_at) VALUES (?, datetime('now'))", m.version)
    }
    return nil
}
```

For production use, `golang-migrate/migrate` automates this with file-based migrations, rollback support, and locking.

**Next:** [149-environment-and-config](../149-environment-and-config/README.md)
