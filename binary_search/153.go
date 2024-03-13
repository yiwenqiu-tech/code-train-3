package binary_search

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		// 将条件转换为 nums[i] <= nums[len(nums) - 1], 数组转换为 0,0,..,1, 1
		if nums[mid] <= nums[len(nums)-1] { // 条件转换后，其实是后继型，找第一个1
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[right]
}
