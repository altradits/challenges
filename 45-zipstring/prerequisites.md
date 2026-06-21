# Prerequisites for 45-zipstring

## Before You Start

### 1. Converting a String to a Rune Slice

`[]rune(s)` converts a string into a slice of individual characters (runes). This lets you access any character by index safely:

```go
s := "Hello"
runes := []rune(s)
fmt.Println(runes[0])         // 72 (code point for 'H')
fmt.Println(string(runes[0])) // "H"
fmt.Println(len(runes))       // 5
```

### 2. Comparing Runes

Runes can be compared directly with `==`:

```go
a := 'H'
b := 'H'
c := 'e'
fmt.Println(a == b)  // true
fmt.Println(a == c)  // false
```

### 3. Tracking State Across Iterations

Use variables declared before the loop to remember information from previous iterations:

```go
prev := rune(0)
for _, c := range "aabbcc" {
    if c == prev {
        fmt.Println("same as previous")
    } else {
        fmt.Println("new character:", string(c))
        prev = c
    }
}
```

### 4. `fmt.Sprintf` for Formatting

`fmt.Sprintf` creates a string using format verbs — `%d` for integers, `%c` for rune characters:

```go
count := 3
char := 'u'
s := fmt.Sprintf("%d%c", count, char)
fmt.Println(s)  // "3u"
```

Alternatively: `strconv.Itoa(count) + string(char)`.

### 5. The "Flush After Loop" Pattern

When you are processing runs of identical items, you always need to flush the last run after the loop, because the loop only flushes when it detects a change:

```go
current := runes[0]
count := 1
for i := 1; i < len(runes); i++ {
    if runes[i] == current {
        count++
    } else {
        fmt.Println(count, string(current))  // flush
        current = runes[i]
        count = 1
    }
}
fmt.Println(count, string(current))  // flush LAST group here
```

## Review If Stuck

- [43-itoa](../43-itoa/skills.md) — covers converting integers to strings
- [44-thirdchar](../44-thirdchar/skills.md) — covers index-based character selection

## You're Ready When You Can...

- [ ] Convert a string to `[]rune` and index into it
- [ ] Compare two runes with `==`
- [ ] Track a "current" variable and count across loop iterations
- [ ] Use `fmt.Sprintf("%d%c", count, char)` to format a count+character pair
- [ ] Remember to flush the last group after the loop ends

## Next Steps

- [46-saveandmiss](../46-saveandmiss/README.md)
