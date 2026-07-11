# Skills for gorm/03-associations

## What You'll Learn

**Previous:** [gorm/02-crud](../02-crud/skills.md) | **Next:** [gorm/04-migrations](../04-migrations/skills.md)

**Challenge:** Model one-to-many and many-to-many relationships. Load related data efficiently.

## Belongs To

A `Post` belongs to a `User`:

```go
type Post struct {
    gorm.Model
    Title  string
    UserID uint   // foreign key: must match User's primary key type
    User   User   // GORM auto-populates this when you use Preload
}
```

Convention: if the related model is `User` and the field is `User User`, GORM assumes the FK is `UserID`. Change with `foreignKey` tag.

## Has Many

A `User` has many `Post`s:

```go
type User struct {
    gorm.Model
    Name  string
    Posts []Post  // GORM uses UserID as the FK
}
```

## Many-to-Many

A `Post` can have many `Tag`s, and a `Tag` can belong to many `Post`s:

```go
type Post struct {
    gorm.Model
    Tags []Tag `gorm:"many2many:post_tags;"`
}

type Tag struct {
    gorm.Model
    Name  string
    Posts []Post `gorm:"many2many:post_tags;"`
}
```

GORM creates a junction table `post_tags` automatically.

## Eager Loading — Preload

```go
// Load a post with its author and comments
var post Post
db.Preload("User").Preload("Comments").First(&post, id)

// Load user with all posts and each post's tags
var user User
db.Preload("Posts.Tags").First(&user, userID)

// Load everything
db.Preload(clause.Associations).First(&post, id)
```

## Creating with Associations

```go
// Create user with posts in one go
user := User{
    Name: "Alice",
    Posts: []Post{
        {Title: "Post 1", Tags: []Tag{{Name: "go"}}},
        {Title: "Post 2"},
    },
}
db.Create(&user)
```

## Appending Associations

```go
// Add a tag to an existing post
tag := Tag{Name: "gorm"}
db.Create(&tag)
db.Model(&post).Association("Tags").Append(&tag)

// Remove a tag
db.Model(&post).Association("Tags").Delete(&tag)

// Replace all tags
db.Model(&post).Association("Tags").Replace([]Tag{{Name: "new"}})
```

## Querying Associations

```go
// Posts by user
var posts []Post
db.Where("user_id = ?", userID).Find(&posts)

// Posts with a specific tag (join)
var posts []Post
db.Joins("JOIN post_tags ON post_tags.post_id = posts.id").
    Joins("JOIN tags ON tags.id = post_tags.tag_id").
    Where("tags.name = ?", "go").
    Find(&posts)

// Count comments per post
var result struct { PostID uint; Count int64 }
db.Model(&Comment{}).
    Select("post_id, count(*) as count").
    Group("post_id").
    Scan(&result)
```

## Documentation

- [GORM associations](https://gorm.io/docs/associations.html)
- [GORM has many](https://gorm.io/docs/has_many.html)
- [GORM many-to-many](https://gorm.io/docs/many_to_many.html)
- [GORM preloading](https://gorm.io/docs/preload.html)

**Next:** [gorm/04-migrations](../04-migrations/README.md)
