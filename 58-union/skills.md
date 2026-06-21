# Skills for union

## What You'll Learn

**Previous:** [57-saveandmiss](../57-saveandmiss/skills.md) | **Next:** [59-wdmatch](../59-wdmatch/README.md)

**Challenge:** Print all characters that appear in either of the two input strings, in the order they appear in the combined input (s1 first, then s2), with no duplicates.

## Core Concept: Set Union Using a Map

### What Is It?

Set union combines two collections and keeps every element that appears in either one, with no duplicates. The efficient Go implementation uses a map to track which characters have already been output.

### How It Works

This is simpler than intersection: just walk both strings in order, printing each character the first time you see it.

**Step 1 — Create a "seen" set:**

```go
seen := make(map[rune]bool)
result := ""
```

**Step 2 — Walk s1, then s2 — output each new character once:**

```go
for _, c := range s1 {
    if !seen[c] {
        result += string(c)
        seen[c] = true
    }
}
for _, c := range s2 {
    if !seen[c] {
        result += string(c)
        seen[c] = true
    }
}
```

**Step 3 — Print:**

```go
fmt.Println(result)
```

**Full example — `"zpadinton"` and `"paqefwtdjetyiytjneytjoeyjnejeyj"`:**

Walk `s1 = "zpadinton"`:
- `z` → new → output
- `p` → new → output
- `a` → new → output
- `d` → new → output
- `i` → new → output
- `n` → new → output
- `t` → new → output
- `o` → new → output
- `n` → already seen → skip

Walk `s2`:
- `p` → seen → skip
- `a` → seen → skip
- `q` → new → output
- `e` → new → output
- `f` → new → output
- `w` → new → output
- `t` → seen → skip
- `d` → seen → skip
- `j` → new → output
- `e` → seen → skip
- ... `y` → new → output

Result: `zpadintoqefwjy` ✓

**Comparison: Union vs. Intersection**

| Operation | What you track | Walk order |
|-----------|---------------|------------|
| Intersection (55-inter) | "in s2" set + "printed" set | Walk s1, check both maps |
| Union (this) | "seen" set | Walk s1 then s2, one map |

Union is simpler: one map, two loops.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using two separate maps (one for s1, one for s2) | Outputs duplicates when the same char appears in both | One shared `seen` map covers both strings |
| Walking s2 before s1 | Output order wrong — s1 chars should come first | Always walk s1 completely before s2 |
| Printing `\n` when arg count is wrong | Check the requirement: print newline when wrong number of args | `fmt.Println()` with no args prints just a newline |

## Solving This Challenge

### Algorithm
1. If arg count != 2, print newline and return.
2. Create `seen := make(map[rune]bool)`.
3. Walk `s1`: for each char not in `seen`, append to result, mark seen.
4. Walk `s2`: same.
5. Print result + newline.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [59-wdmatch](../59-wdmatch/README.md)
