package sort

import (
	"container/heap"
)

type intHeap []int

func (h *intHeap) Len() int {
	return len(*h)
}

func (h *intHeap) Less(i, j int) bool {
	//return (*h)[i] < (*h)[j] // 小根堆
	return (*h)[i] > (*h)[j] // 大根堆
}

func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	item := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return item
}

func HeapSort(nums []int) []int { // 小根堆做法额外开了O(N)的空间
	numHeap := intHeap(nums)
	heap.Init(&numHeap)

	var res []int
	for numHeap.Len() > 0 {
		res = append(res, heap.Pop(&numHeap).(int))
	}

	return res
}

func HeapSort2(nums []int) []int { // 大根堆做法不需要额外开空间，每次Pop都将最大值放到最后一个元素了
	numHeap := intHeap(nums)
	heap.Init(&numHeap)

	for numHeap.Len() > 0 {
		heap.Pop(&numHeap)
	}

	return nums
}
