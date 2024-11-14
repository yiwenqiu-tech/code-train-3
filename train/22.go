package train

var n22 int
var res22 []string

func generateParenthesis(n int) []string {
	n22 = n
	res22 = nil
	dfsFor22(0, "")
	return res22
}

func dfsFor22(curPos int, curStr string) {
	if curPos >= 2*n22 {
		if valid22(curStr) {
			res22 = append(res22, curStr)
		}
		return
	}
	if !valid22(curStr) {
		return
	}
	dfsFor22(curPos+1, curStr+"(")
	dfsFor22(curPos+1, curStr+")")
}

func valid22(str string) bool {
	var tmp int
	var left int
	var right int
	for index := range str {
		if str[index] == '(' {
			tmp++
			left++
		} else {
			tmp--
			right++
		}
		if tmp < 0 {
			return false
		}
		if left > n22 {
			return false
		}
		if right > n22 {
			return false
		}
	}
	return true
}
