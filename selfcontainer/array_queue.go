package selfcontainer

import "sync"

type ArrayQueue struct {
	Values []int
	Size   int
	Lock   sync.RWMutex
}

func (a *ArrayQueue) Add(value int) {
	a.Lock.Lock()
	defer a.Lock.Unlock()

	a.Values = append(a.Values, value)
	a.Size++
}

func (a *ArrayQueue) Remove() int {
	a.Lock.Lock()
	defer a.Lock.Unlock()

	if a.Size == 0 {
		panic("empty")
	}

	head := a.Values[0]

	//a.Values = a.Values[1:a.Size] // 1.但原数据内存不能回收

	newValue := make([]int, a.Size-1) // 2.新建一个新数组，将原Value赋值回去
	for i := 0; i < a.Size-1; i++ {
		newValue[i] = a.Values[i+1]
	}
	a.Values = newValue
	a.Size--

	return head
}

func (a *ArrayQueue) Head() int {
	a.Lock.RLock()
	defer a.Lock.RUnlock()

	if a.Size == 0 {
		panic("empty")
	}

	return a.Values[0]
}

func (a *ArrayQueue) Len() int {
	a.Lock.RLock()
	defer a.Lock.RUnlock()

	return a.Size
}

func (a *ArrayQueue) Empty() bool {
	a.Lock.RLock()
	defer a.Lock.RUnlock()

	return a.Size == 0
}
