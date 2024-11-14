package train

var put52 []int
var n52 int
var res52 int

func totalNQueens(n int) int {
	put52 = make([]int, n)
	for i := range put52 {
		put52[i] = -1
	}
	n52 = n
	res52 = 0
	dfsFor52(0)
	return res52
}

func dfsFor52(cur int) {
	if cur == n52 {
		res52++
		return
	}
	valueCur := findValid52(cur)
	for _, value := range valueCur {
		put52[cur] = value
		dfsFor52(cur + 1)
		put52[cur] = -1
	}
}

func findValid52(cur int) []int {
	var res []int
	if cur == 0 { // 第0行时，都合法
		for i := 0; i < n52; i++ {
			res = append(res, i)
		}
		return res
	}

	for i := 0; i < n52; i++ { // 遍历0-n列位置
		var valid = true
		for j := 0; j < cur; j++ { // 0-cur行
			if put52[j] == i {
				valid = false
				break
			}
			if abs(cur-j) == abs(put52[j]-i) {
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
