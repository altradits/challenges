# Skills for 57-cameltosnakecase

**Previous:** [56-join](../56-join/skills.md) | **Next:** [58-itoa](../58-itoa/skills.md)

**Challenge:** Convert a valid camelCase string to snake_case (inserting `_` before each uppercase letter after the first). Return the string unchanged if it is not valid camelCase.

## Core Concept: `unicode.IsUpper` and Validation-Before-Transformation

### What Is It?

This challenge has two distinct phases: **validate** the input, then **transform** it. If validation fails, return the input unchanged. This pattern — validate first, transform only on valid input — is common in real-world Go code.

### Validation Rules for camelCase

A valid camelCase string must:
1. Contain **only letters** (no digits, no punctuation)
2. Have **at least one uppercase letter** (otherwise it's already lowercase, not camelCase)
3. Have **no consecutive uppercase letters** (`CAMELCase` is invalid)
4. **Not end** with an uppercase letter

```go
func isValidCamelCase(s string) bool {
    if len(s) == 0 {
        return false
    }
    hasUpper := false
    for i, c := range s {
        if !unicode.IsLetter(c) {
            return false  // rule 1: only letters
        }
        if unicode.IsUpper(c) {
            hasUpper = true
            // rule 3: no consecutive uppercase
            if i > 0 {
                prev := []rune(s)[i-1]
                if unicode.IsUpper(prev) {
                    return false
                }
            }
        }
    }
    // rule 2: must have at least one uppercase
    if !hasUpper {
        return false
    }
    // rule 4: must not end with uppercase
    runes := []rune(s)
    if unicode.IsUpper(runes[len(runes)-1]) {
        return false
    }
    return true
}
```

### The Transformation

Once the string is confirmed valid, iterate through it. Each time you find an uppercase letter (that is not the very first character), insert `_` before it:

```go
func CamelToSnakeCase(s string) string {
    if !isValidCamelCase(s) {
        return s
    }
    result := ""
    for i, c := range s {
        if unicode.IsUpper(c) && i > 0 {
            result += "_"
        }
        result += string(c)
    }
    return result
}
```

### Tracing Through `"camelToSnakeCase"`

| i | char | IsUpper? | i>0? | Action | result |
|---|------|---------|------|--------|--------|
| 0 | c | no | — | append 'c' | "c" |
| 1 | a | no | — | append 'a' | "ca" |
| 2 | m | no | — | append 'm' | "cam" |
| 3 | e | no | — | append 'e' | "came" |
| 4 | l | no | — | append 'l' | "camel" |
| 5 | T | yes | yes | append '_','T' | "camel_T" |
| 6 | o | no | — | append 'o' | "camel_To" |
| 7 | S | yes | yes | append '_','S' | "camel_To_S" |
| ... | | | | | "camel_To_Snake_Case" |

### `unicode.IsUpper` vs Manual Check

`unicode.IsUpper(c)` is the clean way. You can also use ASCII arithmetic for letters-only:

```go
// Manual ASCII check (only works for A-Z)
if c >= 'A' && c <= 'Z' {
    // uppercase
}
```

`unicode.IsUpper` works for all Unicode letters, including accented capitals.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not validating first | `"CAMELCase"` becomes `"_C_A_M_E_L_Case"` | Check validity before transforming |
| Adding `_` when `i == 0` | Leading underscore on first letter | Check `i > 0` |
| Returning lowercase | `"HelloWorld"` → `"hello_world"` (but expected `"Hello_World"`) | Do NOT force lowercase unless required |

## Algorithm

1. Validate: check only letters, has uppercase, no consecutive uppercase, does not end with uppercase
2. If invalid, return `s` unchanged
3. Iterate through `s`:
   - If `unicode.IsUpper(c)` and `i > 0`: append `"_"` to result
   - Append `string(c)` to result
4. Return result

## The Challenge

See [README.md](README.md).

**Next:** [58-itoa](../58-itoa/README.md)
