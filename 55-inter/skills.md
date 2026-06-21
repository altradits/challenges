# Skills for inter

## What You'll Learn

**Previous:** [54-hiddenp](../54-hiddenp/skills.md) | **Next:** [56-reversestrcap](../56-reversestrcap/README.md)

**Challenge:** Print characters that appear in BOTH strings, in the order they appear in the first string, with no duplicates.

## Core Concept: Set Intersection Using a Map

### What Is It?

Set intersection finds elements common to two collections. The efficient approach uses a **map as a set**: first record all characters from `s2` in a map, then walk `s1` and output only characters that exist in the map — and haven't been output yet.

### How It Works

**Step 1 — Build a "seen in s2" set:**

```go
inS2 := make(map[rune]bool)
for _, c := range s2 {
    inS2[c] = true
}
```

**Step 2 — Walk `s1`, output characters that are in both strings, no duplicates:**

You need a second map to track what has already been printed.

```go
printed := make(map[rune]bool)
result := ""
for _, c := range s1 {
    if inS2[c] && !printed[c] {
        result += string(c)
        printed[c] = true
    }
}
```

**Step 3 — Print the result:**

```go
fmt.Println(result)
```

**Full example — `"padinton"` and `"paqefwtdjetyiytjneytjoeyjnejeyj"`:**

`s2` contains: p, a, q, e, f, w, t, d, j, y, i, n, o

Walk `s1`:
- `p` → in s2, not printed → output `p`
- `a` → in s2, not printed → output `a`
- `d` → in s2, not printed → output `d`
- `i` → in s2, not printed → output `i`
- `n` → in s2, not printed → output `n`
- `t` → in s2, not printed → output `t`
- `o` → in s2, not printed → output `o`
- `n` → already printed → skip

Result: `padinto`

### Why a Map and Not a Nested Loop?

A naive nested loop — for each char in `s1`, scan all of `s2` — works but is O(len(s1) × len(s2)). With a map, lookup is O(1), making the whole algorithm O(len(s1) + len(s2)).

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using a nested loop without a "printed" set | Outputs duplicates when a char appears multiple times in s1 | Track printed chars in a second map |
| Using `s2` directly in the map | Working correctly only when s2 has no duplicates | Map membership automatically de-duplicates s2 |
| Printing chars in s2's order instead of s1's order | Wrong output order | Walk s1, not s2, for the output loop |

## Solving This Challenge

### Algorithm
1. Validate: exactly 2 args; print newline and return otherwise.
2. Build map `inS2` from all characters in `s2`.
3. Walk `s1`; for each character: if it's in `inS2` and not in `printed`, append to result and mark as printed.
4. Print result followed by newline.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [56-reversestrcap](../56-reversestrcap/README.md)
