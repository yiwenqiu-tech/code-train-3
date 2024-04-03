package dp

import "math"

func jump(nums []int) int {
	var f = make([]int, len(nums)) // 走到某个位置需要的最小步数
	for i := range f {
		f[i] = math.MaxInt32
	}
	f[0] = 0

	for i := 1; i < len(f); i++ {
		for j := 0; j < i; j++ {
			if nums[j]+j >= i { //能够到i位置的
				f[i] = min(f[i], f[j]+1)
			}
		}
	}
	return f[len(nums)-1]
}
