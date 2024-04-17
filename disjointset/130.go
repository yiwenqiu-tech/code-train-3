package disjointset

var x130 = []int{1, 0}
var y130 = []int{0, 1}

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	ds := MakeSet(m*n + 1) // 多留一个结点，作为边缘结点
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			index := (i * n) + j
			if board[i][j] == 'O' {
				if i == 0 || i == m-1 || j == 0 || j == n-1 {
					ds.UnionSet(index, m*n) // 边界点与m*n+1合并
				}
				for k := 0; k < 2; k++ { // 只需要统计下和右两个方向
					if i+x130[k] >= m || j+y130[k] >= n {
						continue // 越界直接continue
					}
					if board[i+x130[k]][j+y130[k]] == 'O' {
						ds.UnionSet(index, (i+x130[k])*n+(j+y130[k]))
					}
				}
			}
		}
	}
	mnRoot := ds.Find(m * n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ds.Find(i*n+j) == mnRoot { // 处于边界的不变
				continue
			} else {
				if board[i][j] == 'O' { // 处于中间的变
					board[i][j] = 'X'
				}
			}
		}
	}
}
