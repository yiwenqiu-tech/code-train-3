package dp

import (
	"fmt"
	"math"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	var preType = make([][]int, m+1)
	var f = make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
		preType[i] = make([]int, n+1)
	}
	for i := range f {
		for j := range f[i] {
			if i == 0 || j == 0 {
				f[i][j] = 0
			} else {
				f[i][j] = math.MinInt32
			}
		}
	}

	text1 = " " + text1
	text2 = " " + text2
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i] == text2[j] {
				f[i][j] = f[i-1][j-1] + 1
				preType[i][j] = 0
			} else {
				if f[i-1][j] > f[i][j-1] {
					f[i][j] = f[i-1][j]
					preType[i][j] = 1
				} else {
					f[i][j] = f[i][j-1]
					preType[i][j] = 2
				}
			}
		}
	}
	print1143(preType, text1, text2, m, n)
	return f[m][n]
}

func print1143(preType [][]int, text1, text2 string, m, n int) {
	if m == 0 || n == 0 {
		return
	}
	if preType[m][n] == 0 {
		print1143(preType, text1, text2, m-1, n-1)
		fmt.Printf("%c", text1[m]) // 当前选了m的位置
	} else if preType[m][n] == 1 {
		print1143(preType, text1, text2, m-1, n)
	} else if preType[m][n] == 2 {
		print1143(preType, text1, text2, m, n-1)
	}
}
