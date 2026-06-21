# Prerequisites for 152-graceful-shutdown

## Before You Start

- [27-context](../27-context/skills.md) — context.WithTimeout is the shutdown deadline
- [147-http-basics](../147-http-basics/skills.md) — http.Server, http.Handler
- [24-goroutines](../24-goroutines/skills.md) — ListenAndServe runs in a goroutine
- [25-channels](../25-channels/skills.md) — signal channel, errCh pattern

## You're Ready When You Can...

- [ ] Create `chan os.Signal` and call `signal.Notify` for SIGINT + SIGTERM
- [ ] Start `srv.ListenAndServe()` in a goroutine and handle `http.ErrServerClosed`
- [ ] Call `srv.Shutdown(ctx)` with a timeout context
- [ ] Combine `select` to wait for either a startup error or a shutdown signal

## Next Steps

- [153-financial-freedom-api](../153-financial-freedom-api/README.md)
