# Skills for 186-compactsize-varint

## What You'll Learn

**Previous:** [185-sliding-window](../185-sliding-window/skills.md) | **Next:** [187-sha256d-txid](../187-sha256d-txid/skills.md)

**Challenge:** Bitcoin's variable-length integer, the first consensus
encoding every contributor meets.

## The CompactSize Format

| Value range | Encoded as |
|---|---|
| 0 – 252 | 1 byte: the value itself |
| 253 – 0xffff | `0xfd` + 2 bytes LE |
| 0x10000 – 0xffffffff | `0xfe` + 4 bytes LE |
| larger | `0xff` + 8 bytes LE |

So 252 is `[fc]`, 253 is `[fd fd 00]`, 65535 is `[fd ff ff]`, 65536 is
`[fe 00 00 01 00]`. Multi-byte payloads are little-endian.

## Canonical (Minimal) Encoding

Consensus rules demand the *shortest* form: `[fd 00 00]` (value 0 in
3 bytes) is invalid — 0 must be `[00]`. A decoder must reject:

- `0xfd` prefix with value < 253
- `0xfe` prefix with value < 0x10000
- `0xff` prefix with value < 0x100000000

Non-canonical encodings are a classic attack vector (transaction
malleability relied on similar laxness), so btcd rejects them — and so
will you.

## Not binary.Uvarint

Go's `encoding/binary` varint uses 7-bits-per-byte with continuation
bits — protobuf style, unrelated to CompactSize. Do not use `PutUvarint`
here; branch on the value and use fixed-width `binary.LittleEndian.Put*`
for the payload, or plain shifts.

## Decoding Safely

`DecodeVarint(b)` returns `(value, bytesConsumed, err)`. Check the
length *before* every read:

```go
if len(b) == 0 {
    return 0, 0, errors.New("empty input")
}
switch b[0] {
case 0xfd:
    if len(b) < 3 { return 0, 0, errShort }
    v := binary.LittleEndian.Uint16(b[1:3])
    if v < 253 { return 0, 0, errNonCanonical }
    return uint64(v), 3, nil
...
}
```

Returning the consumed count lets callers chain decodes over a buffer —
exactly how block parsers walk a raw message.

## Solving the Challenge

- `EncodeVarint`: four branches on `n`, each appending prefix + LE bytes.
- `DecodeVarint`: switch on `b[0]`; values `< 0xfd` return `(b[0], 1, nil)`.
