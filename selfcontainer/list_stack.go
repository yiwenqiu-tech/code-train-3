package selfcontainer

import (
	"code-train-3/list"
	"sync"
)

type ListStack struct {
	Root *list.ListNode
	Size int
	Lock sync.RWMutex
}

func (l *ListStack) Push(value any) {
	l.Lock.Lock()
	defer l.Lock.Unlock()

	if l.Root == nil {
		l.Root = &list.ListNode{
			Val:  value,
			Next: nil,
		}
	} else {
		preRoot := l.Root
		l.Root = &list.ListNode{
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
