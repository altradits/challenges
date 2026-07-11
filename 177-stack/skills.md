# Skills for 177-stack

## What You'll Learn

**Previous:** [176-linked-list](../176-linked-list/skills.md) | **Next:** [178-queue](../178-queue/skills.md)

**Challenge:** Implement a LIFO stack and use it to solve bracket balancing.

## What Is a Stack?

A stack is a LIFO (Last In, First Out) collection. Think of a stack of plates — you always add and remove from the top.

```
Push(1) → [1]
Push(2) → [1, 2]
Push(3) → [1, 2, 3]
Pop()   → 3,  [1, 2]
Pop()   → 2,  [1]
```

## Implementation: Slice-Backed Stack

The simplest and fastest Go stack uses a slice:

```go
type Stack struct {
    data []int
}

func (s *Stack) Push(val int) {
    s.data = append(s.data, val)
}

func (s *Stack) Pop() (int, bool) {
    if len(s.data) == 0 {
        return 0, false
    }
    n := len(s.data) - 1
    val := s.data[n]
    s.data = s.data[:n]  // shrink slice without freeing memory
    return val, true
}

func (s *Stack) Peek() (int, bool) {
    if len(s.data) == 0 {
        return 0, false
    }
    return s.data[len(s.data)-1], true
}
```

The top of the stack is the **last element of the slice** — appending adds to the top, slicing off removes from the top. O(1) amortised for all operations.

## Bracket Balancing with a Stack

Classic stack application: scan left to right, push opening brackets, pop and check on closing brackets.

```go
func IsBalanced(s string) bool {
    pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
    stack := NewStack()  // stores runes as int (cast)

    for _, ch := range s {
        switch ch {
        case '(', '[', '{':
            stack.Push(int(ch))
        case ')', ']', '}':
            top, ok := stack.Pop()
            if !ok || rune(top) != pairs[ch] {
                return false  // unmatched closing
            }
        }
    }
    return stack.IsEmpty()  // unmatched openings left?
}
```

Visual trace for `"({[]})"`

```
ch='('  stack: ['(']
ch='{'  stack: ['(', '{']
ch='['  stack: ['(', '{', '[']
ch=']'  pop '[' == pairs[']'='['] ✓  stack: ['(', '{']
ch='}'  pop '{' == pairs['}'='{'] ✓  stack: ['(']
ch=')'  pop '(' == pairs[')'='('] ✓  stack: []
Empty → true
```

## Use Cases

Stacks appear in:
- **Undo/redo** systems (each action is pushed; undo pops)
- **Call stack** in program execution
- **Expression parsing** (RPN calculator — see challenge 142-rpncalc)
- **Depth-first search** (DFS)
- **Browser history** (back button = pop)

## Solving the Challenge

```go
package piscine

type Stack struct {
    data []int
}

func NewStack() *Stack { return &Stack{} }

func (s *Stack) Push(val int) {
    s.data = append(s.data, val)
}

func (s *Stack) Pop() (int, bool) {
    if len(s.data) == 0 { return 0, false }
    n := len(s.data) - 1
    val := s.data[n]
    s.data = s.data[:n]
    return val, true
}

func (s *Stack) Peek() (int, bool) {
    if len(s.data) == 0 { return 0, false }
    return s.data[len(s.data)-1], true
}

func (s *Stack) IsEmpty() bool { return len(s.data) == 0 }
func (s *Stack) Size() int     { return len(s.data) }

func IsBalanced(input string) bool {
    pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
    st := NewStack()
    for _, ch := range input {
        switch ch {
        case '(', '[', '{':
            st.Push(int(ch))
        case ')', ']', '}':
            top, ok := st.Pop()
            if !ok || rune(top) != pairs[ch] {
                return false
            }
        }
    }
    return st.IsEmpty()
}
```

## Documentation

- [container/list](https://pkg.go.dev/container/list) — standard doubly linked list (can be used as stack)
- [Go Spec — Slice expressions](https://go.dev/ref/spec#Slice_expressions)

**Next:** [178-queue](../178-queue/README.md)
