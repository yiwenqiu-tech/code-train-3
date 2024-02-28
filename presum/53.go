package presum

import "math"

func maxSubArray(nums []int) int {
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0

	for index, num := range nums {
		preSum[index+1] = preSum[index] + num
	}

	// 区间和 = 前缀和的差值，所以需要找到最大值，则找到对应的前缀最小值
	preMin := make([]int, len(nums)+1)
	preMin[0] = 0
	for i := 1; i < len(preSum); i++ {
		preMin[i] = int(math.Min(float64(preMin[i-1]), float64(preSum[i])))
	}

	var max = math.MinInt32
	for i := 1; i < len(preSum); i++ {
		tmp := preSum[i] - preMin[i-1]
		if tmp > max {
			max = tmp
		}
	}
	return max
}
