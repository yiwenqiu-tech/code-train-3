package dp

import "math"

func rob(nums []int) int {
	var f = make([][]int, 2)
	for i := range f {
		f[i] = make([]int, 2)
		for j := range f[i] {
			f[i][j] = math.MinInt32
		}
	}
	f[0][0] = 0
	for i := 1; i <= len(nums); i++ {
		f[i&1][1] = f[(i-1)&1][0] + nums[i-1]
		f[i&1][0] = max(f[(i-1)&1][1], f[(i-1)&1][0])

		f[(i-1)&1][0] = math.MinInt32
		f[(i-1)&1][1] = math.MinInt32
	}
	return max(f[len(nums)&1][1], f[len(nums)&1][0])
}
