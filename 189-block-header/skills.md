# Skills for 189-block-header

## What You'll Learn

**Previous:** [188-merkle-root](../188-merkle-root/skills.md) | **Next:** [190-p2p-message-codec](../190-p2p-message-codec/skills.md)

**Challenge:** Decode the 80 bytes every full node and every miner
passes around.

## The 80-Byte Layout

| Offset | Size | Field |
|---|---|---|
| 0 | 4 | Version (`int32`, LE) |
| 4 | 32 | PrevBlock hash (internal order) |
| 36 | 32 | MerkleRoot (internal order) |
| 68 | 4 | Timestamp (`uint32` Unix time, LE) |
| 72 | 4 | Bits (compact difficulty target, LE) |
| 76 | 4 | Nonce (`uint32`, LE) |

Everything is little-endian; the two hashes are raw internal-order
bytes — no reversal on parse or serialize.

## Parsing

```go
func ParseHeader(b []byte) (Header, error) {
    var h Header
    if len(b) != 80 {
        return h, fmt.Errorf("header must be 80 bytes, got %d", len(b))
    }
    h.Version = int32(binary.LittleEndian.Uint32(b[0:4]))
    copy(h.PrevBlock[:], b[4:36])
    copy(h.MerkleRoot[:], b[36:68])
    h.Timestamp = binary.LittleEndian.Uint32(b[68:72])
    h.Bits = binary.LittleEndian.Uint32(b[72:76])
    h.Nonce = binary.LittleEndian.Uint32(b[76:80])
    return h, nil
}
```

`copy(dst[:], src)` moves bytes into the fixed arrays. Serializing is
the mirror: allocate 80 bytes, `PutUint32` the integers, `copy` the
hashes.

## The Block Hash

`HeaderHash = DoubleSHA(SerializeHeader(h))`, internal order. The famous
leading zeros live at the *end* of the internal bytes and at the *start*
of the display string — same bytes, reversed view:

```
internal: 6fe28c0a...00000000
display:  000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f
```

## The Genesis Test Vector

```
01000000 || 00…00 (32) || 3ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a
|| 29ab5f49 || ffff001d || 1dac2b7c
```

Version 1, prev zero, timestamp 1231006505 (`0x495fab29`), bits
`0x1d00ffff`, nonce 2083236893 (`0x7c2bac1d`). If your parser reads
these fields and your hash matches, you parse headers correctly.

## Solving the Challenge

The parse function above, its mirror for serialize, and a local
double-SHA helper for the hash. Keep a single source of truth for the
offsets — magic numbers in two places drift apart.
