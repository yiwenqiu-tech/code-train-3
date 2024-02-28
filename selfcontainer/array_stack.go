package selfcontainer

import "sync"

type ArrayStack struct {
	Values []int
	Size   int
	Lock   sync.RWMutex
}

func (a *ArrayStack) Push(value int) {
	a.Lock.Lock()
	defer a.Lock.Unlock()

	a.Values = append(a.Values, value)
	a.Size++
}

func (a *ArrayStack) Pop() int {
	a.Lock.Lock()
	defer a.Lock.Unlock()
	if a.Size == 0 {
		panic("empty")
	}
	top := a.Values[a.Size-1] // pop最后push的元素

	//a.Values = a.Values[0 : a.Size-1] // 1.收缩Values，注意这里并没有真正的释放内存

	var newValues = make([]int, a.Size-1) // 2.方式2，新建一块values赋值
	copy(newValues, a.Values)
	a.Values = newValues

	a.Size--
	return top
}

func (a *ArrayStack) Peek() int {
	a.Lock.RLock()
	defer a.Lock.RUnlock()
	if a.Size == 0 {
		panic("empty")
	}
	return a.Values[0]
}

func (a *ArrayStack) Len() int {
	a.Lock.RLock()
	defer a.Lock.RUnlock()
	return a.Size
}

func (a *ArrayStack) Empty() bool {
	a.Lock.RLock()
	defer a.Lock.RUnlock()
	return a.Size == 0
}
