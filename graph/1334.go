package graph

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// 初始化邻接矩阵
	var f = make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			if i == j {
				f[i][j] = 0
			} else {
				f[i][j] = math.MaxInt32
			}
		}
	}
	for _, e := range edges {
		from := e[0]
		to := e[1]
		weight := e[2]
		f[from][to] = weight
		f[to][from] = weight
	}

	// floyd算法, 滚动数组优化版本
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				f[i][j] = min(f[i][j], f[i][k]+f[k][j]) // f[k][i][j] = min(f[k-1][i][j], f[k-1][i][k] + f[k-1][k][j])
			}
		}
	}

	var ans = -1
	var minCity = math.MaxInt32
	for i := 0; i < n; i++ {
		var tmp = 0
		for j := 0; j < n; j++ {
			if f[i][j] <= distanceThreshold { // 小于阈值
				tmp++
			}
		}
		if tmp <= minCity {
			minCity = tmp
			ans = i
		}
	}
	return ans
}
