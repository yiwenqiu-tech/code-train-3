package search

var dx329 = []int{0, 1, 0, -1}
var dy329 = []int{1, 0, -1, 0}
var m329 int
var n329 int
var input329 [][]int
var dist [][]int // 记录从某个位置出发，能走到最远的长度

func longestIncreasingPath(matrix [][]int) int {
	m329 = len(matrix)
	n329 = len(matrix[0])
	input329 = matrix
	dist = make([][]int, m329)
	for index := range dist {
		dist[index] = make([]int, n329) // 初始化dist
	}

	var ans int
	for i := 0; i < m329; i++ {
		for j := 0; j < n329; j++ {
			curAns := dfsForLongestIncreasingPath(i, j)
			if curAns > ans {
				ans = curAns
			}
		}
	}
	return ans
}

func dfsForLongestIncreasingPath(i, j int) int {
	if dist[i][j] != 0 {
		return dist[i][j] // 已经计算过了
	}
	dist[i][j] = 1 // 当前点就是1
	for k := 0; k < 4; k++ {
		nx := i + dx329[k]
		ny := j + dy329[k]
		if nx >= 0 && ny >= 0 && nx < m329 && ny < n329 && input329[nx][ny] > input329[i][j] {
			nextPointLength := dfsForLongestIncreasingPath(nx, ny) + 1
			if dist[i][j] < nextPointLength {
				dist[i][j] = nextPointLength
			}
		}
	}
	return dist[i][j]
}

//var to [][]int
//var in []int
// 广度优先搜索解法
/*func longestIncreasingPath(matrix [][]int) int {
	// 构建图，含出边数组以及入度数组
	m329 = len(matrix)
	n329 = len(matrix[0])
	to = make([][]int, m329*n329)
	in = make([]int, m329*n329)
	maxLength := make([]int, m329*n329)
	for i := 0; i < m329; i++ {
		for j := 0; j < n329; j++ {
			for k := 0; k < 4; k++ {
				cx := i + dx329[k]
				cy := j + dy329[k]
				if cx < 0 || cy < 0 || cx >= m329 || cy >= n329 {
					continue
				}
				if matrix[i][j] < matrix[cx][cy] {
					to[calIndex(i, j)] = append(to[calIndex(i, j)], calIndex(cx, cy))
					in[calIndex(cx, cy)]++
				}
			}
		}
	}

	// 将入度为0的入队，从这些节点开始遍历
	queue := list.List{}
	for index, v := range in {
		if v == 0 { // 入度为0的节点入队列
			queue.PushBack(index)
			maxLength[index] = 1 //只有一个时，路径长度为1
		}
	}
	for queue.Len() != 0 {
		head := queue.Front()
		headVal := head.Value.(int)
		queue.Remove(head)

		toOfHead := to[headVal]
		for _, t := range toOfHead {
			in[t]--
			// 通过maxLength记录 到达点t经过多少个点的最大值
			maxLength[t] = int(math.Max(float64(maxLength[t]), float64(maxLength[headVal]+1)))
			if in[t] == 0 {
				queue.PushBack(t)
			}
		}
	}
	var ans int
	for _, v := range maxLength {
		ans = int(math.Max(float64(v), float64(ans)))
	}
	return ans
}
*/
func calIndex(i, j int) int {
	return i*n329 + j
}
