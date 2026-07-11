# echo/01-hello-echo — Your First Echo Server

## Challenge

Build an HTTP server using Echo that responds to:

```
GET /         → 200 { "message": "Hello from Echo!" }
GET /ping     → 200 { "message": "pong" }
GET /hello/:name → 200 { "message": "Hello, <name>!" }
GET /users?page=1&size=10 → 200 { "page": 1, "size": 10, "users": [] }
```

## Setup

```bash
mkdir hello-echo && cd hello-echo
go mod init hello-echo
go get github.com/labstack/echo/v4
```

## Starter Code

```go
package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // TODO: add routes

    e.Logger.Fatal(e.Start(":8080"))
}
```

Read `skills.md` before you start.
