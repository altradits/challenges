# Skills for strlenindex

## What You'll Learn

**Previous:** [hello](../hello/skills.md)

**Challenge:** Understanding the relationship between len() and string indexes

## New Concepts Explained

### 1. String indexing with s[i]

Every character in a string can be accessed by its position:

```go
text := "stano"

text[0]  // 's'
text[1]  // 't'
text[2]  // 'a'
text[3]  // 'n'
text[4]  // 'o'
```

Indexes start at 0. This is not a convention you choose — it is how Go (and most languages) work.

### 2. What len() returns

`len(str)` returns the total number of bytes (characters, for ASCII strings) in the string. It counts from 1:

```go
text := "stano"
len(text)  // 5
```

The string has 5 characters. But the last character sits at index 4, not index 5. This gap — between the count and the last index — is one of the most common sources of off-by-one bugs.

### 3. Why last index = len - 1

```
s  t  a  n  o
0  1  2  3  4    ← indexes (start at 0)
1  2  3  4  5    ← len() counts (start at 1)
```

The last valid index is always `len(str) - 1`. There is no index equal to `len(str)` — accessing it causes a runtime panic.

### 4. Bytes vs characters: fmt.Printf("%c", ...)

When you do `text[0]`, Go gives you the raw byte value, not a character:

```go
text := "stano"
fmt.Println(text[0])          // prints: 115  (the byte value of 's')
fmt.Printf("%c\n", text[0])   // prints: s    (the character)
```

The `%c` verb in `fmt.Printf` tells Go to display the byte as its Unicode character instead of its numeric value.

```go
text := "stano"

fmt.Printf("The First character is: %c\n", text[0])  // s
fmt.Println(text[0])                                   // 115

fmt.Printf("The third character is: %c\n", text[2])  // a
fmt.Println(text[2])                                   // 97
```

### 5. The fmt verbs for strings

| Verb | Meaning | Example output |
|------|---------|---------------|
| `%c` | Print as character | `s` |
| `%d` | Print as decimal integer | `115` |
| `%v` | Default format | `115` (for byte) |
| `%s` | Print as string | only for string type |

## The Big Idea

`len()` counts from 1. Indexes count from 0. The last character is always at `len(str) - 1`. This relationship is fundamental to every string algorithm: loops, searches, sliding windows, and more.

**Next:** [printlastchar](../printlastchar/skills.md) — Applying len-1 to get the last character
