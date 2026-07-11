# 05-chi — Chi Router

Chi is an idiomatic, composable router for Go's `net/http`. Unlike Gin, Echo, and Fiber, Chi adds **zero abstraction** over `net/http` — you use standard `http.Handler` and `http.HandlerFunc` types everywhere.

## Why Chi?

- Fully compatible with `net/http` — any `net/http` middleware works
- Composable: nest routers and subrouters freely
- Zero external dependencies
- Perfect if you want minimal framework overhead

## Challenges

| Challenge | What You Build | Concept |
|-----------|---------------|---------|
| [01-hello-chi](./01-hello-chi/README.md) | Hello server | Router, standard handlers |
| [02-chi-routing](./02-chi-routing/README.md) | REST API | Subrouters, URL params |
| [03-chi-middleware](./03-chi-middleware/README.md) | Middleware stack | stdlib-compatible middleware |
| [04-chi-api](./04-chi-api/README.md) | Full API + context values | Context threading, stdlib patterns |

## Install

```bash
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
```

## Documentation

- [Chi GitHub](https://github.com/go-chi/chi)
- [Chi Examples](https://github.com/go-chi/chi/tree/master/_examples)
- [net/http — standard library](https://pkg.go.dev/net/http)
