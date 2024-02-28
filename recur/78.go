package recur

var res [][]int

func subsets(nums []int) [][]int {
	res = nil
	subSetsRecur(nums, 0, []int{})
	return res
}

func subSetsRecur(nums []int, index int, curRes []int) {
	if index >= len(nums) {
		var tmp = make([]int, len(curRes)) // 如果直接放进去是浅拷贝，后续修改会影响，因此需要深拷贝进去。
		copy(tmp, curRes)
		res = append(res, tmp)
		return
	}

	subSetsRecur(nums, index+1, curRes) // 不需要值分支

	curRes = append(curRes, nums[index])

	subSetsRecur(nums, index+1, curRes) // 需要值分支
}
