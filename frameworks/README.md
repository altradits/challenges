# Go Frameworks

This section covers the most widely used Go web frameworks and libraries, from your first "hello world" to production-ready applications.

## Prerequisites

Complete the core challenges (01-27) and at least the beginner algorithm challenges before starting here. You should be comfortable with:

- Goroutines and channels (24-25)
- Interfaces (16)
- Error handling (17)
- JSON encoding (146-encoding-json)
- HTTP basics (147-http-basics)

## Framework Map

| Framework | Focus | Best For |
|-----------|-------|----------|
| [01-gin](./01-gin/README.md) | HTTP web framework | REST APIs, high performance |
| [02-echo](./02-echo/README.md) | HTTP web framework | Clean APIs with built-in validation |
| [03-fiber](./03-fiber/README.md) | HTTP web framework | Express.js-like API, ultra-fast |
| [04-gorm](./04-gorm/README.md) | ORM | Database models and migrations |
| [05-chi](./05-chi/README.md) | HTTP router | Lightweight, stdlib-compatible |

## Learning Path

### Beginner (Week 1-2)

1. [01-gin/01-hello-gin](./01-gin/01-hello-gin/README.md) — Your first Gin server
2. [01-gin/02-routing](./01-gin/02-routing/README.md) — Path params, query params
3. [05-chi/01-hello-chi](./05-chi/01-hello-chi/README.md) — Lightweight stdlib-based routing

### Intermediate (Week 3-4)

4. [01-gin/03-middleware](./01-gin/03-middleware/README.md) — Logging, CORS, recovery
5. [01-gin/04-json-crud-api](./01-gin/04-json-crud-api/README.md) — Full CRUD REST API
6. [04-gorm/01-connect-models](./04-gorm/01-connect-models/README.md) — Database with GORM
7. [04-gorm/02-crud](./04-gorm/02-crud/README.md) — GORM CRUD operations
8. [02-echo/01-hello-echo](./02-echo/01-hello-echo/README.md) — Echo framework basics

### Advanced (Week 5-6)

9. [01-gin/05-auth-jwt](./01-gin/05-auth-jwt/README.md) — JWT authentication
10. [04-gorm/03-associations](./04-gorm/03-associations/README.md) — One-to-many, many-to-many
11. [02-echo/04-echo-validation](./02-echo/04-echo-validation/README.md) — Input validation
12. [03-fiber/04-fiber-api](./03-fiber/04-fiber-api/README.md) — Ultra-fast REST API

### Mastery (Week 7+)

13. [01-gin/06-testing-gin](./01-gin/06-testing-gin/README.md) — Testing HTTP handlers
14. [04-gorm/05-advanced-queries](./04-gorm/05-advanced-queries/README.md) — Complex queries
15. [01-gin/07-production-gin](./01-gin/07-production-gin/README.md) — Production patterns
16. [03-fiber/05-fiber-realtime](./03-fiber/05-fiber-realtime/README.md) — WebSockets

## Installing Frameworks

```bash
# Gin
go get github.com/gin-gonic/gin

# Echo
go get github.com/labstack/echo/v4

# Fiber
go get github.com/gofiber/fiber/v2

# GORM + SQLite driver
go get gorm.io/gorm
go get gorm.io/driver/sqlite

# Chi
go get github.com/go-chi/chi/v5

# JWT (for auth challenges)
go get github.com/golang-jwt/jwt/v5
```

## Documentation

- [Gin](https://gin-gonic.com/docs/)
- [Echo](https://echo.labstack.com/docs)
- [Fiber](https://docs.gofiber.io/)
- [GORM](https://gorm.io/docs/)
- [Chi](https://go-chi.github.io/chi/)
