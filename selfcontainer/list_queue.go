package selfcontainer

import (
	"code-train-3/list"
	"sync"
)

type ListQueue struct {
	Root *list.ListNode
	Size int
	Lock sync.RWMutex
}

func (l *ListQueue) Add(value any) {
	l.Lock.Lock()
	defer l.Lock.Unlock()

	if l.Root == nil {
		l.Root = &list.ListNode{
			Val:  value,
			Next: nil,
		}
	} else { // O(n)
		newNode := &list.ListNode{
			Val:  value,
			Next: nil,
		}
		cur := l.Root
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = newNode
	}

	l.Size++
}

func (l *ListQueue) Remove() any {
	l.Lock.Lock()
	defer l.Lock.Unlock()

	if l.Size == 0 {
		panic("empty")
	}

	head := l.Root

	l.Root = l.Root.Next // 取队头，然后队头下一个节点成为队友
	l.Size--

	return head.Val
}

func (l *ListQueue) Head() any {
	l.Lock.RLock()
	defer l.Lock.RUnlock()

	if l.Size == 0 {
		panic("empty")
	}

	return l.Root.Val
}

func (l *ListQueue) Len() int {
	l.Lock.RLock()
	defer l.Lock.RUnlock()

	return l.Size
}

func (l *ListQueue) Empty() bool {
	l.Lock.RLock()
	defer l.Lock.RUnlock()

	return l.Size == 0
}
