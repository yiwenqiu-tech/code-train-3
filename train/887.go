package train

import "math"

func superEggDrop(k int, n int) int {
	var f = make([][]int, k+1)
	for i := 0; i <= k; i++ {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		f[1][i] = i           // 只有一个鸡蛋时初始化
		f[0][i] = math.MinInt // 没有鸡蛋时，无法确定多少层
	}
	for i := 1; i <= k; i++ {
		f[i][0] = 0 // n=0时，初始化为1 // 0层为0
	}
	for i := 2; i <= k; i++ {
		for j := 1; j <= n; j++ {
			minDrop := math.MaxInt
			for z := 1; z <= j; z++ {
				minDrop = min(minDrop, 1+max(f[i][j-z], f[i-1][z-1]))
			}
			f[i][j] = minDrop
		}
	}
	return f[k][n]
}
