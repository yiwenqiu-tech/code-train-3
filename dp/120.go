package dp

import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	n := len(triangle[m-1])
	var f = make([][]int, m)
	var preType = make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
		preType[i] = make([]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = math.MaxInt32
			preType[i][j] = -1
		}
	}
	f[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				f[i][j] = f[i-1][j] + triangle[i][j]
				preType[i][j] = 1
			} else {
				if f[i-1][j] < f[i-1][j-1] {
					f[i][j] = f[i-1][j] + triangle[i][j]
					preType[i][j] = 2
				} else {
					f[i][j] = f[i-1][j-1] + triangle[i][j]
					preType[i][j] = 3
				}
			}
		}
	}
	var ans = math.MaxInt32
	var endPos int
	for i := 0; i < n; i++ {
		if f[m-1][i] < ans {
			ans = f[m-1][i]
			endPos = i
		}
	}
	printForMinimumTotal(m-1, endPos, preType, triangle)
	return ans
}

func printForMinimumTotal(row int, col int, preType [][]int, triangle [][]int) {
	if preType[row][col] == -1 {
		fmt.Printf("%d ", triangle[row][col])
		return
	}
	switch preType[row][col] {
	case 1:
		printForMinimumTotal(row-1, col, preType, triangle)
	case 2:
		printForMinimumTotal(row-1, col, preType, triangle)
	case 3:
		printForMinimumTotal(row-1, col-1, preType, triangle)
	}
	fmt.Printf("%d ", triangle[row][col])
}

func minimumTotal2(triangle [][]int) int {
	m := len(triangle)
	n := len(triangle[m-1])
	var f = make([][]int, 2)
	for i := 0; i < 2; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = math.MaxInt32
		}
	}
	f[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				f[i&1][j] = f[(i-1)&1][j] + triangle[i][j] // 注意：golang里 &的优先级比-高，所以需要加括号
			} else {
				if f[(i-1)&1][j] < f[(i-1)&1][j-1] {
					f[i&1][j] = f[(i-1)&1][j] + triangle[i][j]
				} else {
					col := (i - 1) & 1
					row := j - 1
					newCol := i & 1
					f[newCol][j] = f[col][row] + triangle[i][j]
				}
			}
		}
	}
	var ans = math.MaxInt32
	for i := 0; i < n; i++ {
		if f[(m-1)&1][i] < ans {
			ans = f[(m-1)&1][i]
		}
	}
	return ans
}
