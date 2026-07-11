package piscine

type Stack struct {
	data []int
}

func NewStack() *Stack { return &Stack{} }

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	n := len(s.data) - 1
	val := s.data[n]
	s.data = s.data[:n]
	return val, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
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
