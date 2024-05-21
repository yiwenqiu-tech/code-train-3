package batch_process

import (
	"container/heap"
	"sort"
)

type Event struct {
	pos    int
	height int
	tp     int
	index  int
}

type eventHeap []Event

func (e *eventHeap) Len() int {
	return len(*e)
}

func (e *eventHeap) Less(i, j int) bool {
	return (*e)[i].height > (*e)[j].height
}

func (e *eventHeap) Push(x any) {
	*e = append(*e, x.(Event))
}

func (e *eventHeap) Pop() any {
	item := (*e)[len(*e)-1]
	*e = (*e)[:len(*e)-1]
	return item
}

func (e *eventHeap) Swap(i, j int) {
	(*e)[i], (*e)[j] = (*e)[j], (*e)[i]
}

func getSkyline(buildings [][]int) [][]int {
	var events []Event
	for i, build := range buildings {
		events = append(events, Event{
			pos:    build[0],
			height: build[2],
			tp:     1,
			index:  i,
		})
		events = append(events, Event{
			pos:    build[1],
			height: build[2],
			tp:     -1,
			index:  i,
		})
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].pos < events[j].pos
	})

	var removed = make(map[int]struct{})
	var h eventHeap

	var ans [][]int
	for index, e := range events {
		if e.tp == 1 { // 新增的
			heap.Push(&h, e)
		} else { // 删除的
			removed[e.index] = struct{}{}
		}
		// 如果当前位置为事件的最后一个或者当前位置已经为当前pos的最后一个
		if index == len(events)-1 || e.pos != events[index+1].pos {
			curTop := 0
			for h.Len() > 0 {
				event := heap.Pop(&h).(Event)
				if _, exist := removed[event.index]; !exist { // 没被删除
					curTop = event.height
					heap.Push(&h, event) // 没被删除的放回去
					break
				}
			}

			if len(ans) == 0 || ans[len(ans)-1][1] != curTop {
				ans = append(ans, []int{e.pos, curTop})
			}
		}
	}
	return ans
}
