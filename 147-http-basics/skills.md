# Skills for 147-http-basics

## What You'll Learn

**Previous:** [146-encoding-json](../146-encoding-json/skills.md) | **Next:** [148-database-sql](../148-database-sql/skills.md)

**Challenge:** Write HTTP handlers, a router, and an HTTP client.

## The `net/http` Package

### Hello World HTTP Server

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World!")
    })
    http.ListenAndServe(":8080", nil)  // blocks forever
}
```

### `http.HandlerFunc`

The type of every handler function:

```go
type HandlerFunc func(ResponseWriter, *Request)
```

A `HandlerFunc` implements the `http.Handler` interface, which has one method:
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

### The Request Object

```go
func handler(w http.ResponseWriter, r *http.Request) {
    r.Method          // "GET", "POST", "PUT", "DELETE"
    r.URL.Path        // "/users/123"
    r.URL.Query()     // map of query params — r.URL.Query().Get("id")
    r.Header.Get("Authorization")  // request header

    // Reading the body (POST/PUT)
    var body MyStruct
    json.NewDecoder(r.Body).Decode(&body)
    defer r.Body.Close()
}
```

### The ResponseWriter

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Set headers BEFORE writing the body
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)  // 201 — call once, before Write

    // Write body
    json.NewEncoder(w).Encode(myStruct)
    // or: fmt.Fprintln(w, "text")
    // or: w.Write([]byte("bytes"))
}
```

**Rule:** Set headers → call `WriteHeader` → write body. Each step can only happen once.

### Common Status Codes

```go
http.StatusOK                  // 200
http.StatusCreated             // 201
http.StatusNoContent           // 204
http.StatusBadRequest          // 400
http.StatusUnauthorized        // 401
http.StatusForbidden           // 403
http.StatusNotFound            // 404
http.StatusInternalServerError // 500
```

### Routing with `http.ServeMux`

```go
mux := http.NewServeMux()

mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, `{"status":"ok"}`)
})

mux.HandleFunc("/users/", usersHandler)  // trailing slash: matches /users/anything

http.ListenAndServe(":8080", mux)
```

Go 1.22+ supports `"GET /path"` and `"POST /path"` patterns directly:
```go
mux.HandleFunc("GET /health", healthHandler)
mux.HandleFunc("POST /users", createUserHandler)
```

### JSON Response Helper

```go
func writeJSON(w http.ResponseWriter, status int, v any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(v)
}
```

### HTTP Client

```go
// Simple GET
resp, err := http.Get("https://api.example.com/users")
if err != nil {
    return err
}
defer resp.Body.Close()

var users []User
json.NewDecoder(resp.Body).Decode(&users)

// POST with body
body, _ := json.Marshal(user)
resp, err := http.Post(url, "application/json", bytes.NewReader(body))

// Full control
client := &http.Client{Timeout: 10 * time.Second}
req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
req.Header.Set("Authorization", "Bearer "+token)
resp, err = client.Do(req)
```

Always set a timeout on `http.Client`. The default has no timeout.

### Solving the Challenge

```go
package piscine

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

func JSONHandler(v any) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(v)
    }
}

func Router() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprintln(w, `{"status":"ok"}`)
    })
    mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        msg := r.URL.Query().Get("msg")
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"message": msg})
    })
    mux.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
        var body struct{ Text string `json:"text"` }
        json.NewDecoder(r.Body).Decode(&body)
        runes := []rune(body.Text)
        for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
            runes[i], runes[j] = runes[j], runes[i]
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"result": string(runes)})
    })
    _ = strings.ToLower // suppress import
    return mux
}

func FetchJSON(url string, v any) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    return json.NewDecoder(resp.Body).Decode(v)
}
```

**Next:** [148-database-sql](../148-database-sql/README.md)
