# Skills for echo/01-hello-echo

## What You'll Learn

**Previous:** [gin/07-production-gin](../../01-gin/07-production-gin/skills.md) | **Next:** [echo/02-echo-routing](../02-echo-routing/skills.md)

**Challenge:** Build your first Echo server. Notice the differences from Gin.

## Echo vs Gin

| Feature | Gin | Echo |
|---------|-----|------|
| Handler type | `gin.HandlerFunc` | `echo.HandlerFunc` |
| Context | `*gin.Context` | `echo.Context` |
| JSON response | `c.JSON(code, obj)` | `c.JSON(code, obj)` |
| Path param | `c.Param("id")` | `c.Param("id")` |
| Query param | `c.Query("key")` | `c.QueryParam("key")` |
| Bind JSON | `c.ShouldBindJSON(&v)` | `c.Bind(&v)` |
| Return error | not idiomatic | `return c.JSON(...)` |

Echo handlers **return an error**, which Echo's error handler processes:

```go
e.GET("/ping", func(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
})
```

If you return a non-nil error, Echo calls its `HTTPErrorHandler` which by default returns a JSON error response.

## Query Parameters

```go
e.GET("/users", func(c echo.Context) error {
    page := c.QueryParam("page")   // string
    size := c.QueryParam("size")
    return c.JSON(http.StatusOK, map[string]interface{}{
        "page": page,
        "size": size,
    })
})
```

Or use struct binding:

```go
type Pagination struct {
    Page int `query:"page"`
    Size int `query:"size"`
}

var p Pagination
c.Bind(&p)  // automatically reads query params into struct
```

## Built-In Middleware

```go
import "github.com/labstack/echo/v4/middleware"

e.Use(middleware.Logger())   // request logging
e.Use(middleware.Recover())  // panic recovery
e.Use(middleware.CORS())     // CORS headers
e.Use(middleware.Gzip())     // gzip compression
e.Use(middleware.Secure())   // security headers
```

## Solving the Challenge

```go
package main

import (
    "net/http"
    "strconv"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{
            "message": "Hello from Echo!",
        })
    })

    e.GET("/ping", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
    })

    e.GET("/hello/:name", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{
            "message": "Hello, " + c.Param("name") + "!",
        })
    })

    e.GET("/users", func(c echo.Context) error {
        page, _ := strconv.Atoi(c.QueryParam("page"))
        size, _ := strconv.Atoi(c.QueryParam("size"))
        return c.JSON(http.StatusOK, map[string]interface{}{
            "page":  page,
            "size":  size,
            "users": []interface{}{},
        })
    })

    e.Logger.Fatal(e.Start(":8080"))
}
```

## Error Handling

Echo's error handler formats errors consistently:

```go
// Return a standard HTTP error
return echo.NewHTTPError(http.StatusNotFound, "user not found")

// Return with custom map
return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
    "error": "invalid input",
})
```

Custom error handler:

```go
e.HTTPErrorHandler = func(err error, c echo.Context) {
    he, ok := err.(*echo.HTTPError)
    if ok {
        c.JSON(he.Code, map[string]interface{}{"error": he.Message})
    } else {
        c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
}
```

## Documentation

- [Echo quickstart](https://echo.labstack.com/docs/quick-start)
- [echo.Context](https://pkg.go.dev/github.com/labstack/echo/v4#Context)
- [Echo middleware](https://echo.labstack.com/docs/middleware)

**Next:** [echo/02-echo-routing](../02-echo-routing/README.md)
