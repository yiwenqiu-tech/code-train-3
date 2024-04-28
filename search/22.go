package search

var ans22 []string
var str string
var n22 int

func generateParenthesis(n int) []string {
	n22 = n
	ans22 = nil
	dfsForGenerateParenthesis(0)
	return ans22
}

func dfsForGenerateParenthesis(pos int) {
	if !isValid22() {
		return
	}
	if pos == n22*2 {
		ans22 = append(ans22, str)
		return
	}
	str += "("
	dfsForGenerateParenthesis(pos + 1)
	str = str[0 : len(str)-1]

	str += ")"
	dfsForGenerateParenthesis(pos + 1)
	str = str[0 : len(str)-1]
}

func isValid22() bool {
	var left int
	var totalLeft int
	for _, c := range str {
		if c == '(' {
			left++
			totalLeft++
		} else {
			left--
		}
	}
	if totalLeft > n22 { // 总的left数量>n了，非法。
		return false
	}
	return left >= 0
}
