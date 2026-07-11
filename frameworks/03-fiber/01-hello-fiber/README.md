# fiber/01-hello-fiber — Your First Fiber App

## Challenge

Build an HTTP server using Fiber:

```
GET /                  → 200 { "message": "Hello from Fiber!" }
GET /ping              → 200 { "message": "pong" }
GET /hello/:name       → 200 { "message": "Hello, <name>!" }
GET /echo?text=hello   → 200 { "echo": "hello" }
```

## Setup

```bash
mkdir hello-fiber && cd hello-fiber
go mod init hello-fiber
go get github.com/gofiber/fiber/v2
```

## Starter Code

```go
package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    // TODO: add routes here

    app.Listen(":8080")
}
```

## Test

```bash
go run .
curl http://localhost:8080/
curl http://localhost:8080/hello/World
curl "http://localhost:8080/echo?text=hello"
```

Read `skills.md` before you start.
