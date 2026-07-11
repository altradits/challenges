# Skills for gin/02-routing

## What You'll Learn

**Previous:** [gin/01-hello-gin](../01-hello-gin/skills.md) | **Next:** [gin/03-middleware](../03-middleware/skills.md)

**Challenge:** Build a full REST user API with route groups, path params, query params, and JSON binding.

## Route Groups

Groups avoid repeating the prefix `/users` on every route:

```go
users := r.Group("/users")
{
    users.GET("", listUsers)
    users.GET("/:id", getUser)
    users.POST("", createUser)
    users.PUT("/:id", updateUser)
    users.DELETE("/:id", deleteUser)
}
```

The `{ }` block is optional style — the group is `users`, not the block.

## Binding JSON

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func createUser(c *gin.Context) {
    var u User
    if err := c.ShouldBindJSON(&u); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // u is now populated from the JSON body
}
```

`ShouldBindJSON` returns an error if the body is invalid JSON or required fields are missing. Never use `BindJSON` — it calls `c.AbortWithError` on failure, which is harder to test.

## Path Parameters

```go
func getUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    // find user with id
}
```

## Query Parameters

```go
func searchUsers(c *gin.Context) {
    q := c.Query("q")                   // "" if not provided
    limit := c.DefaultQuery("limit", "10")  // "10" if not provided
}
```

## Returning 404

```go
if user == nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
    return
}
```

## Solving the Challenge

```go
package main

import (
    "net/http"
    "strconv"
    "strings"
    "sync/atomic"
    "github.com/gin-gonic/gin"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var (
    users  []User
    nextID int64 = 1
)

func main() {
    r := gin.Default()

    u := r.Group("/users")
    u.GET("", func(c *gin.Context) {
        c.JSON(http.StatusOK, users)
    })

    u.POST("", func(c *gin.Context) {
        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        user.ID = int(atomic.AddInt64(&nextID, 1)) - 1
        users = append(users, user)
        c.JSON(http.StatusCreated, user)
    })

    u.GET("/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        for _, u := range users {
            if u.ID == id {
                c.JSON(http.StatusOK, u)
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
    })

    r.GET("/search", func(c *gin.Context) {
        q := strings.ToLower(c.Query("q"))
        var results []User
        for _, u := range users {
            if strings.Contains(strings.ToLower(u.Name), q) {
                results = append(results, u)
            }
        }
        c.JSON(http.StatusOK, results)
    })

    r.Run(":8080")
}
```

## Documentation

- [Gin routing](https://gin-gonic.com/docs/examples/routes-params/)
- [Gin model binding](https://gin-gonic.com/docs/examples/binding-models/)
- [Gin query params](https://gin-gonic.com/docs/examples/only-bind-query-string/)

**Next:** [gin/03-middleware](../03-middleware/README.md)
