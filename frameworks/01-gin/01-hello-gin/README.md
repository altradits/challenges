# gin/01-hello-gin — Your First Gin Server

## Challenge

Build an HTTP server using Gin that responds to these endpoints:

```
GET /         → 200 { "message": "Hello, World!" }
GET /ping     → 200 { "message": "pong" }
GET /hello/:name → 200 { "message": "Hello, <name>!" }
```

## Setup

```bash
mkdir hello-gin && cd hello-gin
go mod init hello-gin
go get github.com/gin-gonic/gin
```

## Starter Code

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    // TODO: add routes here

    r.Run(":8080")
}
```

## Test

```bash
go run .

# In another terminal:
curl http://localhost:8080/
curl http://localhost:8080/ping
curl http://localhost:8080/hello/Alice
```

Expected:
```
{"message":"Hello, World!"}
{"message":"pong"}
{"message":"Hello, Alice!"}
```

Read `skills.md` before you start.
