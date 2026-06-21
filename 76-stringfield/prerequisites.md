# Prerequisites for 76-stringfield

## Before You Start

### 1. `strings.Fields` — you already used this

```go
words := strings.Fields("hello world")  // ["hello", "world"]
```

Review: [42-wordcount](../42-wordcount/skills.md), [55-split](../55-split/skills.md)

### 2. `strings.Index` — find a separator

```go
idx := strings.Index("a,b,c", ",")  // 1
```

Review: [72-stringindex](../72-stringindex/skills.md)

### 3. Building a result slice with `append`

```go
result := []string{}
result = append(result, "part1")
result = append(result, "part2")
```

### 4. Advancing past a found position

```go
s = s[idx+len(sep):]  // chop off everything up to and including the separator
```

## You're Ready When You Can...

- [ ] Use `strings.Index` to find where a separator is
- [ ] Slice a string before and after a position
- [ ] Loop until no more separators are found

## Next Steps

- [77-stringmap](../77-stringmap/README.md)
