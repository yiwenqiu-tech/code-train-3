package divide_conquer

import (
	"sync"
)

type AnyListNode struct {
	Val  any
	Next *AnyListNode
}

type ListStack struct {
	Root *AnyListNode
	Size int
	Lock sync.RWMutex
}

func (l *ListStack) Push(value any) {
	l.Lock.Lock()
	defer l.Lock.Unlock()

	if l.Root == nil {
		l.Root = &AnyListNode{
			Val:  value,
			Next: nil,
		}
	} else {
		preRoot := l.Root
		l.Root = &AnyListNode{
			Val:  value,
			Next: preRoot,
		}
	}

	l.Size++
}

func (l *ListStack) Pop() any {
	l.Lock.Lock()
	defer l.Lock.Unlock()

	if l.Size == 0 {
		panic("empty")
	}
	top := l.Root.Val
	l.Root = l.Root.Next
	l.Size--
	return top
}

func (l *ListStack) Peek() any {
	l.Lock.RLock()
	defer l.Lock.RUnlock()

	if l.Size == 0 {
		panic("empty")
	}
	return l.Root.Val
}

func (l *ListStack) Len() any {
	l.Lock.RLock()
	defer l.Lock.RUnlock()

	return l.Size
}

func (l *ListStack) Empty() bool {
	l.Lock.RLock()
	defer l.Lock.RUnlock()

	return l.Size == 0
}

var generateParenthesisRes []string
var generateParenthesisChosen []byte

func generateParenthesis(n int) []string {
	generateParenthesisRes = nil
	recurForGenerateParenthesis(n, n, n, 0)
	return generateParenthesisRes
}

func recurForGenerateParenthesis2(left int, right int, n int, index int) {
	// 剪枝，若前缀不满足直接返回，实测下来这种做法更慢，因为每个分支都需要建栈来分析
	if !isValidPrefix(string(generateParenthesisChosen)) {
		return
	}
	if index >= 2*n { // 判断字符串是否满足栈
		res := string(generateParenthesisChosen)
		if isValid(res) {
			generateParenthesisRes = append(generateParenthesisRes, res)
		}
		return
	}
	if left > 0 {
		generateParenthesisChosen = append(generateParenthesisChosen, '(')
		recurForGenerateParenthesis2(left-1, right, n, index+1)
		generateParenthesisChosen = generateParenthesisChosen[:len(generateParenthesisChosen)-1]
	}
	if right > 0 {
		generateParenthesisChosen = append(generateParenthesisChosen, ')')
		recurForGenerateParenthesis2(left, right-1, n, index+1)
		generateParenthesisChosen = generateParenthesisChosen[:len(generateParenthesisChosen)-1]
	}
}

func recurForGenerateParenthesis(left int, right int, n int, index int) {
	if index >= 2*n { // 通过枚举所有子集可能后，再判断字符串是否满足栈
		res := string(generateParenthesisChosen)
		if isValid(res) {
			generateParenthesisRes = append(generateParenthesisRes, res)
		}
		return
	}
	if left > 0 {
		generateParenthesisChosen = append(generateParenthesisChosen, '(')
		recurForGenerateParenthesis(left-1, right, n, index+1)
		generateParenthesisChosen = generateParenthesisChosen[:len(generateParenthesisChosen)-1]
	}
	if right > 0 {
		generateParenthesisChosen = append(generateParenthesisChosen, ')')
		recurForGenerateParenthesis(left, right-1, n, index+1)
		generateParenthesisChosen = generateParenthesisChosen[:len(generateParenthesisChosen)-1]
	}
}

func isValidPrefix(chosen string) bool {
	stack := ListStack{}
	for _, c := range chosen {
		if c == ')' {
			if stack.Size == 0 {
				return false
			}
			stack.Pop()
		} else {
			stack.Push(c)
		}
	}
	return true
}

func isValid(chosen string) bool {
	stack := ListStack{}
	for _, c := range chosen {
		if c == ')' {
			if stack.Size == 0 {
				return false
			}
			stack.Pop()
		} else {
			stack.Push(c)
		}
	}
	if stack.Size == 0 {
		return true
	}
	return false
}

// TODO 1.分治做法？？

// TODO 2.参考别人的剪枝做法？？
