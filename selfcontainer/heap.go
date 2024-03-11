package selfcontainer

import "container/heap"

type Heap struct { // 堆
	heap.Interface
}

func (h *Heap) Init() {
	n := h.Len() - 1
	for i := n / 2; i >= 0; i-- { // 非叶子节点做down
		h.heapifyDown(i, n)
	}
}

func (h *Heap) Empty() bool {
	return h.Len() == 0
}

func (h *Heap) PopTop() any {
	n := h.Len() - 1
	h.Swap(0, n)
	h.heapifyDown(0, n-1) // 由于最后一个元素不要了，所以down的时候不参与
	return h.Pop()
}

func (h *Heap) Insert(elem any) {
	h.Push(elem)
	h.heapifyUp(h.Len() - 1) // 将最后一个位置做up
}

func (h *Heap) heapifyUp(n int) {
	p := n
	for {
		parent := (p - 1) / 2
		if parent == p {
			break
		}
		if h.Less(p, parent) { // 如果parent < p，则需要交换，否则终止
			h.Swap(parent, p)
			p = parent
		} else {
			break
		}
	}
}

func (h *Heap) heapifyDown(start, end int) {
	p := start
	for {
		child := p*2 + 1 // 左孩子
		if child > end { // 坐标大于end了，终止
			break
		}
		rChild := p*2 + 2 // 右孩子
		if rChild <= end && h.Less(rChild, child) {
			child = rChild // 左孩子小于右孩子，直接将孩子下标赋值为右孩子下标
		}
		if h.Less(child, p) { // 当前节点小于孩子节点则交换
			h.Swap(p, child)
			p = child
		} else {
			break // 当前节点大于等于孩子节点，则不需要再调整了
		}
	}
}
