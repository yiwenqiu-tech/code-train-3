package search

import (
	"math"
)

var board37 [][]byte
var row37 [][10]bool // 若某一行的 某个index为true，代表有这个元素了。
var col37 [][10]bool
var box37 [][10]bool

func solveSudoku(board [][]byte) {
	board37 = board
	row37 = make([][10]bool, len(board))
	col37 = make([][10]bool, len(board))
	box37 = make([][10]bool, len(board))

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0'
				row37[i][num] = true
				col37[j][num] = true
				box37[boxIndex(i, j)][num] = true
			}
		}
	}
	dfsForSolveSudoku()
}

func dfsForSolveSudoku() bool {
	i, j, optional := findLeastPossiblePos()
	if i == -1 {
		return true
	}
	for _, k := range optional {
		board37[i][j] = byte('0' + k)
		row37[i][k] = true
		col37[j][k] = true
		box37[boxIndex(i, j)][k] = true
		if dfsForSolveSudoku() { // 如果为true时，立即返回，不需要还原现场了
			return true
		}
		board37[i][j] = '.'
		row37[i][k] = false
		col37[j][k] = false
		box37[boxIndex(i, j)][k] = false
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
				var curOptional []int
				for k := 1; k < len(rowChosen); k++ {
					if rowChosen[k] {
						continue
					}
					if colChosen[k] {
						continue
					}
					if boxChosen[k] {
						continue
					}
					curOptional = append(curOptional, k)
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
