# Skills for 187-sha256d-txid

## What You'll Learn

**Previous:** [186-compactsize-varint](../186-compactsize-varint/skills.md) | **Next:** [188-merkle-root](../188-merkle-root/skills.md)

**Challenge:** Hash like Bitcoin and display hashes like a block explorer.

## Internal vs Display Byte Order

Inside blocks and transactions, hashes are little-endian bytes. Humans
read them big-endian — reversed. The genesis block hash is stored as
`6fe28c0a…00000000` but every explorer shows
`000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f`.

Rule of thumb: **hash in internal order, reverse only to display**.
Bugs come from reversing twice, or hashing the display string instead
of the raw bytes. btcd keeps `chainhash.Hash` in internal order and
reverses in `String()` — mirror that here.

## DoubleSHA

```go
import "crypto/sha256"

func DoubleSHA(b []byte) [32]byte {
    first := sha256.Sum256(b)
    return sha256.Sum256(first[:])
}
```

`Sum256` takes a slice and returns an array, hence `first[:]`.

## Display and Parse

```go
func TxidDisplay(hash [32]byte) string {
    rev := hash // arrays copy on assignment — safe to mutate
    for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
        rev[i], rev[j] = rev[j], rev[i]
    }
    return hex.EncodeToString(rev[:])
}

func TxidParse(s string) ([32]byte, error) {
    var out [32]byte
    if len(s) != 64 {
        return out, fmt.Errorf("txid must be 64 hex chars, got %d", len(s))
    }
    raw, err := hex.DecodeString(s)
    if err != nil {
        return out, err
    }
    copy(out[:], raw)
    for i, j := 0, 31; i < j; i, j = i+1, j-1 {
        out[i], out[j] = out[j], out[i]
    }
    return out, nil
}
```

Note `rev := hash` copies the whole array — a `[32]byte` assignment is a
value copy, so reversing `rev` leaves the caller's hash untouched.

## Fixed Vector (Memorize It)

```
DoubleSHA("hello") internal = 9595c9df90075148eb06860365df33584b75bff782a510c6cd4883a419833d50
                   display  = 503d8319a48348cdc610a582f7bf754b5833df65038606eb48510790dfc99595
```

## Solving the Challenge

Exactly the three functions above. Validate length before decoding in
`TxidParse`, and always reverse bytes — never hex characters.
