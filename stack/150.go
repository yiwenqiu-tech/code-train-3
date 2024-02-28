package stack

import "strconv"

type IntListNode struct {
	Val  int
	Next *IntListNode
}

type IntListStack struct {
	Root *IntListNode
	Size int
}

func (l *IntListStack) Push(value int) {
	if l.Root == nil {
		l.Root = &IntListNode{
			Val:  value,
			Next: nil,
		}
	} else {
		preRoot := l.Root
		l.Root = &IntListNode{
			Val:  value,
			Next: preRoot,
		}
	}

	l.Size++
}

func (l *IntListStack) Pop() int {
	if l.Size == 0 {
		panic("empty")
	}
	top := l.Root.Val
	l.Root = l.Root.Next
	l.Size--
	return top
}

func (l *IntListStack) Peek() int {
	if l.Size == 0 {
		panic("empty")
	}
	return l.Root.Val
}

func (l *IntListStack) Len() int {
	return l.Size
}

func (l *IntListStack) Empty() bool {
	return l.Size == 0
}

func evalRPN(tokens []string) int {
	var numStack = new(IntListStack)
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			numStack.Push(num) // 遇到数字就Push进去
		} else {
			num2 := numStack.Pop()
			num1 := numStack.Pop()
			switch token {
			case "+":
				numStack.Push(num1 + num2)
			case "-":
				numStack.Push(num1 - num2)
			case "*":
				numStack.Push(num1 * num2)
			case "/":
				numStack.Push(num1 / num2)
			}
		}
	}
	return numStack.Peek()
}
