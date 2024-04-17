package disjointset

func findCircleNum(isConnected [][]int) int {
	disjointSet := MakeSet(len(isConnected))
	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected[i]); j++ {
			if i == j {
				continue
			}
			if isConnected[i][j] == 1 {
				disjointSet.UnionSet(i, j)
			}
		}
	}
	var ans = 0
	for i := range disjointSet.Fa {
		if disjointSet.Fa[i] == i {
			ans++
		}
	}
	return ans
}

// DFS
var visited []bool

func findCircleNum2(isConnected [][]int) int {
	n := len(isConnected)
	visited = make([]bool, n)
	var ans = 0
	for i := range isConnected {
		if !visited[i] {
			dfsForFindCircleNum2(isConnected, i)
			ans++
		}
	}
	return ans
}

func dfsForFindCircleNum2(isConnected [][]int, i int) {
	for j := range isConnected[i] {
		if isConnected[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfsForFindCircleNum2(isConnected, j)
		}
	}
}
