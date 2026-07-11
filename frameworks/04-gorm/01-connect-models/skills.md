# Skills for gorm/01-connect-models

## What You'll Learn

**Previous:** [fiber/05-fiber-realtime](../../03-fiber/05-fiber-realtime/skills.md) | **Next:** [gorm/02-crud](../02-crud/skills.md)

**Challenge:** Connect GORM to SQLite, define models, auto-migrate, and seed data.

## Connecting

```go
import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
if err != nil {
    panic("failed to connect database: " + err.Error())
}
```

For PostgreSQL:

```go
import "gorm.io/driver/postgres"

dsn := "host=localhost user=postgres password=secret dbname=myapp port=5432 sslmode=disable"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

## `gorm.Model`

Embedding `gorm.Model` gives you four fields for free:

```go
type gorm.Model struct {
    ID        uint           `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`  // soft delete
}
```

With soft delete, `db.Delete(&user)` sets `DeletedAt` — the row stays but is excluded from future queries.

## Struct Tags

```go
type User struct {
    gorm.Model
    Name  string `gorm:"not null"`
    Email string `gorm:"uniqueIndex;not null"`
    Age   int    `gorm:"check:age > 0"`
    Role  string `gorm:"default:user;size:50"`
}
```

| Tag | Effect |
|-----|--------|
| `not null` | NOT NULL constraint |
| `uniqueIndex` | UNIQUE index |
| `default:val` | Default value |
| `size:N` | VARCHAR(N) |
| `index` | Regular index |
| `primaryKey` | Primary key |
| `column:name` | Column name override |
| `type:text` | Column type override |

## Auto-Migrate

```go
db.AutoMigrate(&User{}, &Post{}, &Tag{})
```

AutoMigrate:
- Creates tables that don't exist
- Adds missing columns
- Creates missing indexes
- Does NOT remove columns or change column types (safe for production)

## Creating Records

```go
user := User{Name: "Alice", Email: "alice@example.com"}
result := db.Create(&user)
// user.ID is now set
fmt.Println(result.RowsAffected, result.Error)
```

## Solving the Challenge

```go
package main

import (
    "fmt"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

type User struct {
    gorm.Model
    Name  string `gorm:"not null"`
    Email string `gorm:"uniqueIndex;not null"`
    Role  string `gorm:"default:user"`
}

type Post struct {
    gorm.Model
    Title  string `gorm:"not null"`
    Body   string
    UserID uint
    User   User
}

type Tag struct {
    gorm.Model
    Name string `gorm:"uniqueIndex"`
}

func main() {
    db, _ := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

    db.AutoMigrate(&User{}, &Post{}, &Tag{})

    users := []User{
        {Name: "Alice", Email: "alice@example.com"},
        {Name: "Bob", Email: "bob@example.com"},
    }
    db.Create(&users)

    db.Create(&[]Post{
        {Title: "Hello World", UserID: users[0].ID},
        {Title: "Go is great", UserID: users[0].ID},
        {Title: "GORM tutorial", UserID: users[1].ID},
    })

    db.Create(&[]Tag{{Name: "go"}, {Name: "gorm"}, {Name: "tutorial"}})

    tables := []string{"users", "posts", "tags"}
    for _, t := range tables {
        var count int64
        db.Raw(fmt.Sprintf("SELECT COUNT(*) FROM %s", t)).Scan(&count)
        fmt.Printf("Table %s: %d rows\n", t, count)
    }
}
```

## Documentation

- [GORM connecting to database](https://gorm.io/docs/connecting_to_the_database.html)
- [GORM models](https://gorm.io/docs/models.html)
- [GORM conventions](https://gorm.io/docs/conventions.html)
- [database/sql — standard library](https://pkg.go.dev/database/sql)

**Next:** [gorm/02-crud](../02-crud/README.md)
