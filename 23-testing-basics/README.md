# 38-testing-basics — Testing in Go

## Challenge

Write a complete `solution_test.go` for the following functions (which you also implement):

```go
// Abs returns the absolute value of n.
func Abs(n int) int

// Palindrome reports whether s reads the same forwards and backwards.
func Palindrome(s string) bool

// FizzBuzz returns "Fizz" (div by 3), "Buzz" (div by 5), "FizzBuzz" (both), else the number as string.
func FizzBuzz(n int) string
```

Your test file must include:
- At least one table-driven test with `t.Run`
- At least one test for an edge case (empty string, zero, negative)
- At least one benchmark (`func BenchmarkXxx`)

Read `skills.md` before you start.
