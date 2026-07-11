# gorm/01-connect-models — Connect and Define Models

## Challenge

Connect to a SQLite database and define three models:

```go
type User struct {
    gorm.Model
    Name     string `gorm:"not null"`
    Email    string `gorm:"uniqueIndex;not null"`
    Role     string `gorm:"default:user"`
}

type Post struct {
    gorm.Model
    Title   string `gorm:"not null"`
    Body    string
    UserID  uint
    User    User
}

type Tag struct {
    gorm.Model
    Name string `gorm:"uniqueIndex"`
}
```

Tasks:
1. Connect to `./app.db` (SQLite file)
2. Auto-migrate all three models
3. Seed 2 users, 3 posts, and 3 tags
4. Print the database schema (table names and column counts)

## Setup

```bash
go mod init gorm-demo
go get gorm.io/gorm
go get gorm.io/driver/sqlite
```

Read `skills.md` before you start.
