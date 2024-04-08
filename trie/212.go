package trie

var inputWords []string
var m212 int
var n212 int
var x212 = []int{0, 1, 0, -1}
var y212 = []int{1, 0, -1, 0}
var visit [][]bool
var ans212 []string
var ans212Map map[string]struct{}
var board212 [][]byte
var trie Trie

func findWords(board [][]byte, words []string) []string {
	trie = Constructor()
	board212 = board
	m212 = len(board)
	n212 = len(board[0])
	inputWords = words
	for i := range inputWords {
		trie.Insert(inputWords[i]) // 维护trie
	}
	visit = make([][]bool, m212)
	for i := range visit {
		visit[i] = make([]bool, n212)
	}
	ans212 = nil
	ans212Map = make(map[string]struct{})

	// 枚举每个起点，搜索
	for i := 0; i < m212; i++ {
		for j := 0; j < n212; j++ {
			dfsForFind(i, j, nil, trie.Root)
		}
	}

	return ans212
}

func dfsForFind(x, y int, cur []byte, node *TrieNode) {
	if x < 0 || x >= m212 || y < 0 || y >= n212 { // 越界了
		return
	}
	if visit[x][y] { // 已经访问过，直接返回
		return
	}

	if node.Next == nil { // 不存在于对应的前缀树，则直接返回了。
		return
	}

	if _, exist := node.Next[board212[x][y]]; !exist { // 不存在于对应的前缀树，则直接返回了。
		return
	}

	cur = append(cur, board212[x][y])
	visit[x][y] = true
	// 存在该字符，且不在已有结果里，放进结果
	if node.Next[board212[x][y]].Count > 0 && !inMap(string(cur), ans212Map) {
		ans212 = append(ans212, string(cur))
		ans212Map[string(cur)] = struct{}{}
	}
	if len(node.Next[board212[x][y]].Next) == 0 { // next没有子树了，且已经判断过了，则可以删掉了
		delete(node.Next, board212[x][y])
	} else {
		for i := 0; i < 4; i++ {
			dfsForFind(x+x212[i], y+y212[i], cur, node.Next[board212[x][y]])
		}
	}

	cur = cur[:len(cur)-1]
	visit[x][y] = false

}

func inMap(cur string, m map[string]struct{}) bool {
	if _, exist := m[cur]; exist {
		return true
	}
	return false
}
