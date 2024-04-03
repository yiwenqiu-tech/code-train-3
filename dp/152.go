package dp

import (
	"math"
)

func maxProduct(nums []int) int {
	var maxF = make([]int, len(nums))
	var minF = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		maxF[i] = math.MinInt32
		minF[i] = math.MaxInt32
	}
	maxF[0] = nums[0]
	minF[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		maxF[i] = min(min(minF[i-1]*nums[i], maxF[i-1]*nums[i]), nums[i])
		minF[i] = min(min(minF[i-1]*nums[i], maxF[i-1]*nums[i]), nums[i])
	}

	var ans = math.MinInt32
	for i := 0; i < len(maxF); i++ {
		if ans < maxF[i] {
			ans = maxF[i]
		}
	}
	return ans
}
