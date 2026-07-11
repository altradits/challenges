# Skills for chi/01-hello-chi

## What You'll Learn

**Previous:** [gorm/06-gorm-api](../../04-gorm/06-gorm-api/skills.md) | **Next:** [chi/02-chi-routing](../02-chi-routing/skills.md)

**Challenge:** Build a minimal Chi server. Notice that it uses pure `net/http` — no custom context type.

## Chi vs Gin/Echo/Fiber

Chi is intentionally minimal:

| Feature | Gin | Chi |
|---------|-----|-----|
| Handler type | `gin.HandlerFunc` | `http.HandlerFunc` |
| Context | `*gin.Context` | `*http.Request` + `http.ResponseWriter` |
| JSON response | `c.JSON(...)` | `json.NewEncoder(w).Encode(...)` |
| Path param | `c.Param("id")` | `chi.URLParam(r, "id")` |

This means **every `net/http` handler works with Chi** without modification.

## Handler Pattern

```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "hello"})
}
```

Or with a helper:

```go
func writeJSON(w http.ResponseWriter, code int, v interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(v)
}
```

## Route Registration

```go
r.Get("/path", handler)
r.Post("/path", handler)
r.Put("/path", handler)
r.Delete("/path", handler)
r.Patch("/path", handler)
```

## URL Parameters

Chi uses `{param}` syntax (curly braces):

```go
r.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
    name := chi.URLParam(r, "name")
    writeJSON(w, http.StatusOK, map[string]string{"message": "Hello, " + name + "!"})
})
```

## Built-In Middleware

```go
import "github.com/go-chi/chi/v5/middleware"

r.Use(middleware.Logger)     // request logging
r.Use(middleware.Recoverer)  // panic recovery
r.Use(middleware.Compress(5)) // gzip
r.Use(middleware.RealIP)     // set RemoteAddr from X-Real-IP
r.Use(middleware.RequestID)  // add X-Request-ID
```

## Solving the Challenge

```go
package main

import (
    "encoding/json"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(v)
}

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        writeJSON(w, http.StatusOK, map[string]string{"message": "Hello from Chi!"})
    })

    r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
        writeJSON(w, http.StatusOK, map[string]string{"message": "pong"})
    })

    r.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
        name := chi.URLParam(r, "name")
        writeJSON(w, http.StatusOK, map[string]string{"message": "Hello, " + name + "!"})
    })

    http.ListenAndServe(":8080", r)
}
```

## Documentation

- [Chi GitHub](https://github.com/go-chi/chi)
- [net/http](https://pkg.go.dev/net/http)
- [encoding/json](https://pkg.go.dev/encoding/json)

**Next:** [chi/02-chi-routing](../02-chi-routing/README.md)
