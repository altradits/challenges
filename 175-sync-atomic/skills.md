# Skills for 175-sync-atomic

## What You'll Learn

**Previous:** [174-worker-pool](../174-worker-pool/skills.md) | **Next:** [176-linked-list](../176-linked-list/skills.md)

**Challenge:** Share a counter between goroutines with zero locks.

## Why Atomics?

A mutex around a hot counter serializes every goroutine. Atomic
operations use single CPU instructions — `Add`, `Load`, `Store`,
`CompareAndSwap` — that no thread can observe half-done. Use them for
simple shared values (counters, flags, sequence numbers); reach for a
mutex when the shared state is a struct with invariants.

## The Typed Wrappers (Go 1.19+)

```go
import "sync/atomic"

var v atomic.Int64
v.Store(10)        // set
v.Add(5)           // add, returns new value (15)
got := v.Load()    // read
old := v.Swap(0)   // set to 0, return previous
```

Same idea: `atomic.Bool`, `atomic.Int32/Uint64`, `atomic.Pointer[T]`,
`atomic.Value` (any type, one `Store` type forever). Copying a used
atomic is forbidden (`go vet` flags it) — always share by pointer.

## Compare-And-Swap Loops

`CompareAndSwap(old, new)` swaps only if the current value still equals
`old`, reporting success. Retry until you win:

```go
func CASAdd(addr *int64, delta int64) int64 {
    for {
        old := atomic.LoadInt64(addr)
        if atomic.CompareAndSwapInt64(addr, old, old+delta) {
            return old + delta
        }
    }
}
```

Between your `Load` and `CompareAndSwap` another goroutine may change
the value — then the swap fails and you retry with the fresh value.
This is how lock-free data structures (and `atomic.AddInt64` itself)
work internally.

## Solving the Challenge

- `Counter` holds one `atomic.Int64` field; `Add` calls `Add(n)`,
  `Value` calls `Load`. Constructor returns a pointer.
- `CASAdd` is the loop above. One subtlety: use the `sync/atomic`
  package functions (`atomic.LoadInt64`, …) on the plain `*int64`.
