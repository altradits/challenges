# gorm/05-advanced-queries — Advanced Queries and Transactions

## Challenge

Using the blog models from challenge 03, implement:

```go
// GetRecentPosts returns the N most recently created posts with their authors and tags.
func GetRecentPosts(db *gorm.DB, limit int) ([]Post, error)

// GetPostStats returns total posts, average comment count, and most popular tag.
func GetPostStats(db *gorm.DB) (PostStats, error)

// SearchPosts searches title and body with pagination.
func SearchPosts(db *gorm.DB, query string, page, size int) ([]Post, int64, error)

// CreatePostWithTags creates a post and finds-or-creates its tags in a transaction.
// If any step fails, the whole thing rolls back.
func CreatePostWithTags(db *gorm.DB, title, body string, userID uint, tagNames []string) (*Post, error)

// DeleteUserAndPosts deletes a user and all their posts in a transaction.
func DeleteUserAndPosts(db *gorm.DB, userID uint) error
```

Read `skills.md` before you start.
