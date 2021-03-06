package stack

// FILO Stack
type Stack struct {
	stack []string
}

func (s Stack) GetStack() []string {
	return s.stack
}

func (s *Stack) Push(value string) {
	s.stack = append(s.stack, value)
}

func (s *Stack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *Stack) Init() {
	s.stack = []string{}
}
