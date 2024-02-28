// priority_queue - 优先队列。
// 在golang库里的container/heap实现了堆，可基于此实现优先队列

package selfcontainer

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Value    string // 值
	Priority int    // 优先级

	index int // 在堆里的index，记录在Item里，方便更新其优先级时Fix
}

func (i *Item) String() string {
	return fmt.Sprintf("value: %v, priority: %v", i.Value, i.Priority)
}

type PriorityQueue []*Item

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].Priority > p[j].Priority // 从大到小
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	// 交换后，记得更新其index
	p[i].index = i
	p[j].index = j
}

func (p *PriorityQueue) Push(x any) { // 由于需要修改p，所以需要用引用
	curLen := len(*p)
	insert := x.(*Item)
	insert.index = curLen
	*p = append(*p, insert)
}

func (p *PriorityQueue) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1] // 由于heap pop时已经把最后的元素交换到最后了
	// 将待取出的元素引用置为空
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety，注意pop后这里index其实已经是最后一个元素的index了，没有参考意义
	// 缩容
	*p = old[0 : n-1]
	return item
}

func (p *PriorityQueue) Update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(p, item.index)
}
