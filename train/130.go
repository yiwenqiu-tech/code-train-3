package train

var x130 = []int{0, 1, 0, -1}
var y130 = []int{-1, 0, 1, 0}
var m130 int
var n130 int

func solve(board [][]byte) {
	m130 = len(board)
	n130 = len(board[0])
	for i := range board {
		for j := range board[i] {
			if i == 0 || i == m130-1 || j == 0 || j == n130-1 { // 只从边缘开始搜
				if board[i][j] == 'O' {
					dfsFor130(i, j, board)
				}
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'S' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func dfsFor130(i, j int, board [][]byte) {
	if i < 0 || i >= m130 || j < 0 || j >= n130 {
		return
	}
	if board[i][j] == 'X' || board[i][j] == 'S' {
		return
	}
	board[i][j] = 'S' // 表示已经搜过了
	for k := 0; k < 4; k++ {
		dfsFor130(i+x130[k], j+y130[k], board)
	}
}
