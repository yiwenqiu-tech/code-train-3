package train

func sortedSquares(nums []int) []int {
	var res = make([]int, len(nums))
	left := 0
	right := len(nums) - 1
	resP := len(res) - 1
	for left <= right {
		leftNum := nums[left] * nums[left]
		rightNum := nums[right] * nums[right]
		if leftNum < rightNum {
			res[resP] = rightNum
			resP--
			right--
		} else {
			res[resP] = leftNum
			left++
			resP--
		}
	}
	return res
}
