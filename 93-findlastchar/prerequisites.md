# Prerequisites for 93-findlastchar

## Before You Start

To solve this challenge you need to understand:

### 1. `for...range` with Index

From [91-findchar skills.md](../91-findchar/skills.md): capture both index and character:

```go
for i, c := range s {
    // i = byte position, c = the character
}
```

### 2. The `-1` Sentinel for "Not Found"

From [91-findchar skills.md](../91-findchar/skills.md): initialise a result variable to `-1` before the loop so that "not found" is the default return value:

```go
result := -1
// ...loop...
return result   // -1 if nothing was found
```

### 3. Update Without Returning

From [92-countchar skills.md](../92-countchar/skills.md): you know how to accumulate (`count++`) without returning early. The same principle applies here — instead of incrementing a counter, you **update an index**:

```go
for i, c := range s {
    if c == target {
        lastSeen = i   // update, but do NOT return
    }
}
return lastSeen
```

### 4. Contrast with FindChar

You already implemented `FindChar`, which returns on the first match. `FindLastChar` is different: it NEVER returns inside the loop — it updates a variable and returns only after the loop is complete.

## Review If Stuck

- [91-findchar skills.md](../91-findchar/skills.md) — `for...range` with index, `-1` sentinel
- [92-countchar skills.md](../92-countchar/skills.md) — scanning the entire string without early return

## You're Ready When You Can...

- [ ] Use `for i, c := range s` to get both position and character
- [ ] Initialise a variable to `-1` and update it inside a loop
- [ ] Explain why you must NOT `return i` inside the loop for this challenge
- [ ] State what the value of `lastIndex` will be if no match is found

## Next Steps

- [93-findlastchar skills.md](skills.md) — teaches the "keep updating" pattern
- [94-replacechar README](../94-replacechar/README.md) — next challenge
