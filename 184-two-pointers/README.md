# 184-two-pointers — Two Pointer Technique

## Challenge

Implement in `package piscine`:

```go
// TwoSum finds two indices i, j where nums[i]+nums[j]==target.
// nums is sorted. Returns (-1,-1) if no pair exists.
func TwoSum(nums []int, target int) (int, int)

// RemoveDuplicates removes duplicates from sorted nums in-place.
// Returns the new length (elements after it are irrelevant).
func RemoveDuplicates(nums []int) int

// ReverseInPlace reverses a slice in-place.
func ReverseInPlace(nums []int)

// ContainerWithMostWater: given heights, find two lines that form
// the container with the maximum water volume. Returns the volume.
func ContainerWithMostWater(heights []int) int
```

## Examples

```
TwoSum([]int{1,2,3,4,6}, 6)   →  1, 3  (indices of 2 and 4)
TwoSum([]int{1,3,5,7}, 10)    →  1, 3  (indices of 3 and 7)
TwoSum([]int{1,2,3}, 100)     →  -1, -1

RemoveDuplicates([]int{1,1,2,2,3})  →  3  (slice becomes [1,2,3,...])

ReverseInPlace([]int{1,2,3,4,5})  →  [5,4,3,2,1]

ContainerWithMostWater([]int{1,8,6,2,5,4,8,3,7})  →  49
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
go run ./cmd/drill check 184
```

or run the tests directly:

```bash
go test ./184-two-pointers/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 184` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
