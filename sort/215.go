package sort

import "container/heap"

type iHeap []int

func (h *iHeap) Len() int {
	return len(*h)
}

func (h *iHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *iHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *iHeap) Push(x any) {
	(*h) = append((*h), x.(int))
} // add x as element Len()

func (h *iHeap) Pop() any {
	item := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return item
} // remove and return element Len() - 1.

func findKthLargest(nums []int, k int) int {
	h := iHeap(nums)
	heap.Init(&h)
	ans := 0
	for i := k; i > 0; i-- {
		ans = heap.Pop(&h).(int)
	}
	return ans
}
