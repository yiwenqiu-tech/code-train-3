package binary_search

func splitArray(nums []int, k int) int {
	var left, right int
	for _, num := range nums {
		if num > left { // 最小值作为left值
			left = num
		}
		right += num // 累计值作为right
	}
	for left < right {
		mid := left + (right-left)/2
		if validateForSplitArray(nums, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func validateForSplitArray(nums []int, k int, cur int) bool {
	var need = 1
	var accu int

	for _, num := range nums {
		if accu+num > cur { // 累计大于cur了，则新开一个
			accu = num
			need++
		} else { // 否则继续累计
			accu += num
		}
	}
	return need <= k
}
