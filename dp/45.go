package dp

import (
	"math"
)

func jump(nums []int) int {
	var f = make([][]int, len(nums)) // 表示前i个数到达j时花费的最小步数
	for i := range f {
		f[i] = make([]int, len(nums))
		for j := range f[i] {
			f[i][j] = math.MaxInt32
		}
	}
	f[0][0] = 0
	for i := 1; i < min(nums[0], len(nums)); i++ {
		f[0][i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			f[i][j] = f[i-1][j]
			for k := 0; k <= nums[i-1]; k++ {
				if j >= k {
					f[i][j] = min(f[i][j], f[i-1][j-k]+1)
				}
			}
		}
	}
	return f[len(nums)-1][len(nums)-1]
}
