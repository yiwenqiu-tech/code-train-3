package sort

import "math/rand"

// TODO 理解思想！
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func partition(nums []int, left int, right int) int {
	pivotIndex := left + rand.Intn(right-left+1)                // 随机一个数
	pivotVal := nums[pivotIndex]                                // 将值取出来
	nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left] // 将值放到第一位里，方便后面循环，由于值在前面，接下来先从后面找
	for left < right {
		for left < right && nums[right] >= pivotVal {
			right--
		}
		if left == right {
			break
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivotVal {
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
