package train

var col37 [][]bool
var row37 [][]bool
var box37 [][]bool
var board37 [][]byte

func solveSudoku(board [][]byte) {
	col37 = make([][]bool, 9)
	for i := range col37 {
		col37[i] = make([]bool, 10)
	}
	row37 = make([][]bool, 9)
	for i := range row37 {
		row37[i] = make([]bool, 10)
	}
	box37 = make([][]bool, 9)
	for i := range box37 {
		box37[i] = make([]bool, 10)
	}
	board37 = board
	// 初始化
	for i := range board { // i行
		for j := range board[i] { //j列
			if board[i][j] == '.' {
				continue
			}
			value := board[i][j] - '0'
			col37[i][value] = true
			row37[j][value] = true
			box37[getBoxIndex(i, j)][value] = true
		}
	}
	dfsFor37()
	board = board37
}

func dfsFor37() bool {
	i, j := findNextEmpty()
	if i == -1 {
		return true
	}
	for k := 1; k < 10; k++ { // 从1-10去试
		if col37[i][k] || row37[j][k] || box37[getBoxIndex(i, j)][k] {
			continue
		}
		board37[i][j] = byte(k + '0')
		col37[i][k] = true
		row37[j][k] = true
		box37[getBoxIndex(i, j)][k] = true
		if dfsFor37() == true {
			return true
		}
		board37[i][j] = '.'
		col37[i][k] = false
		row37[j][k] = false
		box37[getBoxIndex(i, j)][k] = false
	}
	return false
}

func findNextEmpty() (int, int) {
	for i := range board37 {
		for j := range board37[i] {
			if board37[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}

func getBoxIndex(i, j int) int {
	if i < 3 {
		if j < 3 {
			return 0
		} else if j < 6 {
			return 1
		} else {
			return 2
		}
	} else if i < 6 {
		if j < 3 {
			return 3
		} else if j < 6 {
			return 4
		} else {
			return 5
		}
	} else {
		if j < 3 {
			return 6
		} else if j < 6 {
			return 7
		} else {
			return 8
		}
	}
}
