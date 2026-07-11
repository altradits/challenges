# Prerequisites for 160-defer-panic-recover

## Before You Start

- [17-errorhandling](../17-errorhandling/README.md) — returning errors from functions
- [159-closures](../159-closures/README.md) — anonymous functions and closures
- [14-pointers](../14-pointers/README.md) — how Go passes values

## Key Concepts to Have Solid

### Named Return Values

```go
func divide(a, b int) (result int, err error) {
    result = a / b
    return  // returns result and err
}
```

Named returns let deferred functions modify what gets returned.

### Error Handling Pattern

```go
val, err := someFunc()
if err != nil {
    return 0, err
}
```

You must understand when to return an error vs. when to panic.

## You're Ready When You Can...

- [ ] Write a function that returns `(value, error)`
- [ ] Write an anonymous function and call it immediately
- [ ] Explain the difference between a syntax error and a runtime error
- [ ] Use `fmt.Errorf` to create an error with a message

## Next Steps

- [161-switch-statements](../161-switch-statements/README.md)
