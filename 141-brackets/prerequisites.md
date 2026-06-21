# Prerequisites for 141-brackets

## Before You Start

### 1. Slice as a stack (push/pop)

A stack is the right tool for bracket matching — you push opening brackets and pop on closing brackets.

```go
stack := []rune{}
stack = append(stack, '(')          // push
top := stack[len(stack)-1]          // peek
stack = stack[:len(stack)-1]        // pop
```

Review: [09-slicesintro](../09-slicesintro/skills.md) — covers append, slice reslicing, nil vs empty.

### 2. `for...range` on a string produces runes

```go
for _, r := range s {
    // r is a rune (single Unicode code point)
    if r == '(' { ... }
}
```

Review: [81-checknumber](../81-checknumber/skills.md) — covers for...range on strings.

### 3. `os.Args[1:]` — iterate over all program arguments

```go
for _, arg := range os.Args[1:] {
    // check each argument
}
```

### 4. `switch` on rune — cleaner than nested if/else

```go
switch r {
case '(', '[', '{':
    // opening bracket — push
case ')', ']', '}':
    // closing bracket — pop and verify
}
```

Review: [142-rpncalc](../142-rpncalc/skills.md) — switch usage (even though 73 comes after, the syntax is the same).

### 5. Empty-stack check before popping

```go
if len(stack) == 0 {
    return false  // closing bracket with nothing to match
}
```

## You're Ready When You Can...

- [ ] Push a rune onto a `[]rune` slice and pop the last element
- [ ] Check `len(stack) == 0` before accessing `stack[len(stack)-1]`
- [ ] Iterate over `os.Args[1:]` to process all arguments
- [ ] Use a switch to match `'('`, `'['`, `'{'` opening brackets
- [ ] Return `false` on mismatch or non-empty stack at end

## Next Steps

- [142-rpncalc](../142-rpncalc/README.md)
