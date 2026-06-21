# Prerequisites for 146-encoding-json

## Before You Start

- [13-structs](../13-structs/skills.md) — struct field tags use the same syntax
- [16-interfaces](../16-interfaces/skills.md) — json functions accept `any` (empty interface)
- [19-fileio](../19-fileio/skills.md) — io.Reader/Writer used by Encoder/Decoder
- [145-time](../145-time/skills.md) — time.Time appears in JSON frequently

## You're Ready When You Can...

- [ ] Use `json.Marshal` and `json.Unmarshal`
- [ ] Add `json:"field_name"` struct tags to control serialization
- [ ] Use `omitempty` and `json:"-"` to hide fields
- [ ] Encode to and decode from an `io.Writer`/`io.Reader` with Encoder/Decoder

## Next Steps

- [147-http-basics](../147-http-basics/README.md)
