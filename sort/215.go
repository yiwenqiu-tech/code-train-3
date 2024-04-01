package sort

import (
	"container/heap"
	"math/rand"
)

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

func findKthLargest2(nums []int, k int) int {
	return findK(nums, k, 0, len(nums)-1)
}

func findK(nums []int, k, left, right int) int {
	if left >= right {
		return nums[left]
	}
	pivotIndex := partitionNums(nums, left, right)
	if pivotIndex == k-1 {
		return nums[pivotIndex]
	} else if pivotIndex < k-1 {
		return findK(nums, k, pivotIndex+1, right)
	} else {
		return findK(nums, k, left, pivotIndex-1)
	}
}

func partitionNums(nums []int, left, right int) int {
	pivotIndex := left + rand.Intn(right-left+1)
	pivotVal := nums[pivotIndex]
	nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left]
	for left < right {
		for left < right && nums[right] <= pivotVal {
			right--
		}
		if right == left {
			break
		}
		nums[left] = nums[right]
		for left < right && nums[left] >= pivotVal {
			left++
		}
		if left == right {
			break
		}
		nums[right] = nums[left]
	}
	nums[left] = pivotVal
	return left
}
