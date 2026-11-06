# Skills for 171-encoding-binary

## What You'll Learn

**Previous:** [170-bytes-bufio](../170-bytes-bufio/skills.md) | **Next:** [172-encoding-hex](../172-encoding-hex/skills.md)

**Challenge:** Encode and decode integers in both byte orders.

## Endianness in One Minute

`0x01020304` stored little-endian (LE) is bytes `04 03 02 01` — least
significant byte first. Big-endian (BE) is `01 02 03 04`, the way humans
write numbers. Bitcoin consensus data (txids, amounts, block headers) is
LE on the wire; hashes are *displayed* BE (reversed). Mixing them up is a
classic consensus bug — get it right here.

## encoding/binary

```go
import "encoding/binary"

b := make([]byte, 4)
binary.LittleEndian.PutUint32(b, 0x01020304) // b = [04 03 02 01]
v := binary.LittleEndian.Uint32(b)           // 0x01020304

b8 := make([]byte, 8)
binary.BigEndian.PutUint64(b8, 1) // [00 00 00 00 00 00 00 01]
```

Available widths: `Uint16/32/64` and matching `Put` variants, on both
`binary.LittleEndian` and `binary.BigEndian`. There is also
`binary.Write(w, order, value)` / `binary.Read(r, order, &value)` for
streaming to an `io.Writer`/`io.Reader`.

Panic rule: `Uint32` panics on short slices, so guard lengths or pad.
`binary.Size(v)` reports the encoded size of a fixed-size value.

## Packing a Record

A 6-byte record — `uint32` id + `uint16` score, both LE:

```go
b := make([]byte, 6)
binary.LittleEndian.PutUint32(b[0:4], id)
binary.LittleEndian.PutUint16(b[4:6], score)
```

Unpacking is the mirror image. Always validate the length first and
return an error — never panic on foreign bytes.

## Solving the Challenge

- `Encode*`: allocate the exact size, `Put*`, return.
- `Decode*`: pad short input with zeros (e.g. copy into a zeroed
  4/8-byte array) so short input reads as zero instead of panicking.
- `PackRecord`/`UnpackRecord`: the 6-byte layout above; `UnpackRecord`
  errors unless `len(b) == 6`.
