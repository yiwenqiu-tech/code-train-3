package disjointset

type ds []int

func (d *ds) find(x int) int {
	if (*d)[x] == x {
		return x
	}
	fa := d.find((*d)[x])
	(*d)[x] = fa
	return fa
}

func numIslands(grid [][]byte) int {
	x := []int{1, 0}
	y := []int{0, 1}

	m := len(grid)
	n := len(grid[0])

	fa := make(ds, m*n)
	for i := range fa {
		fa[i] = i
	}

	for i := 0; i < m; i++ { // 行
		for j := 0; j < n; j++ { // 列
			if grid[i][j] == '0' {
				continue
			}
			for k := 0; k < 2; k++ {
				if j+x[k] >= n || i+y[k] >= m {
					continue // 越界了，continue
				}
				if grid[i+y[k]][j+x[k]] == '1' {
					root1 := fa.find((i+y[k])*n + (j + x[k]))
					root2 := fa.find(i*n + j)
					if root1 != root2 {
						fa[root1] = root2
					}
				}
			}
		}
	}

	var ans int
	for i := 0; i < m; i++ { // 行
		for j := 0; j < n; j++ { // 列
			if grid[i][j] == '1' && fa[i*n+j] == i*n+j {
				ans++
			}
		}
	}
	return ans
}
