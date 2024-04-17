package search

var boardM int          // 行
var boardN int          // 列
var inputBoard [][]byte // 输入矩阵
var oPos [][]bool

func solve(board [][]byte) {
	boardM = len(board)
	boardN = len(board[0])
	inputBoard = board
	oPos = make([][]bool, boardM)
	for index := range oPos {
		oPos[index] = make([]bool, boardN)
	}

	// 只遍历边界的点
	for i := 0; i < boardM; i++ {
		for j := 0; j < boardN; j++ {
			if i != 0 && i != boardM-1 && j != 0 && j != boardN-1 { // 不处于边界，不需要参与计算
				continue
			}
			if board[i][j] == 'O' && !oPos[i][j] { // 点为O，且没有遍历过
				oPos[i][j] = true
				dfsForSolve(i, j)
			}
		}
	}
	// 剩下点均填充为X
	for i := range board {
		for j := range board[i] {
			if !oPos[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func dfsForSolve(i, j int) {
	if i < 0 || j < 0 || i >= boardM || j > boardN {
		return
	}
	for index := 0; index < 4; index++ {
		nx := i + dx[index]
		ny := j + dy[index]
		if nx >= 0 && ny >= 0 && nx < boardM && ny < boardN && !oPos[nx][ny] && inputBoard[nx][ny] == 'O' {
			oPos[nx][ny] = true
			dfsForSolve(nx, ny)
		}
	}
}
