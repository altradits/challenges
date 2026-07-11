# 02-echo — Echo Web Framework

Echo is a high-performance, minimalist Go web framework with built-in support for data binding and validation.

## Why Echo?

- Built-in data validation with `go-playground/validator`
- Clean, expressive API
- Native WebSocket support
- Excellent documentation

## Challenges

| Challenge | What You Build | Concept |
|-----------|---------------|---------|
| [01-hello-echo](./01-hello-echo/README.md) | Hello server | Engine, routes, context |
| [02-echo-routing](./02-echo-routing/README.md) | REST routes | Groups, params, binding |
| [03-echo-middleware](./03-echo-middleware/README.md) | Logger, CORS, JWT | Middleware stack |
| [04-echo-validation](./04-echo-validation/README.md) | Validated API | Struct tags, custom validators |
| [05-echo-auth](./05-echo-auth/README.md) | Auth + sessions | JWT + cookie sessions |

## Install

```bash
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
```

## Documentation

- [Echo Website](https://echo.labstack.com/)
- [Echo GitHub](https://github.com/labstack/echo)
