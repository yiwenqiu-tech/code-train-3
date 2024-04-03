package dp

import "math"

// 可以看成是完全背包问题，具有N种不同面值coins，每枚价值为1，每种具有无限枚，求总额为amount时，求最小价值。
func coinChange(coins []int, amount int) int {
	var f = make([]int, amount+1)
	for i := range f {
		f[i] = math.MaxInt32
	}
	f[0] = 0
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			f[j] = min(f[j], f[j-coins[i]]+1)
		}
	}
	if f[amount] == math.MaxInt32 {
		return -1
	}
	return f[amount]
}
