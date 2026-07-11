# 177-stack — Stack (LIFO)

## Challenge

Implement a generic integer stack in `package piscine`:

```go
type Stack struct { /* unexported fields */ }

func NewStack() *Stack
func (s *Stack) Push(val int)
func (s *Stack) Pop() (int, bool)   // returns false if empty
func (s *Stack) Peek() (int, bool)  // look without removing
func (s *Stack) IsEmpty() bool
func (s *Stack) Size() int

// IsBalanced checks if the brackets in s are correctly balanced.
// Brackets: () [] {}
// Example: "({[]})" → true, "([)]" → false, "(((" → false
func IsBalanced(s string) bool
```

## Examples

```
st := NewStack()
st.Push(1)
st.Push(2)
st.Push(3)
st.Peek()    →  3, true
st.Pop()     →  3, true
st.Pop()     →  2, true
st.Size()    →  1

IsBalanced("({[]})")  →  true
IsBalanced("([)]")    →  false
IsBalanced("hello")   →  true  (no brackets = balanced)
IsBalanced("((")      →  false
```

Read `skills.md` before you start.
