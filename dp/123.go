package dp

import (
	"math"
)

func maxProfit123(prices []int) int {
	var f = make([][][]int, 2)
	for i := range f {
		f[i] = make([][]int, 2)
		for j := range f[i] {
			f[i][j] = make([]int, 3)
			for k := range f[i][j] {
				f[i][j][k] = math.MinInt32
			}
		}
	}
	f[0][0][0] = 0
	for i := 1; i <= len(prices); i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k <= 2; k++ {
				f[i&1][j][k] = f[(i-1)&1][j][k] // 不变
				if k > 0 {                      // f[i][1][0] 非法
					f[i&1][1][k] = min(f[(i-1)&1][0][k-1]-prices[i-1], f[i&1][1][k]) // 买
				}
				f[i&1][0][k] = min(f[(i-1)&1][1][k]+prices[i-1], f[i&1][0][k]) // 卖
			}
		}
	}
	var ans = math.MinInt32
	for k := 0; k <= 2; k++ {
		if f[len(prices)&1][0][k] > ans {
			ans = f[len(prices)&1][0][k]
		}
	}
	return ans
}
