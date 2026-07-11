# 04-gorm — GORM ORM

GORM is the most popular Go ORM. It handles database migrations, CRUD, associations, and complex queries.

## Why GORM?

- Auto-migrations (creates/updates tables from struct definitions)
- Convention over configuration
- Built-in hooks (BeforeCreate, AfterUpdate, etc.)
- Supports PostgreSQL, MySQL, SQLite, SQL Server

## Challenges

| Challenge | What You Build | Concept |
|-----------|---------------|---------|
| [01-connect-models](./01-connect-models/README.md) | DB connection + models | GORM setup, auto-migrate |
| [02-crud](./02-crud/README.md) | CRUD operations | Create, Read, Update, Delete |
| [03-associations](./03-associations/README.md) | Users → Posts → Comments | HasMany, BelongsTo |
| [04-migrations](./04-migrations/README.md) | Schema evolution | Migrations, constraints |
| [05-advanced-queries](./05-advanced-queries/README.md) | Complex SQL | Scopes, raw SQL, transactions |
| [06-gorm-api](./06-gorm-api/README.md) | Full Gin + GORM API | Repository pattern |

## Install

```bash
# GORM + SQLite driver (no external DB needed for learning)
go get gorm.io/gorm
go get gorm.io/driver/sqlite

# For PostgreSQL:
go get gorm.io/driver/postgres
```

## Documentation

- [GORM Website](https://gorm.io/)
- [GORM GitHub](https://github.com/go-gorm/gorm)
- [GORM Guides](https://gorm.io/docs/)
