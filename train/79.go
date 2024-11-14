package train

var board79 [][]byte
var word79 string
var x79 = []int{0, 1, 0, -1}
var y79 = []int{1, 0, -1, 0}
var m79 = 0
var n79 = 0
var visited79 [][]bool

func exist(board [][]byte, word string) bool {
	board79 = board
	word79 = word
	m79 = len(board)
	n79 = len(board[0])
	visited79 = make([][]bool, m79)
	for i := range visited79 {
		visited79[i] = make([]bool, n79)
	}
	for i := range board {
		for j := range board[i] {
			res := dfsFor79(i, j, 0)
			if res {
				return true
			}
		}
	}
	return false
}

func dfsFor79(i, j int, k int) bool {
	if i < 0 || i >= m79 {
		return false
	}
	if j < 0 || j >= n79 {
		return false
	}
	if k >= len(word79) {
		return false
	}
	if board79[i][j] != word79[k] {
		return false
	}
	if visited79[i][j] {
		return false
	}
	if k == len(word79)-1 {
		return true
	}

	visited79[i][j] = true
	for l := 0; l < 4; l++ {
		res := dfsFor79(i+x79[l], j+y79[l], k+1)
		if res {
			return true
		}
	}
	visited79[i][j] = false
	return false
}

/*var m, n int
var x = []int{0, 1, 0, -1}
var y = []int{1, 0, -1, 0}
var visited []bool
var targetWord string
var inputBoard [][]byte

func exist(board [][]byte, word string) bool {
	m = len(board)
	n = len(board[0])
	inputBoard = board
	visited = make([]bool, m*n)
	targetWord = word
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(i int, j int, k int) bool {
	if i < 0 || i >= m {
		return false
	}
	if j < 0 || j >= n {
		return false
	}
	if visited[pos(i, j)] {
		return false
	}
	// 当前字母对不上了，直接返回，剪枝
	if inputBoard[i][j] != targetWord[k] {
		return false
	}
	// 遍历到最后一个字母了
	if k == len(targetWord)-1 && inputBoard[i][j] == targetWord[k] {
		return true
	}
	visited[pos(i, j)] = true
	for z := 0; z < 4; z++ {
		if dfs(i+y[z], j+x[z], k+1) {
			return true
		}
	}
	visited[pos(i, j)] = false // 还原现场
	return false
}

func pos(i, j int) int {
	return i*n + j
}*/
