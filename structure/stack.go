package structure

// Stack 栈
type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{items: []interface{}{}}
}

func (s *Stack) Pop() (interface{}, error) {
	slen := len(s.items)
	if slen == 0 {
		return 0, EmptyError
	}
	res := s.items[slen-1]
	s.items = s.items[:slen-1]
	return res, nil
}

func (s *Stack) Top() (interface{}, error) {
	slen := len(s.items)
	if slen == 0 {
		return 0, EmptyError
	}
	return s.items[slen-1], nil
}

func (s *Stack) Push(in interface{}) {
	s.items = append(s.items, in)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
