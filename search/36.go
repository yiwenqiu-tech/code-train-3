package search

var n36 int
var row36 [][10]bool // 若某一行的 某个index为true，代表有这个元素了。
var col36 [][10]bool
var box36 [][10]bool

func isValidSudoku(board [][]byte) bool {
	n36 = len(board)
	row36 = make([][10]bool, 9)
	col36 = make([][10]bool, 9)
	box36 = make([][10]bool, 9)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '0'
			if row36[i][num] {
				return false
			}
			if col36[j][num] {
				return false
			}
			if box36[boxIndex(i, j)][num] {
				return false
			}
			// 赋值
			row36[i][num] = true
			col36[j][num] = true
			box36[boxIndex(i, j)][num] = true
		}
	}
	return true
}
