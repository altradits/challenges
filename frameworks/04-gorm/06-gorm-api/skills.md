# Skills for gorm/06-gorm-api

## What You'll Learn

**Previous:** [gorm/05-advanced-queries](../05-advanced-queries/skills.md) | **Next:** [frameworks/05-chi](../../05-chi/README.md)

**Challenge:** Wire Gin and GORM together with the repository pattern into a production-quality API.

## Architecture

```
cmd/
  main.go              ← startup: DB, router, server
internal/
  models/
    user.go
    post.go
    comment.go
    tag.go
  repository/
    user_repo.go       ← GORM queries
    post_repo.go
  handler/
    auth_handler.go    ← Gin handlers
    post_handler.go
  middleware/
    auth.go            ← JWT middleware
  router/
    router.go          ← route registration
```

## Handler → Repository → DB Flow

```go
// Handler (no DB knowledge)
func (h *PostHandler) Create(c *gin.Context) {
    var input CreatePostInput
    c.ShouldBindJSON(&input)
    userID := c.GetUint("user_id")  // from JWT middleware

    post, err := h.repo.Create(userID, input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, post)
}

// Repository (no HTTP knowledge)
type PostRepository struct{ db *gorm.DB }

func (r *PostRepository) Create(userID uint, input CreatePostInput) (*Post, error) {
    post := &Post{Title: input.Title, Body: input.Body, UserID: userID}
    err := r.db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(post).Error; err != nil { return err }
        for _, name := range input.Tags {
            var tag Tag
            tx.FirstOrCreate(&tag, Tag{Name: name})
            tx.Model(post).Association("Tags").Append(&tag)
        }
        return nil
    })
    return post, err
}
```

## Authorization Patterns

```go
// Author-only middleware
func AuthorOnly(repo *PostRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        postID, _ := strconv.Atoi(c.Param("id"))
        userID := c.GetUint("user_id")

        post, err := repo.GetByID(uint(postID))
        if err != nil {
            c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "post not found"})
            return
        }
        if post.UserID != userID {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "not your post"})
            return
        }
        c.Set("post", post)
        c.Next()
    }
}
```

## Dependency Injection

```go
func main() {
    db, _ := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
    db.AutoMigrate(&User{}, &Post{}, &Comment{}, &Tag{})

    userRepo := &UserRepository{db: db}
    postRepo := &PostRepository{db: db}

    authHandler  := &AuthHandler{userRepo: userRepo}
    postHandler  := &PostHandler{repo: postRepo}

    r := gin.Default()
    router.Setup(r, authHandler, postHandler)
    r.Run(":8080")
}
```

## Documentation

- [GORM + Gin example](https://github.com/gin-gonic/examples/tree/master/gorm)
- [Clean architecture in Go](https://github.com/bxcodec/go-clean-arch)
- [Repository pattern](https://martinfowler.com/eaaCatalog/repository.html)

**Next:** [frameworks/05-chi](../../05-chi/README.md)
