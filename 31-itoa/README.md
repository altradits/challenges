# 31. Integer to String (Itoa)

## What You'll Learn

This exercise teaches you **converting numbers to strings manually**. By the end, you'll understand:
- How to extract digits from an integer
- Converting digits to their character representation
- Handling negative numbers
- Building strings from numeric components

## Theory: Integer to String Conversion

### How Itoa Works

The `Itoa` function converts an integer to its string representation:

```
12345 → "12345"
-1234 → "-1234"
0 → "0"
```

### Extracting Digits

To extract digits from right to left:

```go
n := 12345
for n > 0 {
    digit := n % 10    // Get last digit: 5
    n = n / 10         // Remove last digit: 1234
}
```

### Converting Digit to Character

```go
digit := 5
char := '0' + digit   // '0' (48) + 5 = '5' (53)
```

### Handling Negative Numbers

```go
if n < 0 {
    n = -n
    // Add '-' prefix later
}
```

## Step-by-Step Approach

1. **Handle** zero as special case
2. **Check** if negative, remember sign
3. **Extract** digits in reverse order
4. **Convert** each digit to character
5. **Reverse** the result
6. **Add** sign if needed

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Forgetting negative | `-123` becomes `"123"` | Handle sign separately |
| Wrong digit conversion | `'0' + 5` gives `'5'` | Use `'0' + digit` |
| Not reversing | Digits come out backwards | Build reversed, then reverse |

## The Challenge

Write a function that converts an integer to a string.

### Expected Function

```go
func Itoa(n int) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `12345` | `"12345"` | Positive number |
| `0` | `"0"` | Zero |
| `-1234` | `"-1234"` | Negative number |
| `987654321` | `"987654321"` | Large number |

## Knowledge Check

Before coding, make sure you can answer:
1. How do you extract the last digit of a number?
2. How do you convert a digit (0-9) to its character?
3. Why do we need to reverse the digits?

## Next Steps

After completing this, you'll be ready for:
- **32-thirdchar**: Extracting every nth character
- **33-zipstring**: Run-length encoding
