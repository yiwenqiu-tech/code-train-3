package binary_search

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums)
	var ans []int
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if right == len(nums) || nums[right] != target { // 如果>=target的第一个元素不是target，即不存在
		return []int{-1, -1}
	}
	ans = append(ans, right)

	left = 0
	right = len(nums) - 1
	for left < right {
		mid := left + (right-left+1)/2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	ans = append(ans, right)
	return ans
}
