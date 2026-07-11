# chi/01-hello-chi — Your First Chi Server

## Challenge

Build an HTTP server using Chi:

```
GET /         → 200 { "message": "Hello from Chi!" }
GET /ping     → 200 { "message": "pong" }
GET /hello/{name} → 200 { "message": "Hello, <name>!" }
```

## Setup

```bash
mkdir hello-chi && cd hello-chi
go mod init hello-chi
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
```

## Starter Code

```go
package main

import (
    "encoding/json"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // TODO: add routes here

    http.ListenAndServe(":8080", r)
}
```

Read `skills.md` before you start.
