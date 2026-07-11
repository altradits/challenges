# Skills for echo/02-echo-routing

## What You'll Learn

**Previous:** [echo/01-hello-echo](../01-hello-echo/skills.md) | **Next:** [echo/03-echo-middleware](../03-echo-middleware/skills.md)

**Challenge:** Use Echo's route groups and nested routes to build a blog API.

## Route Groups

```go
api := e.Group("/api/v1")

posts := api.Group("/posts")
posts.GET("", listPosts)
posts.POST("", createPost)
posts.GET("/:id", getPost)
posts.PUT("/:id", updatePost)
posts.DELETE("/:id", deletePost)

// Nested under /api/v1/posts/:id/comments
comments := posts.Group("/:id/comments")
comments.GET("", listComments)
comments.POST("", addComment)
```

Nested groups inherit the parent's prefix. The path param `:id` from the parent is accessible in child handlers with `c.Param("id")`.

## Struct Binding

```go
type CreatePostInput struct {
    Title    string `json:"title" validate:"required,min=1,max=200"`
    Body     string `json:"body"  validate:"required"`
    AuthorID int    `json:"author_id" validate:"required,min=1"`
}

func createPost(c echo.Context) error {
    var input CreatePostInput
    if err := c.Bind(&input); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    // input is populated
    return c.JSON(http.StatusCreated, post)
}
```

Echo's `Bind` reads JSON, form, and query params based on content type.

## Documentation

- [Echo routing](https://echo.labstack.com/docs/routing)
- [Echo binding](https://echo.labstack.com/docs/binding)

**Next:** [echo/03-echo-middleware](../03-echo-middleware/README.md)
