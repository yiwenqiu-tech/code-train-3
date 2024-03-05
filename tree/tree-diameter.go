/* lc vip题目：求树的直径 */

package tree

import (
	"container/list"
)

var to [][]int // 出边数组
var pointNum int

func treeDiameter(edges [][]int) int {
	var maxPoint = -1
	for _, e := range edges {
		for _, p := range e {
			if p > maxPoint {
				maxPoint = p
			}
		}
	}
	pointNum = maxPoint + 1 // 树中的点数量

	// 构造出度数组
	to = make([][]int, pointNum)
	for _, e := range edges {
		a := e[0]
		b := e[1]
		to[a] = append(to[a], b)
		to[b] = append(to[b], a)
	}

	// 从任意点出发，找到距离最远的点p -- 该点肯定为树的直径的一端
	start := 0
	resOfStart := findDeepest(start)
	// 从p点出发，找到离p点最远的点 -- 找到树的直径
	resOfMaxDepth := findDeepest(resOfStart.point)
	return resOfMaxDepth.level - 1 // 层数 - 1是边
}

type deepestRes struct {
	point int
	level int
}

func findDeepest(start int) deepestRes {
	// 通过层序遍历求解层数量
	queue := list.List{}
	queue.PushBack(start)
	var depths = make([]int, pointNum)
	depths[start] = 1
	for queue.Len() != 0 {
		head := queue.Front()
		queue.Remove(head)
		headVal := head.Value.(int)
		for _, toVal := range to[headVal] {
			if depths[toVal] != 0 {
				continue // 已经遍历过，忽略
			}
			queue.PushBack(toVal)
			depths[toVal] = depths[headVal] + 1
		}
	}
	maxDepthNum := 0
	maxDepthPoint := -1
	for index, d := range depths {
		if d > maxDepthNum {
			maxDepthNum = d
			maxDepthPoint = index
		}
	}
	return deepestRes{maxDepthPoint, maxDepthNum}
}
