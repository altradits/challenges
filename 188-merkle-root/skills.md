# Skills for 188-merkle-root

## What You'll Learn

**Previous:** [187-sha256d-txid](../187-sha256d-txid/skills.md) | **Next:** [189-block-header](../189-block-header/skills.md)

**Challenge:** Commit to a transaction list with one 32-byte root.

## The Merkle Tree

Pair up the txids; hash each pair with SHA-256d; repeat with the
results until one hash remains. That hash commits to every transaction:
change one byte anywhere and the root changes.

```
level 0:  a        b        c
               \  /        /
level 1:   H(a||b)    H(c||c)   ← lone leaf duplicated
                 \    /
level 2:       H(Hab||Hcc)      ← root
```

Three Bitcoin rules to get exactly right:

1. **Odd level duplicates the last hash.** With 3 leaves, level 1 is
   `[H(a||b), H(c||c)]` — `c` pairs with itself.
2. **Single txid is its own root.** A one-transaction block's Merkle
   root equals that txid. (Duplicating here too is the classic bug.)
3. **Empty list → zero hash.** Define it; callers of coinbase-only test
   code hit this.

All hashing happens in internal byte order. Display reversal is only
for humans, never for tree math.

## The Loop

```go
func MerkleRoot(txids [][32]byte) [32]byte {
    if len(txids) == 0 {
        return [32]byte{}
    }
    level := txids
    for len(level) > 1 {
        var next [][32]byte
        for i := 0; i < len(level); i += 2 {
            left := level[i]
            right := left
            if i+1 < len(level) {
                right = level[i+1]
            }
            var pair [64]byte
            copy(pair[:32], left[:])
            copy(pair[32:], right[:])
            next = append(next, sha256d(pair[:]))
        }
        level = next
    }
    return level[0]
}
```

Careful: `level := txids` aliases the caller's slice. This loop only
reads it (never writes), so aliasing is safe — but appending to `level`
would corrupt the caller. When in doubt, copy.

## Fixed Vectors

With `a/b/c = DoubleSHA("tx0"/"tx1"/"tx2")`:

```
root([a, b])    = f7ed2a56db7927143defd32e380fb083f74970e0d6f31c00b8aded86aefdf1d2
root([a, b, c]) = e723898188fe36f157a1a97e4a608823a4c64cc81ac658dba49c4b66eef59a1f
```

## Solving the Challenge

The loop above, plus your own local double-SHA helper (each challenge
is self-contained — do not import other challenge packages).
