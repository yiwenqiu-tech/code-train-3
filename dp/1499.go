package dp

import (
	"container/list"
	"math"
)

func findMaxValueOfEquation(points [][]int, k int) int {
	// 如果没有k限制时
	//n := len(points)
	//var ans = math.MinInt32
	//var tmp = math.MinInt32
	//for i := 1; i < n; i++ {
	//	tmp = max(tmp, points[i-1][1]-points[i-1][0])
	//	ans = max(ans, points[i][1]+points[i][0]+tmp)
	//}

	n := len(points)
	var ans = math.MinInt32
	type item struct {
		index int
		value int
	}
	var queue = list.New() // 单调队列，单调递减，队头最大。
	for i := 1; i < n; i++ {
		for queue.Len() > 0 {
			front := queue.Front()
			if points[i][0]-front.Value.(item).index > k { // 当前下标与队头大于k
				queue.Remove(front)
			} else {
				break // 跳出循环
			}
		}
		if points[i][0]-points[i-1][0] > k { // 与前一个都大于k了，则直接跳过计算即可
			continue
		}
		if queue.Len() > 0 { // 剩下没过期的，与准备插入的points[i-1]比较，若小于则移除
			back := queue.Back()
			for back.Value.(item).value < points[i-1][1]-points[i-1][0] {
				queue.Remove(back)
				if queue.Len() == 0 {
					break
				}
				back = queue.Back()
			}
		}
		queue.PushBack(item{
			index: points[i-1][0],
			value: points[i-1][1] - points[i-1][0],
		})
		ans = max(ans, points[i][1]+points[i][0]+queue.Front().Value.(item).value)
	}
	return ans
}
