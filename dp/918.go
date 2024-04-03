package dp

import (
	"container/list"
	"math"
)

func maxSubarraySumCircular(nums []int) int {
	var doubleNums = make([]int, len(nums)*2) // 环形区间
	for i := 0; i < len(nums); i++ {
		doubleNums[i] = nums[i]
	}
	for i := len(nums); i < len(nums)*2; i++ {
		doubleNums[i] = nums[i-len(nums)]
	}

	var preSum = make([]int, len(nums)*2+1)
	preSum[0] = 0
	for i := 1; i < len(nums)*2+1; i++ {
		preSum[i] = preSum[i-1] + doubleNums[i-1]
	}

	var ans = math.MinInt32
	var queue list.List
	type item struct {
		index int
		value int
	}
	queue.PushBack(item{
		index: 0,
		value: 0,
	})
	for i := 1; i < len(nums)*2+1; i++ {
		// 暴力找，超时
		//for j := max(0, i-len(nums)); j < i; j++ {
		//	ans = max(ans, preSum[i]-preSum[j])
		//}

		// 维护单调队列
		front := queue.Front().Value.(item)
		for i-front.index > len(nums) { // 移除过期的
			queue.Remove(queue.Front())
			front = queue.Front().Value.(item)
		}
		ans = max(ans, preSum[i]-queue.Front().Value.(item).value) // 更新ans

		// 继续维护队列
		if queue.Len() > 0 { // 队列有数据
			back := queue.Back() // 队尾
			for back.Value.(item).value > preSum[i] {
				queue.Remove(back)
				if queue.Len() == 0 {
					break
				}
				back = queue.Back() // 队尾
			}
			queue.PushBack(item{
				index: i,
				value: preSum[i],
			})
		}
	}
	return ans
}
