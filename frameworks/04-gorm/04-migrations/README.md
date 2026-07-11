# gorm/04-migrations — Schema Migrations

## Challenge

Use `golang-migrate/migrate` alongside GORM to manage schema migrations:

```bash
go get -u github.com/golang-migrate/migrate/v4
go get -u github.com/golang-migrate/migrate/v4/database/sqlite3
go get -u github.com/golang-migrate/migrate/v4/source/file
```

1. Create migration SQL files:
   - `migrations/000001_create_users.up.sql` — create users table
   - `migrations/000001_create_users.down.sql` — drop users table
   - `migrations/000002_add_bio_to_users.up.sql` — add `bio TEXT` column
   - `migrations/000002_add_bio_to_users.down.sql` — remove `bio` column

2. Write a `migrateUp()` function that applies all pending migrations
3. Write a `migrateDown(steps int)` function that rolls back N steps
4. Write a `currentVersion()` function that returns the current migration version

Read `skills.md` before you start.
