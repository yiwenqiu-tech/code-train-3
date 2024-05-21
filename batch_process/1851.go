package batch_process

import (
	"container/heap"
	"sort"
)

type Event1851 struct {
	pos   int
	len   int
	tp    int
	index int
}

type event1851Heap []Event1851

func (e *event1851Heap) Len() int {
	return len(*e)
}

func (e *event1851Heap) Less(i, j int) bool {
	return (*e)[i].len < (*e)[j].len
}

func (e *event1851Heap) Push(x any) {
	*e = append(*e, x.(Event1851))
}

func (e *event1851Heap) Pop() any {
	item := (*e)[len(*e)-1]
	*e = (*e)[:len(*e)-1]
	return item
}

func (e *event1851Heap) Swap(i, j int) {
	(*e)[i], (*e)[j] = (*e)[j], (*e)[i]
}

func minInterval(intervals [][]int, queries []int) []int {
	var events []Event1851
	for i, interval := range intervals {
		events = append(events, Event1851{
			pos:   interval[0],
			len:   interval[1] - interval[0] + 1,
			tp:    1,
			index: i,
		})
		events = append(events, Event1851{
			pos:   interval[1],
			len:   interval[1] - interval[0] + 1,
			tp:    -1,
			index: i,
		})
	}
	for i, query := range queries {
		events = append(events, Event1851{
			pos:   query,
			len:   0,
			tp:    0,
			index: i,
		})
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i].pos < events[j].pos {
			return true
		} else if events[i].pos > events[j].pos {
			return false
		} else {
			return events[i].tp > events[j].tp
		}
	})

	var h event1851Heap
	var ans = make([]int, len(queries))
	var removed = make(map[int]struct{})
	for _, e := range events {
		if e.tp == 1 {
			heap.Push(&h, e)
		} else if e.tp == -1 {
			removed[e.index] = struct{}{}
		} else { // 查询
			curLen := -1
			for h.Len() > 0 {
				event := heap.Pop(&h).(Event1851)
				if _, exist := removed[event.index]; !exist {
					curLen = event.len
					ans[e.index] = curLen
					heap.Push(&h, event) // 放回去
					break
				}
			}
			if curLen == -1 { // 没找到
				ans[e.index] = -1
			}
		}
	}
	return ans
}
