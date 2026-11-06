# 190-p2p-message-codec — Bitcoin P2P Message Framing

## Task

In `package piscine`, frame and parse Bitcoin P2P messages — the envelope
every `version`, `inv`, `block`, and `tx` travels in:

```go
func EncodeMessage(magic [4]byte, command string, payload []byte) ([]byte, error)
func DecodeMessage(r io.Reader, magic [4]byte) (string, []byte, error)
```

Read `skills.md` before you start.

## Input / Output

- `EncodeMessage` builds `magic || command(12, padded) || len(4 LE) ||
  checksum(4) || payload`. Errors on empty/overlong command or oversize payload.
- `DecodeMessage` reads one message from `r`, verifying magic, length cap,
  and checksum. Returns the command and payload.

## Examples

```
msg, _ := EncodeMessage([4]byte{0xf9,0xbe,0xb4,0xd9}, "ping", []byte{1,2,3,4,5,6,7,8})
cmd, payload, _ := DecodeMessage(bytes.NewReader(msg), magic)
cmd      →  "ping"
payload  →  [1 2 3 4 5 6 7 8]
```

## Constraints

- Command: 1–12 ASCII bytes, null-padded to 12. Anything else errors.
- Payload cap: 32 MiB (`32 << 20`); check *before* allocating.
- Decode must use `io.ReadFull`, verify magic first, checksum last.

## Time limit

**35 minutes** (round 1). Beat it on every repeat: each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

    go run ./cmd/drill check 190

or run the tests directly:

    go test ./190-p2p-message-codec/

Green means done. Red means read the failure message, fix from memory, check again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, beat your previous time. If you cannot, re-read `skills.md` once,
then close it and go again.
