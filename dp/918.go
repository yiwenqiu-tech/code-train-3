package dp

import (
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

	//var queue list.List // 方法2.队列里同时存放值和index
	//type item struct {
	//	index int
	//	value int
	//}
	//queue.PushBack(item{
	//	index: 0,
	//	value: 0,
	//})
	//for i := 1; i < len(nums)*2+1; i++ {
	//	// 方法1.暴力找，超时
	//	//for j := max(0, i-len(nums)); j < i; j++ {
	//	//	ans = max(ans, preSum[i]-preSum[j])
	//	//}
	//
	//	// 维护单调队列
	//	front := queue.Front().Value.(item)
	//	for i-front.index > len(nums) { // 移除过期的
	//		queue.Remove(queue.Front())
	//		front = queue.Front().Value.(item)
	//	}
	//	ans = max(ans, preSum[i]-queue.Front().Value.(item).value) // 更新ans
	//
	//	// 继续维护队列
	//	if queue.Len() > 0 { // 队列有数据
	//		back := queue.Back() // 队尾
	//		for back.Value.(item).value > preSum[i] {
	//			queue.Remove(back)
	//			if queue.Len() == 0 {
	//				break
	//			}
	//			back = queue.Back() // 队尾
	//		}
	//		queue.PushBack(item{
	//			index: i,
	//			value: preSum[i],
	//		})
	//	}
	//}

	//var queue list.List // 方法3.队列里只存放index
	//queue.PushBack(0)
	//for i := 1; i < len(nums)*2+1; i++ {
	//	// 维护单调队列
	//	for i-queue.Front().Value.(int) > len(nums) { // 移除过期的
	//		queue.Remove(queue.Front())
	//	}
	//	ans = max(ans, preSum[i]-preSum[queue.Front().Value.(int)]) // 更新ans
	//
	//	// 继续维护队列
	//	for queue.Len() > 0 && preSum[queue.Back().Value.(int)] > preSum[i] {
	//		queue.Remove(queue.Back())
	//	}
	//	queue.PushBack(i)
	//}
	return ans
}

func maxSubarraySumCircular2(nums []int) int {
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	maxChildSum := math.MinInt32
	minPreMin := math.MaxInt32
	for i := 1; i <= len(nums); i++ {
		// 更新preMin
		if preSum[i-1] < minPreMin {
			minPreMin = preSum[i-1]
		}
		if preSum[i]-minPreMin > maxChildSum {
			maxChildSum = preSum[i] - minPreMin
		}
	}

	minChildSum := math.MaxInt32
	maxPreMin := math.MinInt32
	// 从0-n-1期间，选一段最小值，注意不能全选。这里全选会导致最后结果数组为空，不符合题意
	for i := 2; i <= len(nums); i++ { // 这里先不选数组的头元素，计算出此时的最小和
		// 更新preMax
		if preSum[i-1] > maxPreMin {
			maxPreMin = preSum[i-1]
		}
		if preSum[i]-maxPreMin < minChildSum {
			minChildSum = preSum[i] - maxPreMin
		}

	}
	for i := 1; i < len(nums); i++ { // 这里不选数组的最后一个元素，与之前的最小和做对比
		if preSum[i] < minChildSum {
			minChildSum = preSum[i]
		}
	}

	if maxChildSum > preSum[len(nums)]-minChildSum {
		return maxChildSum
	} else {
		return preSum[len(nums)] - minChildSum
	}
}
