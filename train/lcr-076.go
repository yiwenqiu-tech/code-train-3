package train

import (
	"container/heap"
	"math/rand"
)

type intHeap []int

func (ih *intHeap) Len() int {
	return len(*ih)
}

func (ih *intHeap) Less(i, j int) bool {
	return (*ih)[i] > (*ih)[j]
}

func (ih *intHeap) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
}

func (ih *intHeap) Push(x any) {
	*ih = append(*ih, x.(int))
}

func (ih *intHeap) Pop() any {
	item := (*ih)[len(*ih)-1]
	*ih = (*ih)[:len(*ih)-1]
	return item
}

func findKthLargest(nums []int, k int) int {
	iHeap := intHeap(nums)
	heap.Init(&iHeap)
	var res int
	for k > 0 {
		res = heap.Pop(&iHeap).(int)
		k--
	}
	return res
}

func findKthLargest2(nums []int, k int) int {
	return fastChoose(nums, 0, len(nums)-1, len(nums)-k)
}

func fastChoose(nums []int, left int, right int, index int) int {
	pos := partition(nums, left, right)
	if pos == index {
		return nums[pos]
	}
	if pos < index {
		return fastChoose(nums, pos+1, right, index)
	} else {
		return fastChoose(nums, left, pos-1, index)
	}
}

func partition(nums []int, left int, right int) int {
	randomPos := left + rand.Intn(right-left+1)
	pivot := nums[randomPos]
	nums[left], nums[randomPos] = nums[randomPos], nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		if right <= left {
			break
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		if left >= right {
			break
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}
