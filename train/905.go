package train

func sortArrayByParity(nums []int) []int {
	larger := len(nums) - 1
	smaller := 0
	for smaller < larger {
		for smaller < larger && nums[smaller]&1 != 1 { // 找到第一个奇数
			smaller++
		}
		if smaller >= larger {
			break
		}
		for smaller < larger && nums[larger]&1 != 0 { // 找到第一个偶数
			larger--
		}
		if smaller >= larger {
			break
		}
		nums[smaller], nums[larger] = nums[larger], nums[smaller]
		smaller++
		larger--
	}
	return nums
}
