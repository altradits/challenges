# gorm/03-associations — Associations (One-to-Many, Many-to-Many)

## Challenge

Model a blog system with associations:

```go
type User struct {
    gorm.Model
    Name  string
    Email string `gorm:"uniqueIndex"`
    Posts []Post
}

type Post struct {
    gorm.Model
    Title   string
    Body    string
    UserID  uint
    User    User
    Tags    []Tag `gorm:"many2many:post_tags;"`
    Comments []Comment
}

type Comment struct {
    gorm.Model
    Text   string
    PostID uint
    Post   Post
}

type Tag struct {
    gorm.Model
    Name  string `gorm:"uniqueIndex"`
    Posts []Post `gorm:"many2many:post_tags;"`
}
```

Implement:
1. Create a user with 2 posts, each with 2 comments
2. Attach tags to posts
3. Query a post with all its comments and author (eager loading)
4. Query all posts by a specific user
5. Query all posts with a specific tag
6. Count comments per post

Read `skills.md` before you start.
