# Prerequisites for 168-sync-mutex

## Before You Start

- [24-goroutines](../24-goroutines/README.md) — goroutines and WaitGroup
- [25-channels](../25-channels/README.md) — channels
- [14-pointers](../14-pointers/README.md) — pointer receivers (required for mutex types)

## You're Ready When You Can...

- [ ] Start a goroutine with `go func() {...}()`
- [ ] Use `sync.WaitGroup` to wait for goroutines
- [ ] Run the race detector: `go run -race main.go`
- [ ] Explain why `count++` is not thread-safe

## Next Steps

- [169-testing-advanced](../169-testing-advanced/README.md)
