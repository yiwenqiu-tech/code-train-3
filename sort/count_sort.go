package sort

func CountSort(nums []int) []int {
	tmp := make([]int, 100001) // 题目-5 * 10⁴ <= nums[i] <= 5 * 10⁴
	for _, num := range nums {
		tmp[num+50000]++ // 所以这里需要+50000，以避免负数，最后遍历tmp的时候再减回去
	}
	index := 0
	for num, c := range tmp {
		for c > 0 {
			nums[index] = num - 50000
			c--
			index++
		}
	}
	return nums
}
