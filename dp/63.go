package dp

import "code-train-3/common"

var m int
var n int
var ans int
var ansForXY [][]int

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	ans = 0
	m = len(obstacleGrid)    // 行
	n = len(obstacleGrid[0]) // 列
	ansForXY = make([][]int, m)
	for i := range ansForXY {
		ansForXY[i] = make([]int, n)
	}
	for i := range ansForXY {
		for j := range ansForXY[i] {
			ansForXY[i][j] = -1
		}
	}

	return dfsForUniquePathsWithObstacles(0, 0, obstacleGrid)
}

func dfsForUniquePathsWithObstacles(x, y int, obstacleGrid [][]int) int {
	// 1.纯暴力搜索
	/*	if obstacleGrid[x][y] == 1 {
			return
		}

		if x == m-1 && y == n-1 {
			ans++
			return
		}

		if x < m-1 {
			dfsForUniquePathsWithObstacles(x+1, y, obstacleGrid)
		}

		if y < n-1 {
			dfsForUniquePathsWithObstacles(x, y+1, obstacleGrid)
		}*/

	// 记忆化搜索
	if obstacleGrid[x][y] == 1 {
		return 0
	}
	if x == m-1 && y == n-1 {
		return 1
	}
	// 记忆化搜索
	if ansForXY[x][y] != -1 {
		return ansForXY[x][y]
	}
	var down = 0
	var right = 0
	if x < m-1 {
		down = dfsForUniquePathsWithObstacles(x+1, y, obstacleGrid)
	}

	if y < n-1 {
		right = dfsForUniquePathsWithObstacles(x, y+1, obstacleGrid)
	}
	ansForXY[x][y] = common.MaxInt(ansForXY[x][y], down+right)
	return ansForXY[x][y]
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m = len(obstacleGrid)    // 行
	n = len(obstacleGrid[0]) // 列
	var f [][]int
	f = make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := range f {
		for j := range f[i] {
			f[i][j] = -1
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				f[i][j] = 0
				continue
			}
			if i == 0 && j == 0 {
				f[i][j] = 1
			} else if i == 0 {
				f[i][j] = f[i][j-1]
			} else if j == 0 {
				f[i][j] = f[i-1][j]
			} else {
				f[i][j] = common.MaxInt(f[i][j], f[i-1][j]+f[i][j-1])
			}
		}
	}
	return f[m-1][n-1]
}
