package dp

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	var f = make([]int, len(nums))
	var pre = make([]int, len(nums))
	for i := 0; i < len(f); i++ {
		f[i] = math.MinInt32
		pre[i] = -1
	}
	f[0] = nums[0] // 必须选一个

	for i := 1; i < len(f); i++ {
		if nums[i] > f[i-1]+nums[i] {
			f[i] = nums[i]
		} else {
			f[i] = f[i-1] + nums[i]
			pre[i] = i - 1 // 记录转移路径
		}
	}
	var ans int
	var end int
	for i := 0; i < len(f); i++ {
		if f[i] > ans {
			ans = f[i]
			end = i
		}
	}
	fmt.Println(end)
	fmt.Println(pre)
	print53(pre, end, nums)
	return ans
}

func print53(pre []int, i int, nums []int) {
	if pre[i] == -1 {
		fmt.Printf("%d ", nums[i]) // 注：第一个开始节点没有pre，但是也在结果里。
		return
	}
	print53(pre, pre[i], nums)
	fmt.Printf("%d ", nums[i])
}
