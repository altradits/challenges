# Skills for gin/03-middleware

## What You'll Learn

**Previous:** [gin/02-routing](../02-routing/skills.md) | **Next:** [gin/04-json-crud-api](../04-json-crud-api/skills.md)

**Challenge:** Write custom Gin middleware for logging, CORS, rate limiting, and auth.

## What Is Middleware?

Middleware is a function that runs before (and optionally after) your handler. In Gin:

```go
func MyMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // before handler
        c.Next()   // call the next handler
        // after handler
    }
}
```

Apply globally or per-route:

```go
r.Use(MyMiddleware())        // all routes
r.GET("/path", MyMiddleware(), handler)  // single route
group.Use(MyMiddleware())    // route group
```

## Logger Middleware

```go
func RequestLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        log.Printf("%s %s %d %v",
            c.Request.Method,
            c.Request.URL.Path,
            c.Writer.Status(),
            time.Since(start),
        )
    }
}
```

## CORS Middleware

```go
func CORS() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,X-API-Key")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }
        c.Next()
    }
}
```

## Auth Middleware

```go
func APIKeyAuth(key string) gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.GetHeader("X-API-Key") != key {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "invalid or missing API key",
            })
            return
        }
        c.Next()
    }
}
```

`c.Abort()` stops the middleware chain — subsequent handlers and middleware do not run.

## Abort vs Return

```go
// Wrong: just return doesn't stop the chain
func bad(c *gin.Context) {
    if !authorized {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return   // handler stops, but middleware chain continues
    }
    c.Next()
}

// Correct: AbortWith* stops everything
func good(c *gin.Context) {
    if !authorized {
        c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
        return
    }
    c.Next()
}
```

## Solving the Challenge

```go
package main

import (
    "log"
    "net/http"
    "sync"
    "time"
    "github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        log.Printf("[%d] %s %s %v",
            c.Writer.Status(),
            c.Request.Method,
            c.Request.URL.Path,
            time.Since(start))
    }
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type,X-API-Key")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }
        c.Next()
    }
}

type ipEntry struct {
    count     int
    resetAt   time.Time
    mu        sync.Mutex
}

var (
    ipMap   = make(map[string]*ipEntry)
    ipMapMu sync.Mutex
)

func RateLimit(maxPerSec int) gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := c.ClientIP()
        ipMapMu.Lock()
        entry, ok := ipMap[ip]
        if !ok {
            entry = &ipEntry{resetAt: time.Now().Add(time.Second)}
            ipMap[ip] = entry
        }
        ipMapMu.Unlock()

        entry.mu.Lock()
        if time.Now().After(entry.resetAt) {
            entry.count = 0
            entry.resetAt = time.Now().Add(time.Second)
        }
        entry.count++
        tooMany := entry.count > maxPerSec
        entry.mu.Unlock()

        if tooMany {
            c.AbortWithStatusJSON(http.StatusTooManyRequests,
                gin.H{"error": "rate limit exceeded"})
            return
        }
        c.Next()
    }
}

func APIKeyAuth(key string) gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.GetHeader("X-API-Key") != key {
            c.AbortWithStatusJSON(http.StatusUnauthorized,
                gin.H{"error": "invalid API key"})
            return
        }
        c.Next()
    }
}

func main() {
    r := gin.New()
    r.Use(RequestLogger(), CORSMiddleware(), RateLimit(5))

    r.GET("/public", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "public endpoint"})
    })

    protected := r.Group("/api/v1/protected", APIKeyAuth("secret123"))
    {
        protected.GET("/data", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"secret": "protected data"})
        })
    }

    r.Run(":8080")
}
```

## Documentation

- [Gin middleware](https://gin-gonic.com/docs/examples/middleware/)
- [gin-contrib/cors](https://github.com/gin-contrib/cors) — production CORS middleware
- [gin-contrib/logger](https://github.com/gin-contrib/logger) — production logger

**Next:** [gin/04-json-crud-api](../04-json-crud-api/README.md)
