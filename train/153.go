package train

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	last := nums[right]
	for left < right {
		mid := left + (right-left)/2
		var res int
		if nums[mid] <= last { // 将原数组转换为01数组，找到第一个1（后继型）
			res = 1
		} else {
			res = 0
		}
		if res >= 1 {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[right]
}

func findMin2(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[right]
}
