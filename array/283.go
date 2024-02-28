package array

// ac
func moveZeroes(nums []int) {
	var curIndex = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[curIndex] = nums[i]
			curIndex++
		}
	}
	for i := curIndex; i < len(nums); i++ {
		nums[i] = 0
	}
}
