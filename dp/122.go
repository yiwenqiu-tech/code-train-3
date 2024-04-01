package dp

import (
	"code-train-3/common"
	"math"
)

func maxProfit122(prices []int) int {
	var f = make([][]int, 2)
	for i := range f {
		f[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			f[i][j] = math.MinInt32
		}
	}
	f[0][0] = 0
	for i := 1; i <= len(prices); i++ {
		// 当前没有操作时转移
		f[i&1][0] = common.MaxInt(f[i&1][0], f[(i-1)&1][0])
		f[i&1][1] = common.MaxInt(f[i&1][1], f[(i-1)&1][1])

		f[i&1][0] = common.MaxInt(f[i&1][0], f[(i-1)&1][1]+prices[i-1])
		f[i&1][1] = common.MaxInt(f[i&1][1], f[(i-1)&1][0]-prices[i-1])
	}
	return f[len(prices)&1][0]
}
