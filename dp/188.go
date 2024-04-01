package dp

import (
	"code-train-3/common"
	"math"
)

func maxProfit188(k int, prices []int) int {
	var f = make([][][]int, 2)
	for i := range f {
		f[i] = make([][]int, 2)
		for j := range f[i] {
			f[i][j] = make([]int, k+1)
			for k := range f[i][j] {
				f[i][j][k] = math.MinInt32
			}
		}
	}
	f[0][0][0] = 0
	for i := 1; i <= len(prices); i++ {
		for j := 0; j < 2; j++ {
			for l := 0; l <= k; l++ {
				f[i&1][j][l] = f[(i-1)&1][j][l] // 不变
				if l > 0 {                      // f[i][1][0] 非法
					f[i&1][1][l] = common.MaxInt(f[(i-1)&1][0][l-1]-prices[i-1], f[i&1][1][l]) // 买
				}
				f[i&1][0][l] = common.MaxInt(f[(i-1)&1][1][l]+prices[i-1], f[i&1][0][l]) // 卖
			}
		}
	}
	var ans = math.MinInt32
	for l := 0; l <= k; l++ {
		if f[len(prices)&1][0][l] > ans {
			ans = f[len(prices)&1][0][l]
		}
	}
	return ans
}
