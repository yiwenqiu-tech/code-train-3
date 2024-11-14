package train

var m980 int
var n980 int
var x980 = []int{0, 1, 0, -1}
var y980 = []int{1, 0, -1, 0}
var grid980 [][]int
var visited980 [][]bool
var visitCount980 int

func uniquePathsIII(grid [][]int) int {
	grid980 = grid
	m980 = len(grid)
	n980 = len(grid[0])
	visited980 = make([][]bool, m980)
	for i := range visited980 {
		visited980[i] = make([]bool, n980)
	}
	visitCount980 = 0
	var startI int
	var startJ int
	for i := range grid980 {
		for j := range grid980[i] {
			if grid980[i][j] != -1 && grid980[i][j] != 2 {
				visitCount980++
			}
			if grid980[i][j] == 1 {
				startI = i
				startJ = j
			}
		}
	}
	return dfsFor980(startI, startJ)
}

func checkVisitedCount() bool {
	var tmp int
	for i := range visited980 {
		for j := range visited980[i] {
			if visited980[i][j] {
				tmp++
			}
		}
	}
	return tmp == visitCount980
}

func dfsFor980(i, j int) int {
	if i < 0 || i >= m980 {
		return 0
	}
	if j < 0 || j >= n980 {
		return 0
	}
	if grid980[i][j] == -1 {
		return 0
	}
	if visited980[i][j] {
		return 0
	}
	if grid980[i][j] == 2 {
		if checkVisitedCount() {
			return 1
		}
	}
	var sum = 0
	visited980[i][j] = true
	for k := 0; k < 4; k++ {
		sum += dfsFor980(i+x980[k], j+y980[k])
	}
	visited980[i][j] = false
	return sum
}
