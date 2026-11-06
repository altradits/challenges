# Prerequisites for 27-context

## Before You Start

- [25-channels](../25-channels/skills.md) — channel operations, select
- [26-select-and-sync](../26-select-and-sync/skills.md) — select with timeout pattern

## You're Ready When You Can...

- [ ] Create a context with `context.WithTimeout`
- [ ] Pass `ctx` as the first parameter and check `ctx.Done()` in a select
- [ ] Always call `cancel()` in a defer
- [ ] Explain the difference between `context.Canceled` and `context.DeadlineExceeded`

## Next Steps

- [28-stringlength](../28-stringlength/README.md)
