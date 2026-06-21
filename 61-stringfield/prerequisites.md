# Prerequisites for 61-stringfield

## Before You Start

### 1. `strings.Fields` — you already used this

```go
words := strings.Fields("hello world")  // ["hello", "world"]
```

Review: [27-wordcount](../27-wordcount/skills.md), [40-split](../40-split/skills.md)

### 2. `strings.Index` — find a separator

```go
idx := strings.Index("a,b,c", ",")  // 1
```

Review: [57-stringindex](../57-stringindex/skills.md)

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

- [62-stringmap](../62-stringmap/README.md)
