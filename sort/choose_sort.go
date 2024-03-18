package sort

// 选择排序
func chooseSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		var minIndex = i
		for j := i + 1; j < len(nums); j++ { // 遍历最小的index，与当前位置交换
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex { // 最小index不等于i，则交换
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
	return nums
}
