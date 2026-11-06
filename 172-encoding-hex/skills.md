# Skills for 172-encoding-hex

## What You'll Learn

**Previous:** [171-encoding-binary](../171-encoding-binary/skills.md) | **Next:** [173-crypto-sha256](../173-crypto-sha256/skills.md)

**Challenge:** Hex-encode and decode bytes the Bitcoin way.

## Why Hex?

Hashes, public keys, and raw transactions are bytes, but humans read hex.
A txid like `a1b2…` is the *reversed* byte order of the internal hash —
Bitcoin displays hashes big-endian while storing them little-endian.
`ReverseHex` here is exactly that display conversion.

## encoding/hex

```go
import "encoding/hex"

s := hex.EncodeToString([]byte{0xde, 0xad}) // "dead"
b, err := hex.DecodeString("dead")          // []byte{0xde, 0xad}, nil
b, err = hex.DecodeString("zz")             // nil, encoding/hex: invalid byte
b, err = hex.DecodeString("abc")            // nil, odd length hex string
```

`DecodeString` accepts upper- and lowercase digits but rejects odd
lengths (`hex.ErrLength`) and bad digits (`hex.InvalidByteError`).
Streaming variants exist too: `hex.NewEncoder(w)` wraps a writer,
`hex.NewDecoder(r)` wraps a reader, and `hex.Dump(b)` prints the classic
offset/hex/ASCII table for debugging.

A handy check: `hex.DecodedLen(len(s))` gives the decoded size.

## Reversing Byte Order

Decode first, reverse the *bytes* (not the characters!), re-encode:

```go
b, err := hex.DecodeString(s)
for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
    b[i], b[j] = b[j], b[i]
}
return hex.EncodeToString(b), err
```

Reversing the string `"0102ff"` directly would give `"ff2010"` —
wrong. Always think in bytes.

## Solving the Challenge

- `ToHex`: one call to `hex.EncodeToString`.
- `FromHex`: one call to `hex.DecodeString`; return its error as-is.
- `IsHex`: decode and report `err == nil`.
- `ReverseHex`: decode, reverse bytes in place, encode.
