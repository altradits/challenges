# 33. Zip String (Run-Length Encoding)

## What You'll Learn

This exercise teaches you **run-length encoding and pattern compression**. By the end, you'll understand:
- How to compress consecutive repeated characters
- Counting runs of identical characters
- Building compressed output format
- The concept of data compression

## Theory: Run-Length Encoding

### What is RLE?

Run-Length Encoding compresses consecutive identical characters:

```
"YouuungFellllas" → "1Y1o3u1n1g1F1e4l1a1s"
```

### The Pattern

For each group of identical characters:
1. Count how many times it repeats
2. Output: count + character

```
"aaa" → "3a"
"bb" → "2b"
"c" → "1c"
```

## Step-by-Step Approach

1. **Initialize** counter and result
2. **Iterate** through string
3. **If same as previous**: increment counter
4. **If different**: output count + previous, reset counter
5. **Handle** last group
6. **Return** result

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Forgetting last group | Last chars not output | Handle after loop |
| Wrong count format | `"a"` not `"1a"` | Always output count |
| Not resetting counter | Count accumulates wrong | Reset on change |

## Task

Write a function that compresses consecutive repeated characters.

### Expected Function

```go
func ZipString(s string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"YouuungFellllas"` | `"1Y1o3u1n1g1F1e4l1a1s"` | Mixed runs |
| `"Thee quuick"` | `"1T1h2e1 1q2u1i1c1k"` | With spaces |
| `"Helloo"` | `"1H1e2l2o"` | Simple runs |

## Knowledge Check

Before coding, make sure you can answer:
1. What is run-length encoding?
2. When do you output the count?
3. How do you handle the last group?

## Next Steps

After completing this, you'll be ready for:
- [61-saveandmiss](../61-saveandmiss/README.md) - Saveandmiss
- [62-reversestrcap](../62-reversestrcap/README.md) - Reversestrcap

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 60
```

or run the tests directly:

```bash
go test ./60-zipstring/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 60` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
