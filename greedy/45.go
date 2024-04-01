package greedy

func jump(nums []int) int {
	var cur = 0
	var ans = 0

	for cur < len(nums)-1 {
		// 当前步已经够剩下的步数了
		if nums[cur] >= len(nums)-1-cur {
			ans++
			return ans
		}
		// 统计当前步走了后，下一步能到最远的地方
		longest := 0
		longestIndex := -1
		for i := 1; i <= nums[cur]; i++ {
			if i+nums[cur+i] > longest {
				longest = nums[cur+i] + i
				longestIndex = cur + i
			}
		}
		cur = longestIndex
		ans++
	}
	return ans
}
