package train

var put51 []int
var n51 int
var res51 [][]string

func solveNQueens(n int) [][]string {
	put51 = make([]int, n)
	for i := range put51 {
		put51[i] = -1
	}
	res51 = nil
	n51 = n
	dfsFor51(0)
	return res51
}

func dfsFor51(cur int) {
	if cur == n51 {
		res51 = append(res51, generateResStr())
		return
	}
	valueCur := findValid51(cur)
	for _, value := range valueCur {
		put51[cur] = value
		dfsFor51(cur + 1)
		put51[cur] = -1
	}
}

func generateResStr() []string {
	var res = make([]string, n51)
	for i := range put51 {
		var tmp = make([]byte, n51)
		for j := 0; j < n51; j++ {
			if j != put51[i] {
				tmp[j] = '.'
			} else {
				tmp[j] = 'Q'
			}
		}
		res[i] = string(tmp)
	}
	return res
}

func findValid51(cur int) []int {
	var res []int
	if cur == 0 { // 第0行时，都合法
		for i := 0; i < n51; i++ {
			res = append(res, i)
		}
		return res
	}

	for i := 0; i < n51; i++ { // 遍历0-n列位置
		var valid = true
		for j := 0; j < cur; j++ { // 0-cur行
			if put51[j] == i {
				valid = false
				break
			}
			if abs(cur-j) == abs(put51[j]-i) {
				valid = false
				break
			}
		}
		if valid {
			res = append(res, i)
		}
	}
	return res
}
