package sort

func ShellSort(nums []int) []int {
	numsLen := len(nums)
	initGap := numsLen / 2 // 初始分多少组
	for initGap >= 1 {
		for j := initGap; j < numsLen; j += 1 { // 从initGap开始遍历到最后，依次与同组之前的元素比较
			for k := j; k >= initGap; k -= initGap {
				if nums[k] < nums[k-initGap] {
					nums[k], nums[k-initGap] = nums[k-initGap], nums[k] // 前后两个对掉
				} else {
					break
				}
			}
		}
		initGap = initGap / 2
	}
	return nums
}
