package train

var x200 = []int{-1, 0, 1, 0}
var y200 = []int{0, -1, 0, 1}

func numIslands(grid [][]byte) int {
	m200 = len(grid)
	n200 = len(grid[0])
	target := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				dfsFor200(i, j, grid)
				target++
			}
		}
	}
	return target
}

func dfsFor200(i, j int, grid [][]byte) {
	if i < 0 || i >= m200 || j < 0 || j >= n200 {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	for k := 0; k < 4; k++ {
		dfsFor200(i+x200[k], j+y200[k], grid)
	}
}

type UnionSet200 struct {
	fa    []int
	count int
}

func MakeUnionSet200(n int) *UnionSet200 {
	us := UnionSet200{}
	us.count = n
	us.fa = make([]int, n)
	for i := 0; i < len(us.fa); i++ {
		us.fa[i] = i
	}
	return &us
}

func (u *UnionSet200) FindFa(x int) int {
	if x == u.fa[x] {
		return x
	}
	fafa := u.FindFa(u.fa[x])
	u.fa[x] = fafa
	return fafa
}

func (u *UnionSet200) Union(x, y int) {
	xFa := u.FindFa(x)
	yFa := u.FindFa(y)
	if xFa == yFa {
		return
	}
	u.fa[yFa] = xFa
	u.count--
}

var m200 int
var n200 int
var x2002 = []int{1, 0}
var y2002 = []int{0, 1}

func numIslands2(grid [][]byte) int {
	m200 = len(grid)
	n200 = len(grid[0])
	space := 0
	us200 := MakeUnionSet200(m200 * n200)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				space++
			} else {
				for k := 0; k < 2; k++ {
					newI := i + x2002[k]
					newJ := j + y2002[k]
					if newI >= m200 || newJ >= n200 || grid[newI][newJ] == '0' {
						continue
					}
					us200.Union(i*n200+j, newI*n200+newJ)
				}
			}
		}
	}
	return us200.count - space
}
