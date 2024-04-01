package dp

import "math"

func rob213(nums []int) int {
	if len(nums) == 1 { // 特殊情况，只有一家时,直接返回nums[0]
		return nums[0]
	}
	var f = make([][]int, 2)
	for i := range f {
		f[i] = make([]int, 2)
		for j := range f[i] {
			f[i][j] = math.MinInt32
		}
	}
	f[1][0] = 0 // 第一家不偷
	for i := 2; i <= len(nums); i++ {
		f[i&1][1] = f[(i-1)&1][0] + nums[i-1]
		f[i&1][0] = max(f[(i-1)&1][1], f[(i-1)&1][0])

		f[(i-1)&1][0] = math.MinInt32
		f[(i-1)&1][1] = math.MinInt32
	}
	ans1 := max(f[len(nums)&1][1], f[len(nums)&1][0])
	// 重新初始化f
	for i := range f {
		for j := range f[i] {
			f[i][j] = math.MinInt32
		}
	}
	f[0][0] = 0
	for i := 1; i < len(nums); i++ { // 第1家偷，最后1家不偷
		f[i&1][1] = f[(i-1)&1][0] + nums[i-1]
		f[i&1][0] = max(f[(i-1)&1][1], f[(i-1)&1][0])

		f[(i-1)&1][0] = math.MinInt32
		f[(i-1)&1][1] = math.MinInt32
	}
	ans2 := max(f[(len(nums)-1)&1][1], f[(len(nums)-1)&1][0])

	return max(ans1, ans2)
}
