# Skills for 42-context

## What You'll Learn

**Previous:** [26-select-and-sync](../26-select-and-sync/skills.md) | **Next:** [28-stringlength](../28-stringlength/skills.md)

**Challenge:** Propagate cancellation and deadlines with the `context` package.

## Why Context?

When a user cancels an HTTP request, or a deadline passes, you want all goroutines doing work for that request to stop — not keep consuming resources. The `context.Context` type propagates that cancellation signal through your call stack.

## Creating Contexts

```go
import "context"

// Root contexts — use these at the top level
ctx := context.Background()  // never cancelled, never expires
ctx := context.TODO()        // placeholder when you'll add proper context later
```

### WithCancel — Manual Cancellation

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()  // ALWAYS call cancel to release resources

go func() {
    <-ctx.Done()  // blocks until cancel() is called or parent is cancelled
    fmt.Println("cancelled:", ctx.Err())  // context.Canceled
}()

cancel()  // signals cancellation
```

### WithTimeout — Automatic Deadline

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

select {
case result := <-doWork():
    fmt.Println(result)
case <-ctx.Done():
    fmt.Println("timed out:", ctx.Err())  // context.DeadlineExceeded
}
```

### WithDeadline — Fixed Point in Time

```go
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

### WithValue — Attaching Data

```go
type key string
const userKey key = "user"

ctx = context.WithValue(ctx, userKey, "alice")

// Later, anywhere in the call chain
user, ok := ctx.Value(userKey).(string)
```

Use typed keys (not raw strings) to avoid collisions between packages.

## Checking for Cancellation

```go
// Non-blocking check
select {
case <-ctx.Done():
    return ctx.Err()  // context.Canceled or context.DeadlineExceeded
default:
    // context still alive, keep working
}
```

```go
// In a loop
for _, item := range items {
    if err := ctx.Err(); err != nil {
        return err  // context cancelled
    }
    process(item)
}
```

## Convention: Context as First Parameter

By convention, context is always the first parameter, named `ctx`:

```go
func FetchUser(ctx context.Context, id int) (*User, error) { ... }
func SaveRecord(ctx context.Context, rec Record) error      { ... }
```

Never store a context in a struct field — pass it through function calls.

## Solving the Challenge

```go
package piscine

import (
    "context"
    "time"
)

func FetchWithTimeout(ctx context.Context, workDuration time.Duration) (string, error) {
    done := make(chan struct{})
    go func() {
        time.Sleep(workDuration)
        close(done)
    }()
    select {
    case <-done:
        return "done", nil
    case <-ctx.Done():
        return "", ctx.Err()
    }
}

func ProcessAll(ctx context.Context, items []string, fn func(string)) error {
    for _, item := range items {
        if err := ctx.Err(); err != nil {
            return err
        }
        fn(item)
    }
    return nil
}
```

**Next:** [28-stringlength](../28-stringlength/README.md)
