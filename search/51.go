package search

import "strings"

// 方法1，通过运算求是否合法
/*var ans51 [][]string
var n51 int

func solveNQueens(n int) [][]string {
	n51 = n
	ans51 = nil
	dfsForSolveNQueens(nil)
	return ans51
}

func dfsForSolveNQueens(cur []int) { // cur表示第i行选了哪一列
	if len(cur) == n51 {
		ans51 = append(ans51, transferPosStr(cur))
		return
	}
	for i := 0; i < n51; i++ {
		if valid51(cur, i) {
			cur = append(cur, i)
			dfsForSolveNQueens(cur)
			cur = cur[0 : len(cur)-1]
		}
	}
}

func transferPosStr(cur []int) []string {
	var res []string
	for _, c := range cur {
		var curStr strings.Builder
		for j := 0; j < n51; j++ {
			if j == c {
				curStr.WriteByte('Q')
			} else {
				curStr.WriteByte('.')
			}
		}
		res = append(res, curStr.String())
	}
	return res
}

func valid51(cur []int, pos int) bool {
	curRow := len(cur)
	for index, i := range cur {
		if pos == i { // 同一列冲突
			return false
		}
		if abs(i-pos) == abs(curRow-index) { // 对角线冲突
			return false
		}
	}
	return true
}

func abs(value int) int {
	if value < 0 {
		return -value
	} else {
		return value
	}
}*/

// 方法2，通过位运算求是否合法
var ans51 [][]string
var n51 int
var usedCol int  // 已用的列
var usedDia int  // 往左下的对角线
var usedDia2 int // 往右下的对角线

func solveNQueens(n int) [][]string {
	n51 = n
	ans51 = nil
	usedCol = 0
	usedDia = 0
	usedDia2 = 0
	dfsForSolveNQueens(nil)
	return ans51
}

func dfsForSolveNQueens(cur []int) { // cur表示第i行选了哪一列
	if len(cur) == n51 {
		ans51 = append(ans51, transferPosStr(cur))
		return
	}
	lastUsedDia := usedDia
	lastUsedDia2 := usedDia2
	state := (1<<n51 - 1) - (usedCol | usedDia | usedDia2)
	for i := 0; i < n51; i++ {
		if (state & (1 << i)) > 0 {
			usedCol = usedCol | (1 << i)
			usedDia = usedDia | (1 << i)
			usedDia2 = usedDia2 | (1 << i)
			usedDia2 = usedDia2 >> 1
			usedDia = usedDia << 1
			usedDia = usedDia & (1<<n51 - 1)
			cur = append(cur, i)
			dfsForSolveNQueens(cur)
			usedCol = usedCol ^ (1 << i)
			usedDia = lastUsedDia
			usedDia2 = lastUsedDia2
			cur = cur[0 : len(cur)-1]
		}
	}
}

func transferPosStr(cur []int) []string {
	var res []string
	for _, c := range cur {
		var curStr strings.Builder
		for j := 0; j < n51; j++ {
			if j == c {
				curStr.WriteByte('Q')
			} else {
				curStr.WriteByte('.')
			}
		}
		res = append(res, curStr.String())
	}
	return res
}

//var queensRes [][]int
//
//var size int
//
//func solveNQueens(n int) [][]string {
//	queensRes = nil
//	size = n
//	dfsForSolveNQueens(nil)
//	fmt.Println(queensRes)
//	return transferToRes(queensRes)
//}
//
//func dfsForSolveNQueens(curRes []int) {
//	if len(curRes) == size {
//		tmp := make([]int, len(curRes))
//		copy(tmp, curRes)
//		queensRes = append(queensRes, tmp)
//		return
//	}
//	for i := 0; i < size; i++ {
//		if len(curRes) == 0 || valid(curRes, i) {
//			curRes = append(curRes, i)
//
//			dfsForSolveNQueens(curRes)
//
//			curRes = curRes[:len(curRes)-1]
//		}
//	}
//
//}
//
//func valid(curRes []int, i int) bool {
//	curRow := len(curRes) // 当前层
//	for index, c := range curRes {
//		if c == i { // 列相等，则不合法
//			return false
//		}
//		if math.Abs(float64(curRow)-float64(index)) == math.Abs(float64(i)-float64(c)) { // 在斜对角线
//			return false
//		}
//	}
//	return true
//}
//
//func transferToRes(queensRes [][]int) [][]string {
//	var res = make([][]string, len(queensRes))
//	for index := range res {
//		res[index] = make([]string, len(queensRes[index]))
//		for i := range res[index] {
//			res[index][i] = strings.Repeat(".", size)
//		}
//	}
//	for i := range queensRes {
//		for j, index := range queensRes[i] {
//			bytes := []byte(res[i][j])
//			bytes[index] = 'Q'
//			res[i][j] = string(bytes)
//		}
//	}
//	return res
//}
