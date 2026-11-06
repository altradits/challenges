# Skills for 173-crypto-sha256

## What You'll Learn

**Previous:** [172-encoding-hex](../172-encoding-hex/skills.md) | **Next:** [174-worker-pool](../174-worker-pool/skills.md)

**Challenge:** Hash bytes with SHA-256, Bitcoin's workhorse hash.

## SHA-256 in Go

```go
import "crypto/sha256"

sum := sha256.Sum256([]byte("abc")) // [32]byte, NOT a slice
fmt.Printf("%x\n", sum)              // ba7816bf8f01cfea...
```

`Sum256` returns a `[32]byte` array. That matters: arrays are values,
so `sum1 == sum2` compares digests directly — no `bytes.Equal` needed.

## Double SHA-256

Bitcoin hashes almost everything twice: `SHA256d(x) = SHA256(SHA256(x))`.
Block hashes, txids, and Merkle nodes are all SHA-256d. The P2P message
checksum is the first 4 bytes of the SHA-256d of the payload.

```go
first := sha256.Sum256(payload)
second := sha256.Sum256(first[:]) // array → slice with [:]
```

## Streaming with hash.Hash

For data too big to hold in memory, hash incrementally:

```go
h := sha256.New()          // hash.Hash (an io.Writer)
if _, err := io.Copy(h, r); err != nil {
    return err
}
sum := h.Sum(nil) // appends digest to nil → fresh []byte
```

`Hash` embeds `io.Writer`, so `io.Copy` streams the reader straight in.
`Sum(nil)` returns the 32-byte digest as a slice. Never `io.ReadAll`
first — that defeats the purpose.

## Fixed Test Vectors (Memorize These)

```
SHA256("")    = e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
SHA256("abc") = ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad
SHA256d("")   = 5df6e0e2761359d30a8275058e299fcc0381534545f55cf43e41983f5d4c9456
```

If your code reproduces these three, your hashing is right.

## Solving the Challenge

- `SHA256Hex`: `sha256.Sum256`, format with `%x` (or `hex.EncodeToString`).
- `SHA256d`: hash twice; remember `first[:]`.
- `HashReader`: `sha256.New` + `io.Copy` + copy `Sum(nil)` into `[32]byte`.
- `Checksum4`: `SHA256d`, take bytes `[0:4]`.
