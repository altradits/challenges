# 01-gin — Gin Web Framework

Gin is the most popular Go web framework. It is fast, minimal, and idiomatic.

## Why Gin?

- Fastest HTTP router among major Go frameworks
- Clean API for binding JSON, forms, query params
- Built-in middleware chain
- Large ecosystem and employer familiarity

## Challenges

| Challenge | What You Build | Concept |
|-----------|---------------|---------|
| [01-hello-gin](./01-hello-gin/README.md) | "Hello World" API | Engine, routes, handlers |
| [02-routing](./02-routing/README.md) | Path params, groups | Route groups, params |
| [03-middleware](./03-middleware/README.md) | Logger, CORS, recovery | Middleware chains |
| [04-json-crud-api](./04-json-crud-api/README.md) | Tasks REST API | JSON binding, CRUD |
| [05-auth-jwt](./05-auth-jwt/README.md) | Login + protected routes | JWT middleware |
| [06-testing-gin](./06-testing-gin/README.md) | Test all endpoints | httptest, table tests |
| [07-production-gin](./07-production-gin/README.md) | Production-ready server | Graceful shutdown, config |

## Install

```bash
go get github.com/gin-gonic/gin
```

## Documentation

- [Gin GitHub](https://github.com/gin-gonic/gin)
- [Gin Docs](https://gin-gonic.com/docs/)
- [Gin Examples](https://github.com/gin-gonic/examples)
