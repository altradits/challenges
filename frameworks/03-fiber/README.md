# 03-fiber — Fiber Web Framework

Fiber is inspired by Express.js. If you come from Node.js, Fiber feels familiar. It is built on fasthttp — the fastest HTTP implementation in Go.

## Why Fiber?

- Express.js-like API — gentle learning curve from JavaScript
- Built on fasthttp — extremely fast
- Low memory allocation
- Rich middleware ecosystem

## Note: Standard Library Compatibility

Fiber uses its own `fasthttp.RequestCtx` under the hood, **not** `net/http`. This means:
- Fiber is NOT compatible with `net/http` middleware (you cannot use `http.Handler` directly)
- Use Fiber for performance-critical APIs or when coming from Express.js

## Challenges

| Challenge | What You Build | Concept |
|-----------|---------------|---------|
| [01-hello-fiber](./01-hello-fiber/README.md) | Hello server | App, routes, Ctx |
| [02-fiber-routing](./02-fiber-routing/README.md) | REST API | Groups, params, body parsing |
| [03-fiber-middleware](./03-fiber-middleware/README.md) | Middleware stack | Logger, CORS, rate limit |
| [04-fiber-api](./04-fiber-api/README.md) | Full CRUD API | JSON, validation, error handling |
| [05-fiber-realtime](./05-fiber-realtime/README.md) | WebSocket chat | WebSocket upgrade, hub pattern |

## Install

```bash
go get github.com/gofiber/fiber/v2
```

## Documentation

- [Fiber Website](https://gofiber.io/)
- [Fiber GitHub](https://github.com/gofiber/fiber)
- [Fiber Docs](https://docs.gofiber.io/)
