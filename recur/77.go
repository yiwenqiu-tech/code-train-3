package recur

var combineRes [][]int

func combine(n int, k int) [][]int {
	combineRes = nil
	recurForCombine(n, k, 1, []int{})
	return combineRes
}

func recurForCombine(n int, k int, index int, input []int) {
	if len(input) >= k { // 剪枝，满足条件就直接添加到结果，并return
		var tmpRes = make([]int, len(input))
		copy(tmpRes, input)
		combineRes = append(combineRes, tmpRes)
		return
	}
	if index > n { // 超出n则直接返回
		return
	}
	recurForCombine(n, k, index+1, input)
	input = append(input, index)
	recurForCombine(n, k, index+1, input)
}
