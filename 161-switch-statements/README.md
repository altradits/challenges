# 161-switch-statements — Switch Statements

## Challenge

Implement in `package piscine`:

```go
// DayType returns "weekend" for Saturday/Sunday, "weekday" otherwise.
// Input: "Monday", "Tuesday", ..., "Sunday"
func DayType(day string) string

// Grade converts a numeric score (0-100) to a letter grade.
// 90-100: "A", 80-89: "B", 70-79: "C", 60-69: "D", below 60: "F"
func Grade(score int) string

// Describe takes any value and returns a description string:
// int → "integer: N", string → "string: S", bool → "boolean: true/false",
// anything else → "unknown type"
func Describe(v interface{}) string

// FizzBuzz returns "Fizz" for multiples of 3, "Buzz" for multiples of 5,
// "FizzBuzz" for multiples of both, or the number as a string otherwise.
func FizzBuzz(n int) string
```

## Examples

```
DayType("Saturday")  →  "weekend"
DayType("Monday")    →  "weekday"

Grade(95)  →  "A"
Grade(73)  →  "C"
Grade(55)  →  "F"

Describe(42)       →  "integer: 42"
Describe("hello")  →  "string: hello"
Describe(true)     →  "boolean: true"
Describe(3.14)     →  "unknown type"

FizzBuzz(15)  →  "FizzBuzz"
FizzBuzz(9)   →  "Fizz"
FizzBuzz(10)  →  "Buzz"
FizzBuzz(7)   →  "7"
```

Read `skills.md` before you start.
