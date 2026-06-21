# Skills for options

## What You'll Learn

**Previous:** [68-itoabase](../68-itoabase/skills.md) | **Next:** [70-piglatin](../70-piglatin/README.md)

**Challenge:** Parse command-line flags (`-abc`, `-z`) and store them as bits in a single `int32`. Each letter `a`–`z` corresponds to a bit position (bit 0 for `a`, bit 25 for `z`). Print the result as 4 groups of 8 bits.

## Core Concept: Bitwise Flag Storage

### What Is It?

Bitwise flags pack multiple boolean values (on/off) into a single integer. Each letter `a`–`z` maps to a bit: `a` = bit 0, `b` = bit 1, ..., `z` = bit 25. Setting a flag with `|=` and checking it with `&` are the core operations.

### How It Works

**Step 1 — Check for `-h` flag first (it has priority):**

```go
var flags int32

// Check if -h appears as a standalone argument or at the start of any argument
for _, arg := range os.Args[1:] {
    if len(arg) > 1 && arg[0] == '-' {
        for _, c := range arg[1:] {
            if c == 'h' {
                fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
                return
            }
            break  // -h only has priority when it's the FIRST char in that arg
        }
    }
}
```

Wait — rereading the spec: `-zh` → prints the bitmask (h is in a flag with z before it, so h doesn't take priority). `-z -h` → prints options (h is first in its own argument). The rule is: if `-h` is encountered as the **first character** of any argument (i.e., `arg == "-h..."` and `h` is the first option letter), check priority based on argument order.

Simpler reading of the examples:
- `-abc -hijk` → `h` is at start of second arg → print options (h takes over its arg)
- `-zh` → `z` is first, `h` is later in same arg → h doesn't abort
- `-z -h` → `h` is first in its own arg → print options

The key rule: scan args left to right; if any arg starts with `-h` (after the dash), print options.

```go
for _, arg := range os.Args[1:] {
    if len(arg) >= 2 && arg[0] == '-' && arg[1] == 'h' {
        fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
        return
    }
}
```

**Step 2 — Parse flags and set bits:**

```go
for _, arg := range os.Args[1:] {
    if len(arg) < 2 || arg[0] != '-' {
        fmt.Println("Invalid Option")
        return
    }
    for _, c := range arg[1:] {
        if c < 'a' || c > 'z' {
            fmt.Println("Invalid Option")
            return
        }
        bit := uint(c - 'a')   // 'a'→0, 'b'→1, ..., 'z'→25
        flags |= int32(1) << bit
    }
}
```

**Step 3 — Print as 4 groups of 8 bits, MSB first:**

```go
fmt.Printf("%08b %08b %08b %08b\n",
    (flags>>24)&0xFF,
    (flags>>16)&0xFF,
    (flags>>8)&0xFF,
    flags&0xFF)
```

**Trace `-abc`:**
- `a` = bit 0 → flags = `0b000...001`
- `b` = bit 1 → flags = `0b000...011`
- `c` = bit 2 → flags = `0b000...111`
- Output: `00000000 00000000 00000000 00000111` ✓

**Trace `-z`:**
- `z` = bit 25 → flags = `0b00000010_00000000_00000000_00000000`
- Output: `00000010 00000000 00000000 00000000` ✓

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `flags |= 1 << bit` with plain `int` on a 64-bit system | Upper 4 bytes may be non-zero | Use `int32` or `uint32` and mask with `0xFF` when printing |
| Not checking `-h` priority before processing | `-h` in a non-first arg incorrectly triggers help | Scan for `-h` as first option letter in any arg, in arg order |
| Checking `c < 'a'` but not `c > 'z'` | Digits and symbols pass validation | Both bounds required |

## Solving This Challenge

### Algorithm
1. If no args, print options list and return.
2. Scan args in order: if any arg starts with `-h` (second char), print options and return.
3. Parse each arg: validate starts with `-`, each char is `a`–`z`, set bit `c - 'a'`.
4. Print flags as 4 groups of 8 bits.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [70-piglatin](../70-piglatin/README.md)
