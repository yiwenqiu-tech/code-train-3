package train

var m, n int
var x = []int{0, 1, 0, -1}
var y = []int{1, 0, -1, 0}
var visited []bool
var inputBoard [][]byte
var wordMap map[string]struct{}
var ans []string
var ansMap map[string]struct{}

func findWords(board [][]byte, words []string) []string {
	inputBoard = board
	m = len(board)
	n = len(board[0])
	ans = make([]string, 0, len(words))
	visited = make([]bool, m*n)
	wordMap = make(map[string]struct{})
	for index := range words {
		wordMap[words[index]] = struct{}{}
	}
	ansMap = make(map[string]struct{})
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, "")
		}
	}
	return ans
}

func dfs(i int, j int, str string) {
	if i < 0 || i >= m {
		return
	}
	if j < 0 || j >= n {
		return
	}
	if visited[pos(i, j)] {
		return
	}
	visited[pos(i, j)] = true
	str += string(inputBoard[i][j])

	if _, exist := wordMap[str]; exist {
		if _, exist2 := ansMap[str]; !exist2 {
			ans = append(ans, str)
			ansMap[str] = struct{}{}
		}
	}
	for z := 0; z < 4; z++ {
		dfs(i+y[z], j+x[z], str)
	}
	str = str[:len(str)-1]
	visited[pos(i, j)] = false // 还原现场
	return
}

func pos(i, j int) int {
	return i*n + j
}
