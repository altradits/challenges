# Skills for gorm/05-advanced-queries

## What You'll Learn

**Previous:** [gorm/04-migrations](../04-migrations/skills.md) | **Next:** [gorm/06-gorm-api](../06-gorm-api/skills.md)

**Challenge:** Write complex GORM queries with joins, aggregations, scopes, and transactions.

## Scopes — Reusable Query Modifiers

```go
func Paginate(page, size int) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        offset := (page - 1) * size
        return db.Offset(offset).Limit(size)
    }
}

func Active(db *gorm.DB) *gorm.DB {
    return db.Where("deleted_at IS NULL")
}

// Usage:
db.Scopes(Paginate(1, 10), Active).Find(&posts)
```

## Raw SQL

```go
// Scan into a struct
var stats struct {
    TotalPosts   int64
    AvgComments  float64
}
db.Raw(`
    SELECT COUNT(*) as total_posts,
           AVG(comment_count) as avg_comments
    FROM (
        SELECT p.id, COUNT(c.id) as comment_count
        FROM posts p
        LEFT JOIN comments c ON c.post_id = p.id
        GROUP BY p.id
    )
`).Scan(&stats)

// Exec for non-SELECT
db.Exec("UPDATE posts SET views = views + 1 WHERE id = ?", postID)
```

## Select with Computed Fields

```go
type PostWithCount struct {
    Post
    CommentCount int
}

var results []PostWithCount
db.Model(&Post{}).
    Select("posts.*, COUNT(comments.id) as comment_count").
    Joins("LEFT JOIN comments ON comments.post_id = posts.id").
    Group("posts.id").
    Order("comment_count DESC").
    Scan(&results)
```

## Transactions

```go
// Transaction — auto-commit or rollback
err := db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&post).Error; err != nil {
        return err  // rolls back
    }
    for _, tag := range tags {
        if err := tx.Create(&tag).Error; err != nil {
            return err  // rolls back everything
        }
    }
    return nil  // commits
})

// Manual transaction
tx := db.Begin()
defer func() {
    if r := recover(); r != nil {
        tx.Rollback()
    }
}()

if err := tx.Create(&post).Error; err != nil {
    tx.Rollback()
    return err
}
tx.Commit()
```

## Find or Create

```go
// FindOrCreate: find by Name, create with defaults if not found
var tag Tag
db.FirstOrCreate(&tag, Tag{Name: tagName})

// With defaults for new records only:
db.Where(Tag{Name: tagName}).
    Attrs(Tag{Name: tagName}).
    FirstOrCreate(&tag)
```

## Solving: CreatePostWithTags

```go
func CreatePostWithTags(db *gorm.DB, title, body string, userID uint, tagNames []string) (*Post, error) {
    var post Post
    err := db.Transaction(func(tx *gorm.DB) error {
        post = Post{Title: title, Body: body, UserID: userID}
        if err := tx.Create(&post).Error; err != nil {
            return err
        }
        for _, name := range tagNames {
            var tag Tag
            tx.FirstOrCreate(&tag, Tag{Name: name})
            tx.Model(&post).Association("Tags").Append(&tag)
        }
        return nil
    })
    return &post, err
}
```

## Documentation

- [GORM advanced query](https://gorm.io/docs/advanced_query.html)
- [GORM raw SQL](https://gorm.io/docs/sql_builder.html)
- [GORM transactions](https://gorm.io/docs/transactions.html)
- [GORM scopes](https://gorm.io/docs/scopes.html)

**Next:** [gorm/06-gorm-api](../06-gorm-api/README.md)
