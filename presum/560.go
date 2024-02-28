package presum

func subarraySum(nums []int, k int) int {
	var preSum = make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	var res int
	var preSumNumMap = make(map[int]int)
	for i := 0; i < len(preSum); i++ {
		cur := preSum[i]
		find := cur - k
		if count, exist := preSumNumMap[find]; exist {
			res += count
		}
		preSumNumMap[cur]++
	}
	return res
}
