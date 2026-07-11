# Skills for gin/01-hello-gin

## What You'll Learn

**Previous:** [185-sliding-window](../../../185-sliding-window/skills.md) | **Next:** [gin/02-routing](../02-routing/skills.md)

**Challenge:** Build your first Gin server with three routes.

## Gin Basics

### Engine

```go
r := gin.Default()
```

`gin.Default()` creates an engine with two built-in middleware:
- **Logger** — prints each request to stdout
- **Recovery** — catches panics and returns 500

For production you often call `gin.New()` and add middleware manually.

### Registering Routes

```go
r.GET("/path", handlerFunc)
r.POST("/path", handlerFunc)
r.PUT("/path", handlerFunc)
r.DELETE("/path", handlerFunc)
r.PATCH("/path", handlerFunc)
```

### Handler Functions

```go
func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello, World!",
    })
}
```

`*gin.Context` carries the request and lets you send responses.
`gin.H` is a shortcut for `map[string]interface{}`.

### Path Parameters

```go
r.GET("/hello/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name + "!"})
})
```

`:name` captures anything in that position. Access with `c.Param("name")`.

### Starting the Server

```go
r.Run(":8080")   // listens on 0.0.0.0:8080
```

`Run` blocks. It wraps `http.ListenAndServe` and handles errors.

## Response Methods

| Method | Use |
|--------|-----|
| `c.JSON(code, obj)` | JSON response |
| `c.String(code, fmt, args)` | Plain text response |
| `c.HTML(code, name, obj)` | HTML template |
| `c.Redirect(code, url)` | HTTP redirect |
| `c.Status(code)` | Status only, no body |
| `c.AbortWithStatus(code)` | Stop middleware chain, send status |

## HTTP Status Codes

```go
http.StatusOK           // 200
http.StatusCreated      // 201
http.StatusBadRequest   // 400
http.StatusUnauthorized // 401
http.StatusNotFound     // 404
http.StatusInternalServerError // 500
```

Always import `"net/http"` for the constants. Don't use raw numbers.

## Solving the Challenge

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
    })

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    r.GET("/hello/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name + "!"})
    })

    r.Run(":8080")
}
```

## Common Mistakes

| Mistake | Fix |
|---------|-----|
| `gin.New()` without middleware | Use `gin.Default()` or add Logger+Recovery manually |
| Forgetting `go get github.com/gin-gonic/gin` | Run it first |
| Using raw int status codes `200` | Use `http.StatusOK` constants |
| Not running `go mod tidy` | Run it after `go get` |

## Documentation

- [Gin Quickstart](https://gin-gonic.com/docs/quickstart/)
- [gin.Context](https://pkg.go.dev/github.com/gin-gonic/gin#Context)
- [HTTP status codes](https://pkg.go.dev/net/http#StatusOK)

**Next:** [gin/02-routing](../02-routing/README.md)
