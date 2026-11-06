# Skills for 191-p2pkh-script

## What You'll Learn

**Previous:** [190-p2p-message-codec](../190-p2p-message-codec/skills.md) | **Next:** [192-tooling-practice](../192-tooling-practice/skills.md)

**Challenge:** Handle the script behind most Bitcoin payments.

## The P2PKH Template

A pay-to-pubkey-hash output locks coins to a 20-byte HASH160 of a
public key. The scriptPubKey is always 25 bytes:

```
76 a9 14 <20-byte pubkey hash> 88 ac
```

Byte by byte:

| Byte(s) | Meaning |
|---|---|
| `0x76` | OP_DUP — duplicate the top stack item |
| `0xa9` | OP_HASH160 — hash it (SHA-256 then RIPEMD-160) |
| `0x14` | Push the next 20 bytes |
| 20 bytes | The pubkey hash |
| `0x88` | OP_EQUALVERIFY — must match, else fail |
| `0xac` | OP_CHECKSIG — verify the signature |

To spend, the input provides `<signature> <pubkey>`; the script checks
the pubkey hashes to the committed value and the signature is valid.
LND watches for scripts like these when tracking on-chain funds.

## Building and Parsing

```go
func BuildP2PKH(pubkeyHash []byte) ([]byte, error) {
    if len(pubkeyHash) != 20 {
        return nil, fmt.Errorf("pubkey hash must be 20 bytes, got %d", len(pubkeyHash))
    }
    out := []byte{0x76, 0xa9, 0x14}
    out = append(out, pubkeyHash...)
    return append(out, 0x88, 0xac), nil
}

func ParseP2PKH(script []byte) ([]byte, error) {
    if !IsP2PKH(script) {
        return nil, fmt.Errorf("not a P2PKH script")
    }
    return bytes.Clone(script[3:23]), nil
}

func IsP2PKH(script []byte) bool {
    return len(script) == 25 &&
        script[0] == 0x76 && script[1] == 0xa9 && script[2] == 0x14 &&
        script[23] == 0x88 && script[24] == 0xac
}
```

`bytes.Clone` (Go 1.20+) returns a copy so callers cannot mutate your
input through the returned slice. Check the length *first* — indexing
before the length check panics on short input.

## Solving the Challenge

The three functions above. Single source of truth: `ParseP2PKH`
delegates to `IsP2PKH` so the template lives in exactly one place.
