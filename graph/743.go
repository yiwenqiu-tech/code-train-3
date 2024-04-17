package graph

import (
	"container/heap"
	"math"
)

// Bellman-Ford做法
func networkDelayTime(times [][]int, n int, k int) int {
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
		if i == k {
			dist[i] = 0 // 起点赋值为0
		}
	}
	for i := 1; i <= n-1; i++ { // n-1轮
		flag := false
		for j := range times { // 遍历边
			u := times[j][0]
			v := times[j][1]
			w := times[j][2]
			if dist[u] == math.MaxInt32 { // 起点边为正无穷，直接跳过
				continue
			}
			if dist[u]+w < dist[v] { // 起点边 + w
				flag = true // 更新标识
				dist[v] = dist[u] + w
			}
		}
		// 本轮没有更新，直接break
		if !flag {
			break
		}
	}
	var ans = math.MinInt32
	for i := 1; i < len(dist); i++ {
		if dist[i] > ans {
			ans = dist[i]
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// Dijkstra算法
func networkDelayTime2(times [][]int, n int, k int) int {
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
		if i == k {
			dist[i] = 0 // 起点赋值为0
		}
	}
	type outsideItem struct {
		index  int
		length int
	}
	var outside = make([][]outsideItem, n+1) // 维护各个点的出边数组
	for t := range times {
		outside[times[t][0]] = append(outside[times[t][0]], outsideItem{
			index:  times[t][1],
			length: times[t][2],
		})
	}
	expand := make([]bool, n+1)
	for i := 1; i <= n; i++ { // n轮，每轮取当前最小的值来扩展
		var temp = math.MaxInt32
		minPoint := -1
		for j := 1; j <= n; j++ { // 从n个节点里找到没有扩展过的最小值
			if dist[j] < temp && !expand[j] {
				temp = dist[j]
				minPoint = j
			}
		}
		if minPoint == -1 { // 非法时，找不到minPoint，直接break
			break
		}
		expand[minPoint] = true
		outsideItems := outside[minPoint]
		for _, item := range outsideItems { // 更新当前最小值的出边
			if dist[item.index] > dist[minPoint]+item.length {
				dist[item.index] = dist[minPoint] + item.length
			}
		}
	}
	var ans = math.MinInt32
	for i := 1; i < len(dist); i++ {
		if dist[i] > ans {
			ans = dist[i]
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

type heapItem struct {
	index     int
	dist      int
	heapIndex int
}

type minHeap []*heapItem

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i].dist < (*m)[j].dist
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
	(*m)[i].heapIndex = i
	(*m)[j].heapIndex = j
}

func (m *minHeap) Push(x any) {
	item := x.(*heapItem)
	item.heapIndex = len(*m)
	*m = append(*m, item)
}

func (m *minHeap) Pop() any {
	item := (*m)[m.Len()-1]
	item.heapIndex = -1
	*m = (*m)[:m.Len()-1]
	return item
}

// Dijkstra算法，利用堆优化。
func networkDelayTime3(times [][]int, n int, k int) int {
	dist := make([]*heapItem, n+1)
	for i := range dist {
		dist[i] = &heapItem{
			index: i,
			dist:  math.MaxInt32,
		}
		if i == k { // 起点赋值为0
			dist[i].dist = 0
		}
	}
	type outsideItem struct {
		index  int
		length int
	}
	var outside = make([][]outsideItem, n+1) // 维护各个点的出边数组
	for t := range times {
		outside[times[t][0]] = append(outside[times[t][0]], outsideItem{
			index:  times[t][1],
			length: times[t][2],
		})
	}
	var mh minHeap
	mh.Push(dist[k]) // 把k push进去

	expand := make([]bool, n+1)
	for mh.Len() > 0 {
		minItem := heap.Pop(&mh).(*heapItem)
		expand[minItem.index] = true
		outsideItems := outside[minItem.index] // 取对应的出边更新
		for _, item := range outsideItems {    // 更新当前最小值的出边
			if dist[item.index].dist == math.MaxInt32 { // 之前没有放入堆上。
				dist[item.index].dist = minItem.dist + item.length
				heap.Push(&mh, dist[item.index])
			} else {
				if dist[item.index].dist > minItem.dist+item.length { // 更新dist
					dist[item.index].dist = minItem.dist + item.length
					heap.Fix(&mh, dist[item.index].heapIndex) // 已经Push过的，调整其位置
				}
			}
		}
	}
	var ans = math.MinInt32
	for i := 1; i < len(dist); i++ {
		if dist[i].dist > ans {
			ans = dist[i].dist
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
