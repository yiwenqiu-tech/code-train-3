package train

import "container/list"

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	l := list.New()
	return &Stack{
		list: l,
	}
}

func (s *Stack) Push(value any) {
	s.list.PushBack(value)
}

func (s *Stack) Pop() any {
	backValue := s.list.Back()
	s.list.Remove(backValue)
	return backValue.Value
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func isValid(s string) bool {
	stack := NewStack()
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack.Push(c)
		} else {
			if stack.Len() == 0 {
				return false
			}
			value := stack.Pop().(int32)
			if !isMatch(byte(c), byte(value)) {
				return false
			}
		}
	}
	if stack.Len() != 0 {
		return false
	}
	return true
}

func isMatch(c byte, v byte) bool {
	if (c == ')' && v == '(') || (c == ']' && v == '[') || (c == '}' && v == '{') {
		return true
	}
	return false
}
