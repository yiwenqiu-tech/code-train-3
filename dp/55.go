package dp

func canJump(nums []int) bool {
	if len(nums) == 1 { // 只有一个数字时，直接返回true
		return true
	}
	if nums[0] == 0 { // 在起点就动不了了，直接返回false
		return false
	}
	if len(nums) == 2 && nums[0] >= 1 { // 长度为2时，进不了循环判断，额外判断下，如果nums[0]>=1，则返回true
		return true
	}

	target := len(nums) - 1
	var f = make([]int, len(nums)) // 状态：记录到第i个位置时，可以到达的最远的下标
	f[0] = nums[0]

	for i := 1; i < len(nums)-1; i++ { // 最后一个数不统计，因为最后一个数跳到哪对题意没有意义了
		f[i] = max(i+nums[i], f[i-1])
		if f[i] >= target {
			return true
		}

		if f[i] == i {
			return false
		}
	}
	return false
}
