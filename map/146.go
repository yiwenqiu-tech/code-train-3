package _map

import "container/list"

type Item struct {
	Key   int
	Value int
}

type LRUCache struct {
	// 注意：这里需要存Item，不能只存Value，只存Value的话，当容量不足时，从队列移除数据时无法根据value找到Key，从而无法从map中移除数据
	Deque   *list.List            // 双端队列，Value存Item,
	ElemMap map[int]*list.Element // key为key值，Value为双端队列里的元素
	Cap     int
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{}
	lruCache.ElemMap = make(map[int]*list.Element)
	lruCache.Deque = list.New()
	lruCache.Cap = capacity
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if v, exist := this.ElemMap[key]; exist {
		this.Deque.MoveToFront(this.ElemMap[key])
		return v.Value.(*Item).Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, exist := this.ElemMap[key]; exist { // 本来就存在了，更新值，并且调整deque位置
		elem := this.ElemMap[key]
		elem.Value = &Item{
			Key:   key,
			Value: value,
		}
		this.Deque.MoveToFront(elem)
	} else {
		if this.Deque.Len()+1 > this.Cap { // 当加入当前值后，大于容量时，则需先腾出最后一个元素
			backElem := this.Deque.Back()
			delete(this.ElemMap, backElem.Value.(*Item).Key)
			this.Deque.Remove(backElem)
		}
		this.Deque.PushFront(&Item{
			Key:   key,
			Value: value,
		})
		this.ElemMap[key] = this.Deque.Front()
	}
}
