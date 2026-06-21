# Skills for 55-split

**Previous:** [54-replaceall](../54-replaceall/skills.md) | **Next:** [56-join](../56-join/skills.md)

**Challenge:** Implement `Split(s, sep string) []string` that splits a string by a separator, returning a slice of substrings.

## Core Concept: Tokenizing with a Start Pointer

### What Is It?

**Tokenizing** means breaking a string into meaningful pieces (tokens) by finding separator boundaries. The key insight is tracking a `start` pointer that marks where the current token began, then extracting `s[start:i]` whenever you hit a separator.

### The Algorithm

```go
func Split(s, sep string) []string {
    if sep == "" {
        // Split into individual characters
        var result []string
        for _, c := range s {
            result = append(result, string(c))
        }
        return result
    }

    var result []string
    start := 0
    sepLen := len(sep)

    for i := 0; i <= len(s)-sepLen; i++ {
        if s[i:i+sepLen] == sep {
            result = append(result, s[start:i])  // token before separator
            start = i + sepLen                    // new start is after separator
            i += sepLen - 1                       // skip past separator (loop will i++ too)
        }
    }
    result = append(result, s[start:])  // add the last token
    return result
}
```

### Visualizing "a,b,c" split by ","

```
s:     a , b , c
index: 0 1 2 3 4
start: 0

i=1: s[1:2]=="," → append s[0:1]="a", start=2, i=1
i=2: loop increments to 3
i=3: s[3:4]=="," → append s[2:3]="b", start=4, i=3
i=4: loop increments to 5 > 4 = len(s)-sepLen, loop ends
After loop: append s[4:]="c"

Result: ["a", "b", "c"]
```

### The Last Token Is Always Added After the Loop

If you only append when you find a separator, the text after the LAST separator is never added. The post-loop `append(result, s[start:])` handles this.

Example: `"hello"` with no separator — the loop body never runs, but `s[0:]` = `"hello"` is appended after the loop.

### What Happens with Consecutive Separators

`"a,,b"` split by `","`:
- At i=1: separator found, append `s[0:1]` = `"a"`, start=2
- At i=2: separator found, append `s[2:2]` = `""` (empty string), start=3
- After loop: append `s[3:]` = `"b"`
- Result: `["a", "", "b"]` — the empty token is preserved, matching how `strings.Split` behaves.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Forgetting the post-loop append | Last token silently dropped | Always `append(result, s[start:])` after loop |
| Loop bound too large | `s[i:i+sepLen]` panics past end | Use `i <= len(s)-sepLen` |
| Not handling empty separator | Infinite loop or wrong result | Check `sep == ""` first and split by character |

## Algorithm

1. If `sep == ""`, return slice of individual characters
2. Set `start = 0`, `result = []`
3. Scan `i` from `0` to `len(s)-len(sep)` inclusive:
   - If `s[i:i+len(sep)] == sep`: append `s[start:i]` to result, set `start = i+len(sep)`, advance `i` past separator
4. Append `s[start:]` (last token)
5. Return `result`

## The Challenge

See [README.md](README.md).

**Next:** [56-join](../56-join/README.md)
