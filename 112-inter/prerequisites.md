# Prerequisites for 112-inter

## Before You Start

### 1. Using a Map as a Membership Set

Build a set from a string, then use it for O(1) lookup:

```go
s2 := "hello"
inS2 := make(map[rune]bool)
for _, c := range s2 {
    inS2[c] = true
}

// Now check membership:
fmt.Println(inS2['h'])  // true
fmt.Println(inS2['z'])  // false
```

### 2. Two Maps for Two Questions

This challenge requires answering two separate questions per character:
- "Is this character in s2?" → use `inS2` map
- "Have I already added this to result?" → use `seen` map

```go
inS2 := make(map[rune]bool)   // built from s2
seen := make(map[rune]bool)   // tracks what we've added

for _, c := range s1 {
    if inS2[c] && !seen[c] {
        seen[c] = true
        // add to result
    }
}
```

### 3. Order Matters: Iterate s1, Look Up in s2

The result must follow the order of s1. So:
- Iterate **s1** (to preserve s1's character order)
- Look up in **s2's set** (to filter)

If you iterated s2 instead, you'd get s2's order — wrong.

### 4. What "Intersection" Means

Set intersection: only elements that are in BOTH sets. If s1=`"cat"` and s2=`"act"`, the intersection is `"cat"` filtered to only chars in s2 = `"cat"` (all chars appear in both). But if s1=`"dog"` and s2=`"cat"`, intersection is `""` (no common characters).

## Review If Stuck

- [111-union](../111-union/skills.md) — covers maps as sets; this challenge adds a second map for deduplication

## You're Ready When You Can...

- [ ] Build a `map[rune]bool` lookup set from one string
- [ ] Check two conditions: `inS2[c] && !seen[c]`
- [ ] Explain why you iterate s1 (not s2) to preserve correct order
- [ ] Explain the difference between union ("OR") and intersection ("AND")

## Next Steps

- [113-stringbuilder](../113-stringbuilder/README.md)
