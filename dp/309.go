package dp

import "math"

func maxProfit309(prices []int) int {
	// f[i][j][k]  表示到第i个prices时，持有或不持有，且是否处于冷冻期的最高收益
	var f = make([][][]int, 2)
	for i := range f {
		f[i] = make([][]int, 2)
		for j := range f[i] {
			f[i][j] = make([]int, 2)
			for k := range f[i][j] {
				f[i][j][k] = math.MinInt32
			}
		}
	}
	f[0][0][0] = 0
	for i := 1; i <= len(prices); i++ {
		for j := 0; j < 2; j++ {
			// 不操作，注意冷冻期会变成0
			f[i&1][j][0] = max(f[(i-1)&1][j][1], f[(i-1)&1][j][0])

			// 买，买时不存在冷冻期
			f[i&1][1][0] = max(f[i&1][1][0], f[(i-1)&1][0][0]-prices[i-1])
			// 卖，卖时需要修改冷冻期
			f[i&1][0][1] = max(f[i&1][0][1], f[(i-1)&1][1][0]+prices[i-1])
		}
	}

	var ans = math.MinInt32
	for k := 0; k < 2; k++ {
		if f[len(prices)&1][0][k] > ans {
			ans = f[len(prices)&1][0][k]
		}
	}
	return ans
}

func maxProfit3092(prices []int) int { // 列表法，考虑下一个状态如何由当前状态转移
	// f[i][j][k]  表示到第i个prices时，持有或不持有，且是否处于冷冻期的最高收益
	var f = make([][][]int, len(prices)+1)
	for i := range f {
		f[i] = make([][]int, 2)
		for j := range f[i] {
			f[i][j] = make([]int, 2)
			for k := range f[i][j] {
				f[i][j][k] = math.MinInt32
			}
		}
	}
	f[0][0][0] = 0
	for i := 0; i < len(prices); i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				if f[i][j][k] == math.MinInt32 {
					continue // 跳过非法值
				}
				if j == 0 && k == 0 { // 只有当 没持有且不在冷冻期，可以购买
					f[i+1][1][0] = max(f[i][j][k]-prices[i], f[i+1][j][k])
				}
				if j == 1 && k == 0 { // 只有当持有股票才可以卖出
					f[i+1][0][1] = max(f[i][j][k]+prices[i], f[i+1][j][k])
				}
				// 保持不变
				f[i+1][j][0] = max(f[i][j][k], f[i+1][j][0])
			}
		}
	}
	return max(f[len(prices)][0][0], f[len(prices)][0][1])
}
