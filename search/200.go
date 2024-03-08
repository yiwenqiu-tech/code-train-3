package search

var m int
var n int
var visited [][]bool
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}
var inputGrid [][]byte

func numIslands(grid [][]byte) int {
	inputGrid = grid
	m = len(grid)    // m行
	n = len(grid[0]) // n列
	visited = make([][]bool, m)
	for index := range visited {
		visited[index] = make([]bool, n)
	}

	var count int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && !visited[i][j] { // 遍历到陆地，且之前没有遍历过的
				dfsForNumIslands(i, j)
				count++
			}
		}
	}
	return count
}

func dfsForNumIslands(i, j int) {
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	for index := 0; index < 4; index++ {
		nx := i + dx[index]
		ny := j + dy[index]

		if nx >= 0 && ny >= 0 && nx < m && ny < n && !visited[nx][ny] && inputGrid[nx][ny] == '1' {
			visited[nx][ny] = true
			dfsForNumIslands(nx, ny)
		}
	}
}
