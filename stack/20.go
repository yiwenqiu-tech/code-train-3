package stack

type CharListNode struct {
	Val  byte
	Next *CharListNode
}

type CharListStack struct {
	Root *CharListNode
	Size int
}

func (l *CharListStack) Push(value byte) {
	if l.Root == nil {
		l.Root = &CharListNode{
			Val:  value,
			Next: nil,
		}
	} else {
		preRoot := l.Root
		l.Root = &CharListNode{
			Val:  value,
			Next: preRoot,
		}
	}

	l.Size++
}

func (l *CharListStack) Pop() byte {
	if l.Size == 0 {
		panic("empty")
	}
	top := l.Root.Val
	l.Root = l.Root.Next
	l.Size--
	return top
}

func (l *CharListStack) Peek() byte {

	if l.Size == 0 {
		panic("empty")
	}
	return l.Root.Val
}

func (l *CharListStack) Len() int {
	return l.Size
}

func (l *CharListStack) Empty() bool {
	return l.Size == 0
}

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	var sStack CharListStack
	for index := range s {
		if s[index] == '(' || s[index] == '{' || s[index] == '[' {
			sStack.Push(s[index])
		} else if s[index] == ')' {
			if sStack.Empty() {
				return false
			}
			tmp := sStack.Pop()
			if tmp != '(' {
				return false
			}
		} else if s[index] == '}' {
			if sStack.Empty() {
				return false
			}
			tmp := sStack.Pop()
			if tmp != '{' {
				return false
			}
		} else if s[index] == ']' {
			if sStack.Empty() {
				return false
			}
			tmp := sStack.Pop()
			if tmp != '[' {
				return false
			}
		} else {
			return false
		}
	}
	if sStack.Empty() {
		return true
	}
	return false
}
