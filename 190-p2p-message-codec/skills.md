# Skills for 190-p2p-message-codec

## What You'll Learn

**Previous:** [189-block-header](../189-block-header/skills.md) | **Next:** [191-p2pkh-script](../191-p2pkh-script/skills.md)

**Challenge:** Speak the Bitcoin wire protocol's outer envelope.

## The 24-Byte Header

| Offset | Size | Field |
|---|---|---|
| 0 | 4 | Magic (network ID; mainnet `f9 be b4 d9`) |
| 4 | 12 | Command, ASCII, null-padded (`version\0\0\0\0\0`) |
| 16 | 4 | Payload length, LE |
| 20 | 4 | Checksum: first 4 bytes of SHA-256d(payload) |

Then `length` bytes of payload. Total on the wire: `24 + length`.

## Encoding

```go
const MaxPayload = 32 << 20 // 32 MiB, protocol limit

func EncodeMessage(magic [4]byte, command string, payload []byte) ([]byte, error) {
    if len(command) == 0 || len(command) > 12 {
        return nil, fmt.Errorf("command must be 1-12 chars, got %d", len(command))
    }
    if len(payload) > MaxPayload {
        return nil, fmt.Errorf("payload %d bytes exceeds cap", len(payload))
    }
    out := make([]byte, 0, 24+len(payload))
    out = append(out, magic[:]...)
    cmd := make([]byte, 12)
    copy(cmd, command) // null-pads automatically
    out = append(out, cmd...)
    var ln [4]byte
    binary.LittleEndian.PutUint32(ln[:], uint32(len(payload)))
    out = append(out, ln[:]...)
    sum := doubleSHA(payload)
    out = append(out, sum[0], sum[1], sum[2], sum[3])
    out = append(out, payload...)
    return out, nil
}
```

## Decoding Defensively

Network input is hostile. Decode in this order:

1. `io.ReadFull(r, header)` — 24 bytes or an error. (`Read` may return
   short counts; `ReadFull` loops until full or failed.)
2. Compare magic — wrong network, drop it.
3. Read length; reject `> MaxPayload` **before** `make` — otherwise a
   lying peer makes you allocate 4 GB.
4. `io.ReadFull(r, payload)` — truncated stream errors here.
5. Verify checksum — corrupted payload errors here.

Return the command with padding stripped (`string(bytes.TrimRight(cmd, "\x00"))`).
Also validate the command bytes are printable ASCII — belt and braces
against garbage.

## Solving the Challenge

Encode as above; decode in the five steps. Keep your own local
double-SHA helper — every challenge is self-contained.
