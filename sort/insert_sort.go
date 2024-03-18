package sort

func InsertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		value := nums[i] // 取当前value值
		j := i - 1
		for ; j >= 0; j-- { // 找出i前一个位置，往前逐个对比
			if value < nums[j] { // 如果大于value的，往后挪下位置
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = value // 将value放在对应的位置
	}
	return nums
}

func InsertSort2(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j >= 1; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
	}
	return nums
}
