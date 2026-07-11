# Skills for gorm/04-migrations

## What You'll Learn

**Previous:** [gorm/03-associations](../03-associations/skills.md) | **Next:** [gorm/05-advanced-queries](../05-advanced-queries/skills.md)

**Challenge:** Use golang-migrate to version your database schema with up/down migrations.

## Why Versioned Migrations?

GORM's `AutoMigrate` is great for development but dangerous in production:
- It never removes columns (your code may depend on columns that should be gone)
- It cannot rename columns or change types
- It has no rollback mechanism
- Multiple deploys can cause race conditions

**Versioned migrations** solve this:
- Each migration has a unique version number
- Applied in order; tracked in a `schema_migrations` table
- Each migration has an `up` (apply) and `down` (rollback) SQL file

## Migration File Naming

```
migrations/
  000001_create_users.up.sql
  000001_create_users.down.sql
  000002_add_bio_to_users.up.sql
  000002_add_bio_to_users.down.sql
```

## Migration File Contents

```sql
-- 000001_create_users.up.sql
CREATE TABLE users (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    name       TEXT NOT NULL,
    email      TEXT UNIQUE NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);

-- 000001_create_users.down.sql
DROP TABLE IF EXISTS users;

-- 000002_add_bio_to_users.up.sql
ALTER TABLE users ADD COLUMN bio TEXT;

-- 000002_add_bio_to_users.down.sql
-- SQLite doesn't support DROP COLUMN directly; usually recreate table
-- For PostgreSQL: ALTER TABLE users DROP COLUMN bio;
```

## Running Migrations

```go
import (
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/sqlite3"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func newMigrate(dbPath string) (*migrate.Migrate, error) {
    return migrate.New(
        "file://migrations",              // source: directory with .sql files
        "sqlite3://"+dbPath,              // database DSN
    )
}

func migrateUp(dbPath string) error {
    m, err := newMigrate(dbPath)
    if err != nil { return err }
    if err = m.Up(); err != nil && err != migrate.ErrNoChange {
        return err
    }
    return nil
}

func migrateDown(dbPath string, steps int) error {
    m, _ := newMigrate(dbPath)
    return m.Steps(-steps)
}

func currentVersion(dbPath string) (uint, error) {
    m, _ := newMigrate(dbPath)
    version, _, err := m.Version()
    return version, err
}
```

## GORM + Migrate Together

The typical pattern in production:

1. Use `golang-migrate` for schema migrations (run at startup or with a CLI command)
2. Use GORM for all data operations

```go
func main() {
    // Step 1: run migrations
    if err := migrateUp("app.db"); err != nil {
        log.Fatal("migration failed:", err)
    }

    // Step 2: open GORM connection (schema already set up)
    db, _ := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

    // Step 3: use db for CRUD
}
```

## Documentation

- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [GORM migrations](https://gorm.io/docs/migration.html)

**Next:** [gorm/05-advanced-queries](../05-advanced-queries/README.md)
