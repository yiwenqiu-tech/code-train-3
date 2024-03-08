package search

import (
	"fmt"
	"math"
	"strings"
)

var queensRes [][]int

var size int

func solveNQueens(n int) [][]string {
	queensRes = nil
	size = n
	dfsForSolveNQueens(nil)
	fmt.Println(queensRes)
	return transferToRes(queensRes)
}

func dfsForSolveNQueens(curRes []int) {
	if len(curRes) == size {
		tmp := make([]int, len(curRes))
		copy(tmp, curRes)
		queensRes = append(queensRes, tmp)
		return
	}
	for i := 0; i < size; i++ {
		if len(curRes) == 0 || valid(curRes, i) {
			curRes = append(curRes, i)

			dfsForSolveNQueens(curRes)

			curRes = curRes[:len(curRes)-1]
		}
	}

}

func valid(curRes []int, i int) bool {
	curRow := len(curRes) // 当前层
	for index, c := range curRes {
		if c == i { // 列相等，则不合法
			return false
		}
		if math.Abs(float64(curRow)-float64(index)) == math.Abs(float64(i)-float64(c)) { // 在斜对角线
			return false
		}
	}
	return true
}

func transferToRes(queensRes [][]int) [][]string {
	var res = make([][]string, len(queensRes))
	for index := range res {
		res[index] = make([]string, len(queensRes[index]))
		for i := range res[index] {
			res[index][i] = strings.Repeat(".", size)
		}
	}
	for i := range queensRes {
		for j, index := range queensRes[i] {
			bytes := []byte(res[i][j])
			bytes[index] = 'Q'
			res[i][j] = string(bytes)
		}
	}
	return res
}
