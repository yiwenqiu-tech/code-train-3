package dp

import "math"

func maxProfit714(prices []int, fee int) int {
	//f[i][j] 表示到第i个prices时，持有或不持有的最高收益
	var f = make([][]int, 2)
	for index := range f {
		f[index] = make([]int, 2)
		for i := 0; i < 2; i++ {
			f[index][i] = math.MinInt32
		}
	}
	f[0][0] = 0
	for i := 1; i <= len(prices); i++ {
		// 不变
		f[i&1][0] = f[(i-1)&1][0]
		f[i&1][1] = f[(i-1)&1][1]

		//买
		f[i&1][1] = max(f[i&1][1], f[(i-1)&1][0]-prices[i-1])

		//卖
		f[i&1][0] = max(f[i&1][0], f[(i-1)&1][1]+prices[i-1]-fee)

		// 重新初始化复用空间
		f[(i-1)&1][0] = math.MinInt32
		f[(i-1)&1][1] = math.MinInt32
	}
	return f[len(prices)&1][0]
}
