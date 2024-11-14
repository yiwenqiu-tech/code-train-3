package train

var m695 int
var n695 int
var x695 = []int{0, 1, 0, -1}
var y695 = []int{-1, 0, 1, 0}
var visit [][]bool

func maxAreaOfIsland(grid [][]int) int {
	m695 = len(grid)
	n695 = len(grid[0])
	visit = make([][]bool, m695)
	for i := range visit {
		visit[i] = make([]bool, n695)
		for j := range visit[i] {
			visit[i][j] = false
		}
	}
	maxArea := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 || visit[i][j] {
				continue
			} else {
				res := dfsFor695(i, j, grid)
				if res > maxArea {
					maxArea = res
				}
			}
			visit[i][j] = true
		}
	}
	return maxArea
}

func dfsFor695(i, j int, grid [][]int) int {
	if i < 0 || i >= m695 || j < 0 || j >= n695 {
		return 0
	}
	if grid[i][j] != 1 {
		return 0
	}
	if visit[i][j] {
		return 0
	}
	visit[i][j] = true
	var sum = 1
	for k := 0; k < 4; k++ {
		sum += dfsFor695(i+x695[k], j+y695[k], grid)
	}
	return sum
}
