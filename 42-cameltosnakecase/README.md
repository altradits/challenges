# 30. Camel to Snake Case

## What You'll Learn

This exercise teaches you **advanced string transformation and validation**. By the end, you'll understand:
- How to detect and transform naming conventions
- Combining multiple string operations
- Validation logic for string patterns
- Real-world text processing challenges

## Theory: Naming Conventions

### Camel Case

Camel case joins words without spaces, capitalizing each word after the first:

```
lowerCamelCase: helloWorld, countVowels
UpperCamelCase: HelloWorld, CountVowels (PascalCase)
```

### Snake Case

Snake case uses lowercase with underscores:

```
hello_world, count_vowels
```

### Validation Rules for CamelCase

A valid camelCase string:
1. Contains only letters (no numbers, no punctuation)
2. Does NOT end with an uppercase letter
3. Does NOT have consecutive uppercase letters
4. Has at least one uppercase letter (otherwise it's just lowercase)

## Step-by-Step Approach

1. **Validate** the input is valid camelCase
2. **If invalid**: return original string
3. **Iterate** through characters
4. **When uppercase found (not first)**: add `_` before it
5. **Convert** to lowercase
6. **Return** result

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Not validating | Would transform invalid strings | Check rules first |
| Adding `_` at start | First char shouldn't have `_` | Check `i > 0` |
| Forgetting lowercase | Should be all lowercase | Convert each char |

## The Challenge

Write a function that converts camelCase to snake_case.

### Expected Function

```go
func CamelToSnakeCase(s string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"HelloWorld"` | `"Hello_World"` | UpperCamelCase |
| `"helloWorld"` | `"hello_World"` | lowerCamelCase |
| `"camelCase"` | `"camel_Case"` | Basic case |
| `"CAMELtoSnackCASE"` | `"CAMELtoSnackCASE"` | Invalid (consecutive caps) |
| `"camelToSnakeCase"` | `"camel_To_Snake_Case"` | Multiple words |
| `"hey2"` | `"hey2"` | Invalid (has number) |

## Knowledge Check

Before coding, make sure you can answer:
1. What makes a string valid camelCase?
2. When should you add an underscore?
3. Why is validation important?

## Next Steps

After completing this, you'll be ready for:
- [43-itoa](../43-itoa/README.md) - Itoa
- [44-thirdchar](../44-thirdchar/README.md) - Thirdchar