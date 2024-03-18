package sort

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ { // 需要遍历 len(nums)-1 次，最后一轮只比较前两个元素
		for j := 0; j < len(nums)-1-i; j++ { // 每一轮，比较len(nums)-1-i次前后元素的大小，例如第一轮进来，需要比较 0和1，1和2，..., len(nums)-2和len(nums)-1
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
