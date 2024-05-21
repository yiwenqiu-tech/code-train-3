package search

import (
	"fmt"
	"math"
)

var board37 [][]byte

//var row37 [][10]bool // 若某一行的 某个index为true，代表有这个元素了。
//var col37 [][10]bool
//var box37 [][10]bool

var row37 [10]int
var col37 [10]int
var box37 [10]int

func solveSudoku(board [][]byte) {
	board37 = board
	row37 = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	col37 = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	box37 = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0'
				row37[i] = row37[i] | (1 << num) // 已存在时，置为1；
				col37[j] = col37[j] | (1 << num)
				box37[boxIndex(i, j)] = box37[boxIndex(i, j)] | (1 << num)
			}
		}
	}
	fmt.Println(row37)
	fmt.Println(col37)
	fmt.Println(box37)
	dfsForSolveSudoku()
}

func dfsForSolveSudoku() bool {
	i, j, optional := findLeastPossiblePos()
	if i == -1 {
		return true
	}
	for _, k := range optional {
		board37[i][j] = byte('0' + k)
		row37[i] = row37[i] | (1 << k) // 已存在时，置为1；
		col37[j] = col37[j] | (1 << k)
		box37[boxIndex(i, j)] = box37[boxIndex(i, j)] | (1 << k)
		if dfsForSolveSudoku() { // 如果为true时，立即返回，不需要还原现场了
			return true
		}
		board37[i][j] = '.'
		row37[i] = row37[i] & (^(1 << k)) // 还原现场，置为0
		col37[j] = col37[j] & (^(1 << k))
		box37[boxIndex(i, j)] = box37[boxIndex(i, j)] & (^(1 << k))
	}
	return false
}

func findNextEmptyPos() (int, int) {
	for i := 0; i < len(board37); i++ {
		for j := 0; j < len(board37[0]); j++ {
			if board37[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}

func findLeastPossiblePos() (int, int, []int) {
	minCount := math.MaxInt32
	var resI = -1
	var resJ = -1
	var optional []int
	for i := 0; i < len(board37); i++ {
		for j := 0; j < len(board37[0]); j++ {
			if board37[i][j] == '.' {
				rowChosen := row37[i]
				colChosen := col37[j]
				boxChosen := box37[boxIndex(i, j)]
				tmp := rowChosen | colChosen | boxChosen // 取或后，为0的位置就是可选
				var curOptional []int
				for k := 1; k <= 9; k++ {
					if (tmp>>k)&1 == 0 {
						curOptional = append(curOptional, k)
					}
				}
				if len(curOptional) < minCount {
					resI = i
					resJ = j
					optional = curOptional
					minCount = len(curOptional)
				}
			}
		}
	}
	return resI, resJ, optional
}

func boxIndex(i, j int) int {
	return i/3*3 + j/3
}
