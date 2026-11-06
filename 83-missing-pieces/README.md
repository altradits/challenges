# 83-missing-pieces

## What This Directory Is

These are the "missing pieces" — foundational concepts that were skipped over too quickly in the early challenges but that every Go developer needs to understand deeply.

The three sub-challenges here fill in those gaps before you move forward into more complex string manipulation.

## Sub-Challenges

### [hello/](./hello/)

A "Hello World" program with a twist. It prints a message that previews what comes next: sliding windows. Before you can understand algorithms, you need to understand that strings are sequences. This is where that idea is planted.

### [printlastchar/](./printlastchar/)

How do you print the last character of a string without knowing its length in advance? This challenge forces you to use `len()` and understand the gap between human counting (1, 2, 3...) and computer indexing (0, 1, 2...).

### [strlenindex/](./strlenindex/)

A deeper look at how `len()` and string indexing relate. `len()` counts from 1. Indexes start at 0. This challenge makes that relationship concrete through experimentation with `%c`, byte values, and string conversion.

## Why These Matter

These concepts are the building blocks for every string algorithm that follows. Sliding windows, substring search, character replacement — all of them depend on understanding:

- How strings are laid out in memory as sequences of bytes
- How to navigate those sequences using indexes
- The relationship between length (count) and the last valid index

## Order to Work Through

1. `hello/` — start here, it sets the frame
2. `strlenindex/` — understand len() and index before going further
3. `printlastchar/` — apply the index concept to get the last character

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

This is a manual challenge — there are no hidden tests. Run:

```bash
go run ./cmd/drill check 83
```

`drill check` walks a self-attestation checklist and records an honest
pass. Be honest; you are only cheating your own memory.

## Repeat schedule

No 1/3/7/21-day repeats — this is reference material or terminal
practice. Revisit whenever `prerequisites.md` of a later challenge points
back here.
