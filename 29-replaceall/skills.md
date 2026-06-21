# Skills for replaceall

## What You'll Learn

**Previous:** [../28-longestword/skills.md](../28-longestword/skills.md) | **Next:** [../30-searchreplace/README.md](../30-searchreplace/README.md)

**Challenge:** Manually replace all occurrences of a substring `old` in `text` with `new`, without using `strings.ReplaceAll`.

## Core Concept: Manual Substring Search and String Building

### What Is It?
To replace substrings manually, scan the input string looking for the start of `old`. When found, append `new` to the result and skip ahead by `len(old)`. When not found, copy the current character and advance by 1.

This teaches how substring replacement actually works internally.

### How It Works

```go
func ReplaceAll(text, old, new string) string {
    if old == "" {
        return text
    }
    result := ""
    i := 0
    for i < len(text) {
        // Check if old starts at position i
        if i+len(old) <= len(text) && text[i:i+len(old)] == old {
            result += new          // append the replacement
            i += len(old)          // skip past the matched part
        } else {
            result += string(text[i]) // copy current character
            i++
        }
    }
    return result
}
```

### Step-by-Step for `ReplaceAll("banana", "na", "NA")`
- i=0: `text[0:2]="ba"` ≠ `"na"` → copy `'b'`, i=1
- i=1: `text[1:3]="an"` ≠ `"na"` → copy `'a'`, i=2
- i=2: `text[2:4]="na"` == `"na"` → append `"NA"`, i=4
- i=4: `text[4:6]="na"` == `"na"` → append `"NA"`, i=6
- i=6: loop ends
- result: `"baNANA"`

### String Slicing
`text[i:j]` extracts bytes from index `i` up to (not including) `j`:
```go
s := "banana"
fmt.Println(s[2:4]) // "na"
fmt.Println(s[0:3]) // "ban"
```

### Checking String Equality
Two string slices can be compared directly with `==`:
```go
if text[i:i+len(old)] == old {
    // match found
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not guarding `i+len(old) <= len(text)` | Slice out of bounds panic | Always check the slice bounds before slicing |
| Not returning early when `old == ""` | Undefined behavior | Return `text` unchanged if `old` is empty |
| Forgetting `i += len(old)` after a match | Re-scans the same match, infinite or wrong loop | Skip the length of `old` after each replacement |

## Solving This Challenge

### Algorithm
1. If `old == ""`, return `text` unchanged.
2. Initialize `result := ""`, `i := 0`.
3. While `i < len(text)`:
   - If `text[i:i+len(old)] == old` (and bounds are safe): append `new`, advance `i` by `len(old)`.
   - Else: append `string(text[i])`, advance `i` by 1.
4. Return `result`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../30-searchreplace/README.md](../30-searchreplace/README.md)
