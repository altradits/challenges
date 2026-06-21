# Prerequisites for 150-logging

## Before You Start

- [16-interfaces](../16-interfaces/skills.md) — http.Handler is an interface (RequestLogger wraps it)
- [147-http-basics](../147-http-basics/skills.md) — http.HandlerFunc, http.Handler
- [145-time](../145-time/skills.md) — time.Since(start) for duration

## You're Ready When You Can...

- [ ] Use `log.Printf` for simple logging
- [ ] Create a `slog.Logger` with `slog.NewJSONHandler` or `slog.NewTextHandler`
- [ ] Log structured key-value pairs with `logger.Info("msg", "key", value)`
- [ ] Build a middleware that wraps `http.Handler`

## Next Steps

- [151-generics](../151-generics/README.md)
