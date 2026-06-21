# Prerequisites for 25-countrepeats

## Before You Start

To solve this challenge you need to understand:

### 1. The `for...range` Loop with State

From [13-stringlength skills.md](../13-stringlength/skills.md): iterating over characters. Here you also need to **remember the previous character** across iterations, which means declaring a variable before the loop:

```go
var prev rune      // declared outside the loop — persists across iterations
for _, c := range s {
    // use prev here
    prev = c       // update at the end of each iteration
}
```

### 2. Boolean Flags

A boolean variable used to remember whether you are in a certain state:

```go
inRepeat := false   // start: not in a repeat

// Inside the loop:
if someCondition {
    inRepeat = true   // enter the state
} else {
    inRepeat = false  // exit the state
}
```

You used a similar idea in [16-isempty](../16-isempty/skills.md) (returning on first match), but here the flag persists across iterations instead of causing an immediate return.

### 3. The Accumulator Pattern

From [13-stringlength skills.md](../13-stringlength/skills.md): `count := 0` then `count++` inside the loop. Here you only increment under a compound condition.

### 4. `var` Declaration for Zero Value

`var prevChar rune` declares a `rune` initialised to its zero value (`0`). This is safe as a "nothing seen yet" sentinel because `0` is not a printable character:

```go
var prevChar rune   // = 0
```

Compare to `:=` which requires an explicit initial value.

## Review If Stuck

- [13-stringlength skills.md](../13-stringlength/skills.md) — `for...range` and counting
- [16-isempty skills.md](../16-isempty/skills.md) — boolean logic
- [20-checknumber skills.md](../20-checknumber/skills.md) — why having a state variable changes behaviour

## You're Ready When You Can...

- [ ] Declare a variable before a loop and read/write it inside the loop
- [ ] Track a boolean flag that changes state based on the current character vs the previous one
- [ ] Increment a count only when a specific compound condition is true
- [ ] Explain what a "sentinel value" is and why `var prevChar rune` gives a safe one

## Next Steps

- [26-retainfirsthalf](../26-retainfirsthalf/README.md)
- [26-retainfirsthalf README](../26-retainfirsthalf/README.md) — next challenge
