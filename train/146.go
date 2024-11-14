package train

import "container/list"

type eleValue struct {
	key   int
	value int
}

type LRUCache struct {
	cap int
	l   *list.List
	m   map[int]*list.Element
}

func Constructor146(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   list.New(),
		m:   make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, exist := this.m[key]; exist {
		this.l.MoveToFront(elem) // 更新到前面
		return elem.Value.(*eleValue).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	newElem := &eleValue{
		key:   key,
		value: value,
	}
	if elem, exist := this.m[key]; exist {
		elem.Value = newElem
		this.l.MoveToFront(elem) // 更新到前面
	} else {
		length := len(this.m)
		if length >= this.cap {
			backElem := this.l.Back()
			backElemValue := backElem.Value.(*eleValue)
			this.l.Remove(backElem)
			delete(this.m, backElemValue.key)
		}
		this.l.PushFront(newElem)
		this.m[key] = this.l.Front()
	}
}
