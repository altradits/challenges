# 179-binary-search — Binary Search

## Task

Implement in `package piscine`:

```go
// BinarySearch returns the index of target in sorted slice nums,
// or -1 if not found. Slice must be sorted in ascending order.
func BinarySearch(nums []int, target int) int

// SearchInsertPosition returns the index where target should be inserted
// to keep nums sorted. If target exists, returns its index.
func SearchInsertPosition(nums []int, target int) int

// FindFirstOccurrence returns the index of the FIRST occurrence of target,
// or -1 if not found.
func FindFirstOccurrence(nums []int, target int) int

// FindLastOccurrence returns the index of the LAST occurrence of target,
// or -1 if not found.
func FindLastOccurrence(nums []int, target int) int
```

## Examples

```
BinarySearch([]int{1,3,5,7,9}, 5)   →  2
BinarySearch([]int{1,3,5,7,9}, 6)   →  -1
BinarySearch([]int{}, 5)             →  -1

SearchInsertPosition([]int{1,3,5,6}, 5)  →  2
SearchInsertPosition([]int{1,3,5,6}, 2)  →  1
SearchInsertPosition([]int{1,3,5,6}, 7)  →  4

FindFirstOccurrence([]int{1,2,2,2,3}, 2)  →  1
FindLastOccurrence([]int{1,2,2,2,3}, 2)   →  3
```

Read `skills.md` before you start.

---

## Time limit

**30 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 179
```

or run the tests directly:

```bash
go test ./179-binary-search/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 179` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
