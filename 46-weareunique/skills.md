# Skills for weareunique

## What You'll Learn

**Previous:** [../45-thirdtimeisacharm/skills.md](../45-thirdtimeisacharm/skills.md) | **Next:** [../47-zipstring/README.md](../47-zipstring/README.md)

**Challenge:** Count characters that appear in one string but not the other, counting each unique character only once.

## Core Concept: Using a Map as a Set for Deduplication

### What Is a Set?
A "set" holds unique values — no duplicates. Go doesn't have a built-in set type, but `map[rune]bool` works perfectly: the key is the value, and the presence of the key means "this item exists in the set."

```go
seen := map[rune]bool{}
seen['a'] = true
seen['b'] = true
seen['a'] = true  // duplicate — no effect, key already exists
fmt.Println(len(seen)) // 2
```

### The Algorithm
1. Build a set of all unique characters in `str1`.
2. Build a set of all unique characters in `str2`.
3. Count characters in set1 that are NOT in set2, plus characters in set2 that are NOT in set1.

```go
func WeAreUnique(str1, str2 string) int {
    if str1 == "" && str2 == "" {
        return -1
    }
    set1 := map[rune]bool{}
    set2 := map[rune]bool{}
    for _, c := range str1 {
        set1[c] = true
    }
    for _, c := range str2 {
        set2[c] = true
    }
    count := 0
    for c := range set1 {
        if !set2[c] {
            count++
        }
    }
    for c := range set2 {
        if !set1[c] {
            count++
        }
    }
    return count
}
```

### Map Lookup Returns Zero Value for Missing Keys
In Go, looking up a missing key in a `map[rune]bool` returns `false` (the zero value for bool). So `!set2[c]` is true when `c` is not in `set2`:
```go
seen := map[rune]bool{}
seen['a'] = true
fmt.Println(seen['a']) // true
fmt.Println(seen['b']) // false — key missing, zero value returned
```

### Step-by-Step for `WeAreUnique("foo", "boo")`
- str1 "foo": set1 = {'f':true, 'o':true}
- str2 "boo": set2 = {'b':true, 'o':true}
- In set1 but not set2: 'f' → count=1
- In set2 but not set1: 'b' → count=2
- Result: 2

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Counting character occurrences instead of unique chars | `"foo"` and `"boo"` would count `o` multiple times | Use a set (map) to deduplicate first |
| Not handling both-empty case | Should return -1 | Check `str1 == "" && str2 == ""` first |
| Forgetting to check characters unique to str2 | Only counts what str1 has that str2 doesn't | Also loop over set2 and check against set1 |

## Solving This Challenge

### Algorithm
1. If both strings are empty, return -1.
2. Build `set1` from unique chars of `str1`; `set2` from unique chars of `str2`.
3. Count chars in set1 not in set2, plus chars in set2 not in set1.
4. Return count.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../47-zipstring/README.md](../47-zipstring/README.md)
