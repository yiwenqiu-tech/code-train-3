package array

import "math"

// ac
func removeDuplicates(nums []int) int {
	var tmp = math.MinInt32
	var curIndex = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != tmp {
			nums[curIndex] = nums[i]
			curIndex++
			tmp = nums[i]
		}
	}
	return curIndex
}
