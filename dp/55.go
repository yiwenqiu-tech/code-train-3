package dp

func canJump(nums []int) bool {
	if len(nums) == 1 { // 只有一个元素，本身就在终点，直接返回
		return true
	}
	var f = make([]int, len(nums)-1) // 记录每个位置可达到最远的地方，最后下标不需要统计了

	for i := 0; i < len(nums)-1; i++ {
		if i >= 1 {
			f[i] = max(i+nums[i], f[i-1]) // 当前位置+nums[i]，或与f[i-1]比较，取较大值
		} else {
			f[i] = nums[i] + i
		}
		if f[i] == i { // 原地踏步，则直接返回false即可
			return false
		}
		if f[i] >= len(nums)-1 { // 到终点了直接返回true
			return true
		}
	}
	return false
}
