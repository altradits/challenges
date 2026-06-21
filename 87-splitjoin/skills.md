# Skills for splitjoin

## What You'll Learn

**Previous:** [86-searchreplace](../86-searchreplace/skills.md) | **Next:** [88-wordanatomy](../88-wordanatomy/skills.md)

**Challenge:** Implement `Split` and `Join` manually — without using `strings.Split` or `strings.Join`.

## Core Concept: Building Split and Join From Scratch

### What Is Split?
`Split(s, sep)` divides a string into parts wherever `sep` appears, returning a slice of substrings. Understanding how to write it manually reveals what the standard library does internally.

```go
func Split(s string, sep string) []string {
    // Empty separator: each character becomes its own element
    if sep == "" {
        result := []string{}
        for _, c := range s {
            result = append(result, string(c))
        }
        return result
    }
    // Empty input: return slice with one empty string
    if s == "" {
        return []string{""}
    }
    result := []string{}
    current := ""
    i := 0
    for i < len(s) {
        if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
            result = append(result, current)
            current = ""
            i += len(sep)
        } else {
            current += string(s[i])
            i++
        }
    }
    result = append(result, current) // don't forget the last segment
    return result
}
```

### What Is Join?
`Join(arr, sep)` concatenates all elements of a slice with `sep` between them. The key is to add the separator between elements — not before the first or after the last.

```go
func Join(arr []string, sep string) string {
    if len(arr) == 0 {
        return ""
    }
    result := arr[0]
    for i := 1; i < len(arr); i++ {
        result += sep + arr[i]
    }
    return result
}
```

### The append Function
`append` adds elements to a slice and returns the new slice. Always reassign the result:
```go
parts := []string{}
parts = append(parts, "hello")  // parts == ["hello"]
parts = append(parts, "world")  // parts == ["hello", "world"]
```

### String Slicing for Separator Detection
Check if the separator appears at position `i` by slicing and comparing:
```go
if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
    // separator found at position i
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Forgetting the last segment in Split | Last part after the final separator is lost | `result = append(result, current)` after the loop |
| Adding separator before the first element in Join | Output starts with sep | Start with `result = arr[0]`, loop from index 1 |
| Not handling `sep == ""` in Split | Each character should become its own element | Special-case empty sep at the start |

## Solving This Challenge

### Split Algorithm
1. If `sep == ""`: return each character as its own slice element.
2. If `s == ""`: return `[]string{""}`.
3. Scan through `s`, collecting characters into `current`. When `sep` matches, append `current` to result and reset. After the loop, append the remaining `current`.

### Join Algorithm
1. If `arr` is empty, return `""`.
2. Start with `result = arr[0]`. Loop from index 1, appending `sep + arr[i]`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [88-wordanatomy](../88-wordanatomy/README.md)
